package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/commitdev/commit0/internal/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func generateProject(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var projectConfig config.Commit0Config
		err := decoder.Decode(&projectConfig)
		if err != nil {
			panic(err)
		}
		log.Println(projectConfig.Name)
		// createProject(projectConfig)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Post successful"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not found"}`))
	}
}

func Commit0Api() {
	// React Frontend is served on port 3000 while in development mode.
	allowOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	var router = mux.NewRouter()
	var api = router.PathPrefix("/v1/generate").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(generateProject)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("internal/ui/build/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	buildHandler := http.FileServer(http.Dir("internal/ui/build"))
	router.PathPrefix("/").Handler(buildHandler)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowOrigins, allowedMethods, allowedHeaders)(router)))
}
