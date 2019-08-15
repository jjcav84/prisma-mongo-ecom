//go:generate gorunpkg github.com/99designs/gqlgen

package tmp

import (
	context "context"
	graphql "prisma-examples/go/graphql"
	main "prisma-examples/go/graphql/server"
)

type Resolver struct{}

func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]graphql.User, error) {
	panic("not implemented")
}
