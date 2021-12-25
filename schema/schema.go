package schema

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	Phase     string `json:"phase"`
}

var PodList []Pod

var podType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pod",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"namespace": &graphql.Field{
			Type: graphql.String,
		},
		"replicas": &graphql.Field{
			Type: graphql.Int,
		},
		"phase": &graphql.Field{
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
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"namespace": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, ok := params.Args["name"].(string)
				if !ok {
					return nil, errors.New("name is not provided")
				}
				ns, ok := params.Args["namespace"].(string)
				if !ok {
					return nil, errors.New("namespace is not provided")
				}

				for _, p := range PodList {
					if p.Name == name && p.Namespace == ns {
						return p, nil
					}
				}
				return &Pod{}, nil
			},
		},

		"podList": &graphql.Field{
			Type:        graphql.NewList(podType),
			Description: "List of Pods",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return PodList, nil
			},
		},
	},
})
