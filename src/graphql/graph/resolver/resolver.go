package resolver

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
	postgres.Postgres
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *Resolver) Init() error {
	err := r.DB.AutoMigrate(
		&model.BaseExercise{},
		&model.User{},
		&model.MuscleGroup{},
		&model.SpecificParts{},
		&model.UserWorkoutPlan{},
		&model.EachExercise{},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return r.DB.Error
}

func Hasher(toHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(toHash))
	return hex.EncodeToString(hasher.Sum(nil))
}
