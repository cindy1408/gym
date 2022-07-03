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
	return r.Postgres.GetBaseExerciseByName(ctx, input)
}

func (r *queryResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	postgres := postgres.Postgres{}
	return r.GetAllBaseExercise(ctx)
}

func (r *queryResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	postgres := postgres.Postgres{}
	updatedExercise, err := postgres.UpdateBaseExercise(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "postgres.UpdateBaseexercise")
	}

	return updatedExercise, nil
}

func (r *queryResolver) HydrateBaseExercise(ctx context.Context) (string, error) {
	postgres := postgres.Postgres{}
	result, err := postgres.HydrateBaseExercise(ctx)
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
	postgres := postgres.Postgres{}
	exists := postgres.ValidateBaseExercise(ctx, input.Name)

	if !exists {
		return "base exercise already exists", nil
	}

	result, err := postgres.AddBaseExercise(ctx, input)
	if err != nil {
		return "", errors.Wrap(err, "postgres.AddBaseExercise")
	}

	return result, nil
}
