package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pkbhowmick/kql/schema"
)

func init() {
	pod1 := schema.Pod{Name: "nginx", Namespace: "demo", Replicas: 1, Phase: "Ready"}
	pod2 := schema.Pod{Name: "mongodb", Namespace: "demo", Replicas: 1, Phase: "Ready"}
	pod3 := schema.Pod{Name: "postgres", Namespace: "demo", Replicas: 1, Phase: "Ready"}

	schema.PodList = append(schema.PodList, pod1, pod2, pod3)
}

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
		log.Println("Get request from url: ", r.URL.String())
		result := execQuery(r.URL.Query().Get("query"), schema.PodSchema)
		json.NewEncoder(w).Encode(result)
	})
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

// curl command (pod): curl -g 'http://localhost:8080/graphql?query={pod(name:"nginx",namespace:"demo"){replicas,phase}}'
// curl command (podList): curl -g 'http://localhost:8080/graphql?query={podList{name,namespace,phase}}'
