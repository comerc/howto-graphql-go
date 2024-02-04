//go:generate go run github.com/99designs/gqlgen --verbose
package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"example/internal/gql/runtime"
)

type Resolver struct {
	// todos []*model.Todo
}

func NewRootResolver() runtime.ResolverRoot {
	return &Resolver{
		// todos: []*model.Todo{
		// 	{
		// 		ID:   "Todo:1",
		// 		Text: "Buy a cat food",
		// 		// State:    StateNotYet,
		// 		// Verified: false,
		// 	},
		// 	{
		// 		ID:   "Todo:2",
		// 		Text: "Check cat water",
		// 		// State:    StateDone,
		// 		// Verified: true,
		// 	},
		// 	{
		// 		ID:   "Todo:3",
		// 		Text: "Check cat meal",
		// 		// State:    StateDone,
		// 		// Verified: true,
		// 	},
		// },
	}
}
