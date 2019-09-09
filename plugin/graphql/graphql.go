package graphql

import (
	"bytes"
	"fmt"
	"github.com/bi-foundation/protobuf-graphql-extension/graphqlproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"strconv"
	"strings"
	"unicode"
)

const graphqlTypes = "graphqlproto.types"

type graphql struct {
	*generator.Generator
	generator.PluginImports
	messages []*generator.Descriptor
	oneofs   map[*descriptor.OneofDescriptorProto]*oneof
	enums    []*generator.EnumDescriptor
}

type oneof struct {
	message      *generator.Descriptor
	fields       []*descriptor.FieldDescriptorProto
	messageIndex int
	oneofIndex   int
}

func init() {
	generator.RegisterPlugin(NewGraphQL())
}

func NewGraphQL() *graphql {
	return &graphql{}
}

func (p *graphql) Name() string {
	return "graphql"
}

func (p *graphql) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *graphql) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.messages = make([]*generator.Descriptor, 0)
	p.oneofs = make(map[*descriptor.OneofDescriptorProto]*oneof)
	p.enums = make([]*generator.EnumDescriptor, 0)

	if graphqlproto.GetGraphQLFile(file.FileDescriptorProto) != true {
		return
	}

	graphQLPkg := p.NewImport("github.com/graphql-go/graphql")
	schemaPkg := p.NewImport("github.com/bi-foundation/protobuf-graphql-extension/plugin/graphql/scalars")
	fmtPkg := p.NewImport("fmt")

	for _, enum := range file.Enums() {
		p.P(`var `, p.graphQLTypeEnumName(enum), ` *`, graphQLPkg.Use(), `.Enum`)
		p.enums = append(p.enums, enum)
	}

	for mi, message := range file.Messages() {
		if len(message.DescriptorProto.Field) == 0 {
			continue
		}

		if message.DescriptorProto.GetOptions().GetMapEntry() {
			p.P(`var `, p.graphQLTypeVarName(message), ` = `, schemaPkg.Use(), `.Map`)
			continue
		}

		p.messages = append(p.messages, message)

		// interfaces for unions
		tname := p.TypeName(message)
		p.P(`type `, tname, `Getter interface{`)
		p.In()
		p.P(`Get`, tname, `() *`, tname)
		p.Out()
		p.P(`}`)

		// var declarations
		p.P(`var `, p.graphQLTypeVarName(message), ` *`, graphQLPkg.Use(), `.Object`)

		for i, field := range message.DescriptorProto.OneofDecl {
			p.P(`var `, graphQLUnionVarName(message, field), ` *`, graphQLPkg.Use(), `.Union`)

			// collect the unions to make them easier to access in the file
			p.oneofs[field] = oneofFields(message, mi, i)
		}
	}

	// getter funcs for oneof fields
	for _, oo := range p.oneofs {
		ccTypeName := generator.CamelCaseSlice(oo.message.TypeName())

		// hack our structs to define getters
		for _, field := range oo.fields {
			obj := p.ObjectNamed(field.GetTypeName())
			tname := generator.CamelCaseSlice(obj.TypeName())
			fname := generator.CamelCase(field.GetName())

			p.P(`func (g *`, ccTypeName, `_`, fname, `) Get`, tname, `() *`, p.TypeName(obj), ` {`)
			p.In()
			p.P(`return g.`, fname)
			p.Out()
			p.P(`}`)
		}
	}

	p.P()
	p.P(`func init() {`)
	p.In()

	// enum type declarations
	for _, enum := range p.enums {
		p.P(p.graphQLTypeEnumName(enum), ` = `, graphQLPkg.Use(), `.NewEnum(`, graphQLPkg.Use(), `.EnumConfig{`)
		p.In()
		p.P(`Name:        "`, p.TypeName(enum), `",`)
		p.P(`Values: `, graphQLPkg.Use(), `.EnumValueConfigMap{`)
		p.In()

		for _, enumValue := range enum.GetValue() {
			p.P(`"`, enumValue.GetName(), `": &`, graphQLPkg.Use(), `.EnumValueConfig{`)
			p.P(`Value: `, int(enumValue.GetNumber()), `,`)
			p.P(`},`)
		}
		p.Out()
		p.P(`},`)
		p.Out()
		p.P(`})`)
	}

	for mi, message := range p.messages {
		messageGQL := p.comment(fmt.Sprintf("4,%d", mi))
		ccTypeName := generator.CamelCaseSlice(message.TypeName())

		p.P(p.graphQLTypeVarName(message), ` = `, graphQLPkg.Use(), `.NewObject(`, graphQLPkg.Use(), `.ObjectConfig{`)
		p.In()
		p.P(`Name:        "`, p.TypeName(message), `",`)
		p.P(`Description: `, messageGQL, `,`)
		p.P(`Fields: (`, graphQLPkg.Use(), `.FieldsThunk)(func() `, graphQLPkg.Use(), `.Fields {`)
		p.In()
		p.P(`return `, graphQLPkg.Use(), `.Fields{`)
		p.In()
		for fi, field := range message.DescriptorProto.Field {
			// skip defining a regular object field for unions, that comes next
			if field.OneofIndex != nil {
				continue
			}

			var (
				fieldGQL = p.comment(fmt.Sprintf("4,%d,2,%d", mi, fi))
				gtype, _ = p.GoType(message, field)
				hasStar  = strings.Index(gtype, "*") == 0
			)

			p.P(`"`, field.GetName(), `": &`, graphQLPkg.Use(), `.Field{`)
			p.In()
			p.P(`Type:        `, p.graphQLType(message, field, graphQLPkg, schemaPkg), `,`)
			p.P(`Description: `, fieldGQL, `,`)
			p.P(`Resolve: func(p `, graphQLPkg.Use(), `.ResolveParams) (interface{}, error) {`)
			p.In()
			p.P(`obj, ok := p.Source.(*`, ccTypeName, `)`)
			p.P(`if ok {`)
			p.In()

			if hasStar {
				p.P(`if obj.`, p.GetFieldName(message, field), ` == nil {`)
				p.In()
				p.P(`return nil, nil`)
				p.Out()
				p.P(`}`)
				p.P(`return obj.Get`, p.GetFieldName(message, field), `(), nil`)
			} else {
				if field.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
					p.P(`return int(`, p.graphQLEnumLookupType(field), `_value[obj.`, p.GetFieldName(message, field), `.String()]), nil`)
				} else {
					p.P(`return obj.`, p.GetFieldName(message, field), `, nil`)
				}
			}

			p.Out()
			p.P(`}`)
			p.P(`inter, ok := p.Source.(`, ccTypeName, `Getter)`)
			p.P(`if ok {`)
			p.In()
			p.P(`face := inter.Get`, ccTypeName, `()`)
			p.P(`if face == nil {`)
			p.In()
			p.P(`return nil, nil`)
			p.Out()
			p.P(`}`)

			if hasStar {
				p.P(`if face.`, p.GetFieldName(message, field), ` == nil {`)
				p.In()
				p.P(`return nil, nil`)
				p.Out()
				p.P(`}`)
				p.P(`return face.Get`, p.GetFieldName(message, field), `(), nil`)
			} else {
				if field.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
					p.P(`return int(`, p.graphQLEnumLookupType(field), `_value[face.`, p.GetFieldName(message, field), `.String()]), nil`)
				} else {
					p.P(`return face.`, p.GetFieldName(message, field), `, nil`)
				}
			}

			p.Out()
			p.P(`}`)
			p.P(`return nil, `, fmtPkg.Use(), `.Errorf("field `, field.GetName(), ` not resolved")`)
			p.Out()
			p.P(`},`)
			p.Out()
			p.P(`},`)
		}
		for fi, field := range message.DescriptorProto.OneofDecl {
			fieldGQL := p.comment(fmt.Sprintf("4,%d,8,%d", mi, fi))

			p.P(`"`, field.GetName(), `": &`, graphQLPkg.Use(), `.Field{`)
			p.In()
			p.P(`Type:        `, graphQLUnionVarName(message, field), `,`)
			p.P(`Description: `, fieldGQL, `,`)
			p.P(`Resolve: func(p `, graphQLPkg.Use(), `.ResolveParams) (interface{}, error) {`)
			p.In()
			p.P(`obj, ok := p.Source.(*`, ccTypeName, `)`)
			p.P(`if !ok {`)
			p.In()
			p.P(`return nil, `, fmtPkg.Use(), `.Errorf("field `, field.GetName(), ` not resolved")`)
			p.Out()
			p.P(`}`)
			p.P(`return obj.Get`, generator.CamelCase(field.GetName()), `(), nil`)
			p.Out()
			p.P(`},`)
			p.Out()
			p.P(`},`)

		}
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`}),`)
		p.Out()
		p.P(`})`)
	}

	// declare our unions last, since the types will have needed to be defined from all messages first
	for decl, oo := range p.oneofs {
		ccTypeName := generator.CamelCaseSlice(oo.message.TypeName())
		fieldGQL := p.comment(fmt.Sprintf("4,%d,8,%d", oo.messageIndex, oo.oneofIndex))

		p.P(graphQLUnionVarName(oo.message, decl), ` = `, graphQLPkg.Use(), `.NewUnion(`, graphQLPkg.Use(), `.UnionConfig{`)
		p.In()
		p.P(`Name:        "`, graphQLUnionName(oo.message, decl), `",`)
		p.P(`Description: `, fieldGQL, `,`)
		p.P(`Types:       []*`, graphQLPkg.Use(), `.Object{`)
		p.In()
		for _, field := range oo.fields {
			p.P(p.graphQLTypeVarName(p.ObjectNamed(field.GetTypeName())), `,`)
		}
		p.Out()
		p.P(`},`)
		p.P(`ResolveType: func (p `, graphQLPkg.Use(), `.ResolveTypeParams) *`, graphQLPkg.Use(), `.Object {`)
		p.In()
		for _, field := range oo.fields {
			obj := p.ObjectNamed(field.GetTypeName())
			fname := generator.CamelCase(field.GetName())

			p.P(`if _, ok := p.Value.(*`, ccTypeName, `_`, fname, `); ok {`)
			p.In()
			p.P(`return `, p.graphQLTypeVarName(obj))
			p.Out()
			p.P(`}`)
		}
		p.P(`return nil`)
		p.Out()
		p.P(`},`)
		p.Out()
		p.P(`})`)
	}

	p.Out()
	p.P(`}`)
}

func (p *graphql) graphQLType(message *generator.Descriptor, field *descriptor.FieldDescriptorProto, pkgName, schemaPkgName generator.Single) string {
	var gqltype string
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE, descriptor.FieldDescriptorProto_TYPE_FLOAT:
		gqltype = fmt.Sprint(pkgName.Use(), ".", "Float")
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32, descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32, descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SINT64:
		gqltype = fmt.Sprint(pkgName.Use(), ".", "Int")
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		gqltype = fmt.Sprint(pkgName.Use(), ".", "Boolean")
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		gqltype = fmt.Sprint(pkgName.Use(), ".", "String")
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		panic("mapping a proto group type to graphql is unimplemented")
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		mobj := p.ObjectNamed(field.GetTypeName())
		gqltype = p.graphQLTypeEnumName(mobj)
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		// TODO: fix this to be more robust about imported objects
		mobj := p.ObjectNamed(field.GetTypeName())
		// fmt.Fprint(os.Stderr, mobj.PackageName())
		if strings.EqualFold(mobj.File().GetPackage(), graphqlTypes) {
			gqltype = fmt.Sprint(schemaPkgName.Use(), ".", generator.CamelCaseSlice(mobj.TypeName()))
			break
		}

		gqltype = p.graphQLTypeVarName(mobj)
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		gqltype = fmt.Sprint(schemaPkgName.Use(), ".", "ByteString")
	default:
		panic("unknown proto field type")
	}

	if field.IsRepeated() && !p.IsMap(field) {
		gqltype = fmt.Sprint(pkgName.Use(), ".NewList(", gqltype, ")")
	}

	if field.IsRequired() {
		gqltype = fmt.Sprint(pkgName.Use(), ".NewNonNull(", gqltype, ")")
	}

	return gqltype
}

func (p *graphql) comment(path string) string {
	return strconv.Quote(strings.TrimSpace(p.Comments(path)))
}

func (p *graphql) graphQLTypeVarName(obj generator.Object) string {
	return fmt.Sprint(p.DefaultPackageName(obj), "GraphQL", generator.CamelCaseSlice(obj.TypeName()), "Type")
}

func (p *graphql) graphQLTypeEnumName(obj generator.Object) string {
	return fmt.Sprint(p.DefaultPackageName(obj), "GraphQL", generator.CamelCaseSlice(obj.TypeName()), "Enum")
}

func (p *graphql) graphQLEnumLookupType(field *descriptor.FieldDescriptorProto) string {
	return generator.CamelCaseSlice(p.ObjectNamed(field.GetTypeName()).TypeName())
}

func graphQLUnionName(message *generator.Descriptor, oneof *descriptor.OneofDescriptorProto) string {
	return generator.CamelCaseSlice(message.TypeName()) + generator.CamelCase(oneof.GetName())
}

func graphQLUnionVarName(message *generator.Descriptor, oneof *descriptor.OneofDescriptorProto) string {
	return fmt.Sprint("GraphQL", graphQLUnionName(message, oneof), "Union")
}

func oneofFields(message *generator.Descriptor, messageIndex, oneofIndex int) *oneof {
	fields := make([]*descriptor.FieldDescriptorProto, 0)

	for _, field := range message.DescriptorProto.Field {
		if field.OneofIndex != nil && *field.OneofIndex == int32(oneofIndex) {
			fields = append(fields, field)
		}
	}

	return &oneof{message, fields, messageIndex, oneofIndex}
}

func snakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)
	out := bytes.NewBuffer(make([]byte, 0, length))

	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out.WriteRune('_')
		}
		out.WriteRune(unicode.ToLower(runes[i]))
	}

	return out.String()
}
