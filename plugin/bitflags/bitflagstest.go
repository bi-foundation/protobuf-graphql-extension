package bitflags

import (
	"github.com/bi-foundation/protobuf-graphql-extension/graphqlproto"
	"github.com/gogo/protobuf/plugin/testgen"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type test struct {
	*generator.Generator
}

func init() {
	testgen.RegisterTestPlugin(NewTest)
}

func NewTest(g *generator.Generator) testgen.TestPlugin {
	return &test{g}
}

func (p *test) Generate(imports generator.PluginImports, file *generator.FileDescriptor) bool {
	for _, message := range file.Messages() {
		if graphqlproto.IsBitflags(file.FileDescriptorProto, message.DescriptorProto) {
			return true
		}
	}
	return false
}
