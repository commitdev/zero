module github.com/yourrepo/hello-world

go 1.12

replace github.com/yourrepo/hello-world-idl => ./hello-world-idl

require (
	github.com/grpc-ecosystem/grpc-gateway v1.11.3 // indirect
	github.com/yourrepo/hello-world-idl v0.0.0
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.24.0 // indirect
)
