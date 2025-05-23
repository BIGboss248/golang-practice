# Golang practice

my simple repository to learn more about golang

## gplgen

The package to create production grade graphql server
You can install the package to use locally with

```console
go get github.com/99designs/gqlgen
go install  github.com/99designs/gqlgen

```

and run the rest of the commands with gqlgen command line or use go run to run the requiered commands as done below

to start with it first get the package and run init

```console
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

```

This will generate files and code to create a graphql server using gplgen

after that we modift graph\schema.graphqls to modify our schema and use the command bellow to generate go code and models

```console

go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen generate

```

now that we generated our models and some code we need to implement out resolvers (function that will get the data querried from the client to return it to the client)
The resolvers are in graph\schema.resolvers.go

