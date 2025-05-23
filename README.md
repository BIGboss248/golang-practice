# Golang practice

my simple repository to learn more about golang

## gplgen

The package to create production grade graphql server

to start with it first get the package and run init

```console
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

```

This will generate files and code to create a graphql server using gplgen

after that we modift graph\schema.graphqls to modify our schema and use the command bellow to generate go code and models

```console

go run github.com/99designs/gqlgen generate

```
