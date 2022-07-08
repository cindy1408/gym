package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/api"
	"github.com/cindy1408/gym/src/api/graphql/graph/model"
	"github.com/pkg/errors"
)

func (r *queryResolver) GetBaseExerciseByName(ctx context.Context, input string) (*model.BaseExercise, error) {
	baseExercise, err := r.PgRepo.GetBaseExerciseByName(ctx, input)
	if err != nil {
		return nil, err
	}
	exercise := api.InternalBaseExerciseToExternalMapper(*baseExercise)

	return &exercise, nil
}

func (r *queryResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	allBaseExercise, err := r.PgRepo.GetAllBaseExercise(ctx)
	if err != nil {
		return nil, err
	}

	var exercises []*model.BaseExercise
	for _, baseExercise := range allBaseExercise {
		exercise := api.InternalBaseExerciseToExternalMapper(*baseExercise)
		exercises = append(exercises, &exercise)
	}

	return exercises, nil
}

func (r *queryResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	baseExerciseInput := api.ExternalBaseExerciseInputToInternalMapper(*input)
	
	updatedExercise, err := r.PgRepo.UpdateBaseExercise(ctx, &baseExerciseInput)
	if err != nil {
		return nil, errors.Wrap(err, "postgres.UpdateBaseexercise")
	}

	updateExercise := api.InternalBaseExerciseToExternalMapper(*updatedExercise)
	return &updateExercise, nil
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

	internalInput := api.ExternalBaseExerciseInputToInternalMapper(*input)
	result, err := r.PgRepo.AddBaseExercise(ctx, &internalInput)
	if err != nil {
		return "", errors.Wrap(err, "postgres.AddBaseExercise")
	}

	return result, nil
}
