package graphql

import (
	"github.com/graphql-go/graphql"
)

// Define os tipos
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"nome":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		},
	},
)

/*
var pedidoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pedido",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.Int},
			"produto": &graphql.Field{Type: graphql.String},
			"userId":  &graphql.Field{Type: graphql.Int},
		},
	},
)
*/
// Define o schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"user": &graphql.Field{
						Type: userType,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{Type: graphql.Int},
						},
						Resolve: resolveUser,
					},
					/*"pedido": &graphql.Field{
						Type: pedidoType,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{Type: graphql.Int},
						},
						Resolve: resolvePedido,
					},*/
				},
			},
		),
	},
)
