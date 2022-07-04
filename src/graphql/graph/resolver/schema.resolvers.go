package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
)

func (r *mutationResolver) HydrateMuscleGroups(ctx context.Context) (string, error) {
	err := r.PgRepo.HydrateMuscleGroups(ctx)
	if err != nil {
		return "", errors.Wrap(err, "HydrateMuscleGroups")
	}
	return "Muscle group table has been hydrated", nil
}

func (r *mutationResolver) HydrateSpecificParts(ctx context.Context) (string, error) {
	err := r.PgRepo.HydrateSpecificParts(ctx)
	if err != nil {
		errors.Wrap(err, "HydrateSpecificParts")
	}

	return "Specific Parts has been hydrated", nil
}

func (r *queryResolver) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) ([]string, error) {

	return r.PgRepo.GetMuscleSpecifics(ctx, input)
}
