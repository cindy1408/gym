package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"github.com/pkg/errors"
)

func (r *mutationResolver) AddExercise(ctx context.Context, input *model.AddExerciseInput) (string, error) {
	validated := postgres.ValidateUser(r.DB, input.UserEmail)
	if !validated {
		return "Please create an account first", nil
	}

	// validate user workout plan
	workoutPlanID, validated := postgres.ValidateUserWorkoutPlan(r.DB, input.UserEmail, input.GymDay)
	if !validated {
		return "You need to create a workout plan", nil
	}

	return postgres.AddEachExercises(r.DB, workoutPlanID, input.EachExercise)
}

func (r *mutationResolver) IncreaseRep(ctx context.Context, input model.IncreaseInput) (*model.EachExercise, error) {
	exercise, err := postgres.Increase(ctx, r.DB, input, model.Rep)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) IncreaseSet(ctx context.Context, input model.IncreaseInput) (*model.EachExercise, error) {
	exercise, err := postgres.Increase(ctx, r.DB, input, model.Set)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) UpdateEachExercise(ctx context.Context, input model.UpdateExerciseInput) (*model.EachExercise, error) {
	userDetails, err := postgres.GetUserByEmail(ctx, r.DB, input.UserEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetUserByEmail")
	}

	requestedExercise, err := postgres.GetExerciseByNameAndWorkoutPlanID(r.DB, input.EachExercise.Name, *userDetails.UserWorkoutPlanID)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetExerciseByNameAndWorkoutPlanID")
	}

	requestedExercise.Name = input.EachExercise.Name
	requestedExercise.Weight = input.EachExercise.Weight
	requestedExercise.Unit = input.EachExercise.Unit
	requestedExercise.Reps = input.EachExercise.Reps
	requestedExercise.Sets = input.EachExercise.Sets

	err = postgres.UpdateExercise(r.DB, requestedExercise)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.UpdateExercise")
	}

	return requestedExercise, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	return postgres.GetAllEachExercise(r.DB)
}
