package flavortown

import (
	"bytes"
	"encoding/json"
	"testing"

	dessert "github.com/bi-foundation/protobuf-graphql-extension/examples/dessert"
	opsee_types "github.com/bi-foundation/protobuf-graphql-extension/opseeproto/types"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func init() {
	perms := opsee_types.NewPermissionsBitmap("peas", "cornbread", "nothing", "another thing", "???")
	opsee_types.PermissionsRegistry.Register("flavortown", perms)
}

func TestSchema(t *testing.T) {
	// some sides
	sa, err := opsee_types.NewPermissions("flavortown", "peas", "nothing")
	assert.Nil(t, err)
	sb, err := opsee_types.NewPermissions("flavortown", "peas")
	assert.Nil(t, err)
	sc, err := opsee_types.NewPermissions("flavortown", "cornbread")
	assert.Nil(t, err)

	populatedMenu := &Menu{
		Items: []*LineItem{
			{
				Dish: &LineItem_Lunch{&Lunch{
					Name:        "hogslop",
					Description: []byte("disgusting"),
					Tags: map[string]string{
						"coolness": "no",
						"tips":     "frosted",
					},
				}},
				PriceCents:     100,
				CreatedAt:      &opsee_types.Timestamp{100, 100},
				UpdatedAt:      &opsee_types.Timestamp{200, 200},
				Sides:          sa,
				QualityControl: Quality_FAIR,
			},
			{
				Dish: &LineItem_TastyDessert{&dessert.Dessert{
					Name:      "coolwhip",
					Sweetness: 9,
				}},
				PriceCents:     50,
				CreatedAt:      &opsee_types.Timestamp{100, 100},
				UpdatedAt:      &opsee_types.Timestamp{200, 200},
				Sides:          sb,
				Nothing:        nil,
				QualityControl: Quality_EXPENSIVE,
			},
			{
				Dish: &LineItem_TastyDessert{&dessert.Dessert{
					Name:      "coolwhip",
					Sweetness: 9,
				}},
				PriceCents:     50,
				CreatedAt:      &opsee_types.Timestamp{100, 100},
				UpdatedAt:      &opsee_types.Timestamp{200, 200},
				Sides:          sc,
				Nothing:        nil,
				QualityControl: Quality_CHEAP,
			},
		},
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"menu": &graphql.Field{
					Type: GraphQLMenuType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return populatedMenu, nil
					},
				},
			},
		}),
	})

	if err != nil {
		t.Fatal(err)
	}

	queryResponse := graphql.Do(graphql.Params{Schema: schema, RequestString: `query goodQuery {
		menu {
			items {
				dish {
					... on Lunch {
						name
						description
						tags
					}
					... on Dessert {
						name
						sweetness
					}
				}
				price_cents
				created_at
				updated_at
				sides
				nothing {
					void
				}
				qualityControl
			}
		}
	}`})

	if queryResponse.HasErrors() {
		t.Fatalf("graphql query errors: %#v\n", queryResponse.Errors)
	}

	t.Logf("repsonse: %s", formatResponse(queryResponse))

	lunchitem := populatedMenu.Items[0]
	assert.Equal(t, lunchitem.GetLunch().Name, getProp(queryResponse.Data, "menu", "items", 0, "dish", "name"))
	assert.Equal(t, string(lunchitem.GetLunch().Description), getProp(queryResponse.Data, "menu", "items", 0, "dish", "description"))
	assert.Equal(t, lunchitem.GetLunch().Tags, getProp(queryResponse.Data, "menu", "items", 0, "dish", "tags"))
	assert.EqualValues(t, lunchitem.PriceCents, getProp(queryResponse.Data, "menu", "items", 0, "price_cents"))
	assert.EqualValues(t, lunchitem.CreatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 0, "created_at"))
	assert.EqualValues(t, lunchitem.UpdatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 0, "updated_at"))
	assert.EqualValues(t, lunchitem.Sides.Permissions(), getProp(queryResponse.Data, "menu", "items", 0, "sides"))
	assert.EqualValues(t, lunchitem.QualityControl.String(), getProp(queryResponse.Data, "menu", "items", 0, "qualityControl"))
	t.Logf("%v", getProp(queryResponse.Data, "menu", "items", 0, "sides"))

	dessertitem := populatedMenu.Items[1]
	assert.Equal(t, dessertitem.GetTastyDessert().Name, getProp(queryResponse.Data, "menu", "items", 1, "dish", "name"))
	assert.EqualValues(t, dessertitem.GetTastyDessert().Sweetness, getProp(queryResponse.Data, "menu", "items", 1, "dish", "sweetness"))
	assert.EqualValues(t, dessertitem.PriceCents, getProp(queryResponse.Data, "menu", "items", 1, "price_cents"))
	assert.EqualValues(t, dessertitem.CreatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 1, "created_at"))
	assert.EqualValues(t, dessertitem.UpdatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 1, "updated_at"))
	assert.EqualValues(t, dessertitem.Sides.Permissions(), getProp(queryResponse.Data, "menu", "items", 1, "sides"))
	assert.EqualValues(t, dessertitem.QualityControl.String(), getProp(queryResponse.Data, "menu", "items", 1, "qualityControl"))

	dessertitem = populatedMenu.Items[2]
	assert.Equal(t, dessertitem.GetTastyDessert().Name, getProp(queryResponse.Data, "menu", "items", 2, "dish", "name"))
	assert.EqualValues(t, dessertitem.GetTastyDessert().Sweetness, getProp(queryResponse.Data, "menu", "items", 2, "dish", "sweetness"))
	assert.EqualValues(t, dessertitem.PriceCents, getProp(queryResponse.Data, "menu", "items", 2, "price_cents"))
	assert.EqualValues(t, dessertitem.CreatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 2, "created_at"))
	assert.EqualValues(t, dessertitem.UpdatedAt.Millis(), getProp(queryResponse.Data, "menu", "items", 2, "updated_at"))
	assert.EqualValues(t, dessertitem.Sides.Permissions(), getProp(queryResponse.Data, "menu", "items", 2, "sides"))
	assert.EqualValues(t, dessertitem.QualityControl.String(), getProp(queryResponse.Data, "menu", "items", 2, "qualityControl"))

}

func formatResponse(v interface{}) string {
	jsonObject, err := json.Marshal(v)
	if err != nil {
		return "failed to format"
	}

	var out bytes.Buffer
	err = json.Indent(&out, []byte(jsonObject), "", "\t")
	if err != nil {
		return string(jsonObject)
	}
	return out.String()
}

func getProp(i interface{}, path ...interface{}) interface{} {
	cur := i

	for _, s := range path {
		switch cur.(type) {
		case map[string]interface{}:
			cur = cur.(map[string]interface{})[s.(string)]
			continue
		case []interface{}:
			cur = cur.([]interface{})[s.(int)]
			continue
		default:
			return cur
		}
	}

	return cur
}
