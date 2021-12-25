package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pkbhowmick/kql/schema"
)

func execQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		log.Printf("errors occured: %v\n", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serveing request for", r.URL.String())
		result := execQuery(r.URL.Query().Get("query"), schema.PodSchema)
		json.NewEncoder(w).Encode(result)
	})
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

// curl command: curl -g 'http://localhost:8080/graphql?query={pod{apiVersion,kind}}'
