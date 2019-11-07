# Commit0 Api

## Usage
 - To run:
 `go run internal/api/generate_api.go internal/api/create_project.go`
 - Endpoint:
 `localhost:8080/{version}/generate`
 - Post request body json example:
 ``` {
      "projectName":"funApp",
      "language":"go",
      "organization":"commit org",
      "description":"this app will do amazing things",
      "gitRepoName":"fun-repo",
      "maintainers":[
         {
            "name":"Lill",
            "email":"ll@gmail.com"
         },
         {
            "name":"Pi",
            "email":"pi@live.ca"
         }
      ]
   }
   ```
