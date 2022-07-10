package resolver

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/cindy1408/gym/packages/api/graphql/graph/generated"
	"github.com/cindy1408/gym/packages/api/postgres"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	postgres.PgRepo
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func Hasher(toHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(toHash))
	return hex.EncodeToString(hasher.Sum(nil))
}
