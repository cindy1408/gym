package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/packages/api/graphql/graph/model"
	"github.com/pkg/errors"
)

func (r *queryResolver) GetBaseExerciseByName(ctx context.Context, input string) (*model.BaseExercise, error) {
	baseExercise, err := r.PgRepo.GetBaseExerciseByName(ctx, input)
	if err != nil {
		return nil, err
	}

	return baseExercise, nil
}

func (r *queryResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	allBaseExercise, err := r.GetAllBaseExercise(ctx)
	if err != nil {
		return nil, err
	}

	var exercises []*model.BaseExercise
	exercises = append(exercises, allBaseExercise...)

	return exercises, nil
}

func (r *queryResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise, err := r.PgRepo.UpdateBaseExercise(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "postgres.UpdateBaseexercise")
	}

	return updatedExercise, nil
}

func (r *queryResolver) HydrateBaseExercise(ctx context.Context) (string, error) {
	result, err := r.HydrateBaseExercise(ctx)
	if err != nil {
		return "", errors.Wrap(err, "postgres.HydrateBaseExercise")
	}

	return result, nil
}

func (r *queryResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (string, error) {
	if input.Name == "" ||
		input.MuscleGroup == "" ||
		len(input.SpecificParts) == 0 ||
		input.Level == 0 ||
		input.MovementType == "" {
		return "at least one of the required field is missing", nil
	}

	exists := r.PgRepo.ValidateBaseExercise(ctx, input.Name)

	if !exists {
		return "base exercise already exists", nil
	}

	result, err := r.PgRepo.AddBaseExercise(ctx, input)
	if err != nil {
		return "", errors.Wrap(err, "postgres.AddBaseExercise")
	}

	return result, nil
}
