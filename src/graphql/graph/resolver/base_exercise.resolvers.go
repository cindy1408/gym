package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"github.com/pkg/errors"
)

func (r *queryResolver) GetBaseExerciseByName(ctx context.Context, input string) (*model.BaseExercise, error) {
	return postgres.GetBaseExerciseByName(ctx, r.DB, input)
}

func (r *queryResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	return postgres.GetAllBaseExercise(ctx, r.DB)
}

func (r *queryResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise, err := postgres.UpdateBaseExercise(ctx, r.DB, input)
	if err != nil {
		return nil, errors.Wrap(err, "postgres.UpdateBaseexercise")
	}

	return updatedExercise, nil
}

func (r *queryResolver) HydrateBaseExercise(ctx context.Context) (string, error) {
	result, err := postgres.HydrateBaseExercise(ctx, r.DB)
	if err != nil {
		return "", errors.Wrap(err, "postgres.HydrateBaseExercise")
	}

	return result, nil
}

func (r *queryResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (string, error) {
	if input.Name == "" ||
		input.MuscleGroup == "" ||
		input.SpecificParts == "" ||
		input.Level == 0 ||
		input.MovementType == "" {
		return "at least one of the required field is missing", nil
	}

	exists := postgres.ValidateBaseExercise(ctx, r.DB, input.Name)

	if !exists {
		return "base exercise already exists", nil
	}

	result, err := postgres.AddBaseExercise(ctx, r.DB, input)
	if err != nil {
		return "", errors.Wrap(err, "postgres.AddBaseExercise")
	}

	return result, nil
}
