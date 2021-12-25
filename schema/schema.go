package schema

import (
	"github.com/graphql-go/graphql"
)

type Pod struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

var podType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pod",
	Fields: graphql.Fields{
		"apiVersion": &graphql.Field{
			Type: graphql.String,
		},
		"kind": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PodSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",

	Fields: graphql.Fields{
		"pod": &graphql.Field{
			Type:        podType,
			Description: "Get single pod",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &Pod{
					Kind:       "pod",
					APIVersion: "v1",
				}, nil
			},
		},
	},
})
