package resolver

import(
	"github.com/miraikeitai2020/backend-bff/pkg/bff"
)

type Resolver struct{}

// Mutation returns bff.MutationResolver implementation.
func (r *Resolver) Mutation() bff.MutationResolver { return &mutationResolver{r} }

// Query returns bff.QueryResolver implementation.
func (r *Resolver) Query() bff.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
