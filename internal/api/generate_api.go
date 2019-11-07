package main


import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	GitRepo     string `json:"gitRepo"`
}

type ProjectConfiguration struct {
	ProjectName       string       `json:"projectName"`
	FrontendLanguage  string       `json:"frontendLanguage"`
	Organization      string       `json:"organization"`
	Description       string       `json:"description"`
	Maintainers       []Maintainer `json:"maintainers"`
	Services          []Service    `json:"services"`
}

func generateProject(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var projectConfig ProjectConfiguration
		err := decoder.Decode(&projectConfig)
		if err != nil {
			panic(err)
		}
		log.Println(projectConfig.ProjectName)
		CreateProject(projectConfig)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Post successful"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not found"}`))
	}

}

func main() {
	var router = mux.NewRouter()
	var api = router.PathPrefix("/v1/generate").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(generateProject)

	log.Fatal(http.ListenAndServe(":8080", router))
}
