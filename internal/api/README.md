# Commit0 Api

## Usage
 - To run:
 `commit0 api`
 - Endpoint:
    - POST `localhost:8080/{version}/generate`
 - Post request body json example:
 ``` {
      "projectName":"funApp",
      "frontendFramework":"react",
      "organization":"commit org",
      "description":"this app will do amazing things",
      "maintainers":[
         {
            "name":"Lill",
            "email":"ll@gmail.com"
         },
         {
            "name":"Pi",
            "email":"pi@live.ca"
         }
      ],
      "services":[
         {
            "name":"user",
            "description":"user service",
            "language":"go",
            "gitRepo":"github.com/user"
         },
         {
            "name":"account",
            "description":"bank account service",
            "language":"go",
            "gitRepo":"github.com/account"
         }
      ]
   }
   ```
