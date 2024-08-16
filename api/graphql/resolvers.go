package graphql

import (
	"estudo-test/internal/service"

	"github.com/graphql-go/graphql"
)

var UserService service.UserService

//var pedidoService service.PedidoService

/*
	func NewUserResolver(userService service.UserService) UserResolver {
		return &userResolver{
			UserService: userService,
		}
	}
*/
func resolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		return UserService.GetUsersById(int64(id))
	}
	return nil, nil
}

/*
func resolvePedido(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int64)
	if ok {
		return pedidoService.GetUsersById(id)
	}
	return nil, nil
}
*/
