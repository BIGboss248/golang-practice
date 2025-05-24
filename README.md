# Golang practice

## Table of Contents

- [Golang practice](#golang-practice)
  - [Table of Contents](#table-of-contents)
  - [graphql fundamentals](#graphql-fundamentals)
    - [Schema Definition Language](#schema-definition-language)
      - [General types](#general-types)
      - [Query types](#query-types)
      - [Mutations types](#mutations-types)
      - [Subscriptions types](#subscriptions-types)
    - [Query data from API](#query-data-from-api)
      - [Query with parameters](#query-with-parameters)
    - [Modify data with mutations](#modify-data-with-mutations)
    - [Subscribe to data](#subscribe-to-data)
  - [gplgen](#gplgen)
  - [Cobra](#cobra)

my simple repository to learn more about golang

## graphql fundamentals

### Schema Definition Language

GraphQL has its own type system that’s used to define the schema of an API. The syntax for writing schemas is called Schema Definition Language (SDL).

The schema is one of the most important concepts when working with a GraphQL API. It specifies the capabilities of the API and defines how clients can request the data. It is often seen as a contract between the server and client.

Generally, a schema is simply a collection of GraphQL types. However, when writing the schema for an API, there are some special root types:

1. type Query { ... }
2. type Mutation { ... }
3. type Subscription { ... }

#### General types

For example below we defined two types, person and posts and if they are related we express them in each others defenition meaning a person can have many posts so we put posts: [Post!]!
and on the other end in post we put auther: Person!
meaning each post is written by a Person

```json

type Person {
  name: String!
  age: Int!
  posts: [Post!]!
}

type Post {
  title: String!
  author: Person!
}

```

#### Query types

As shown above we defined our types now for the client to be able to [query](#query-data-from-api) these data we have to create a query type for example we define allPersons Query type with last attribute

```json

type Query {
  allPersons(last: Int): [Person!]!
}

```

#### Mutations types

Now that we can query the data we might want to be able to update the data of backend server through the API ([mutaitions](#modify-data-with-mutations)) to deine a way to mutate data throuhg the API we define a mutations type

```json

type Mutation {
  createPerson(name: String!, age: Int!): Person!
}

```

#### Subscriptions types

Finally to implement the [subscription](#subscribe-to-data) function (the function in which the client is modified by the server on data update) we define a subscription type

```json

type Subscription {
  newPerson: Person!
}

```

### Query data from API

In graphql we explicitly say what we need and we only get that nothing more nothing less and thats one of the advantages of graphql over REST it prevents over or under fetching
an exaple of a query is below

```json

{
  allPersons {
    name
  }
}

```

the respone we recive is below

```json

{
  "allPersons": [
    { "name": "Johnny" },
    { "name": "Sarah" },
    { "name": "Alice" }
  ]
}

```

#### Query with parameters

we can query data with specific parameters for exaple the query below will return the last 2 persons

```json

{
  allPersons(last: 2) {
    name
  }
}

```

### Modify data with mutations

You can create, update, delete data of the backend server by querying API this is called mutations

To modify data we use **mutation** keyword in our query and pass in arguments for exaple to create a person we query bellow

```json

mutation {
  createPerson(name: "Bob", age: 36)
}

```

### Subscribe to data

Another important requirement for many applications today is to have a realtime connection to the server in order to get immediately informed about important events. For this use case, GraphQL offers the concept of subscriptions.

When a client subscribes to an event, it will initiate and hold a steady connection to the server. Whenever that particular event then actually happens, the server pushes the corresponding data to the client. Unlike queries and mutations that follow a typical “request-response-cycle”, subscriptions represent a stream of data sent over to the client.

Subscriptions are written using the same syntax as queries and mutations. Here’s an example where we subscribe on events happening on the Person type:

```json

subscription {
  newPerson {
    name
    age
  }
}

```

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

## Cobra

Cobra is a package to create a command line app using golang

first we initialize our golang repository and install package

```console

go mod init github.com/{username}/{repo name}
go install github.com/spf13/cobra-cli@latest

```

After that we initialize a cobra-cli repository

```console

cobra-cli init

```

now that we initialized our repository we can edit cmd\root.go to edit the description of what our cli tool works

now we can add more commands using this command

```console

cobra-cli add {command name}

```

after we run this command a file is added in cmd directory
