package schema

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Node      string `json:"node"`
	Phase     string `json:"phase"`
}

var PodList map[string]Pod

func init() {
	PodList = make(map[string]Pod)
}

var podType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pod",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"namespace": &graphql.Field{
			Type: graphql.String,
		},
		"node": &graphql.Field{
			Type: graphql.String,
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

				key := fmt.Sprintf("%s/%s", ns, name)
				val, ok := PodList[key]
				if !ok {
					return &Pod{}, nil
				}
				return val, nil
			},
		},

		"pods": &graphql.Field{
			Type:        graphql.NewList(podType),
			Description: "List of Pods",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				l := make([]Pod, 0, len(PodList))
				for _, v := range PodList {
					l = append(l, v)
				}
				return l, nil
			},
		},
	},
})
