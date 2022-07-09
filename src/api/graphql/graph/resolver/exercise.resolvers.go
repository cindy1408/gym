package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/api/graphql/graph/model"
	"github.com/pkg/errors"
)

func (r *mutationResolver) AddExercise(ctx context.Context, input *model.AddExerciseInput) (string, error) {
	validated := r.PgRepo.ValidateUser(input.UserEmail)
	if !validated {
		return "Please create an account first", nil
	}

	// validate user workout plan
	workoutPlanID, validated := r.PgRepo.ValidateUserWorkoutPlan(input.UserEmail, input.GymDay)
	if !validated {
		return "You need to create a workout plan", nil
	}

	return r.PgRepo.AddEachExercises(workoutPlanID, input.EachExercise)
}

func (r *mutationResolver) IncreaseRep(ctx context.Context, input model.IncreaseInput) (*model.EachExercise, error) {
	target := model.Details{
		Rep: true, 
	}
	exercise, err := r.PgRepo.Increase(ctx, input, target)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) IncreaseSet(ctx context.Context, input model.IncreaseInput) (*model.EachExercise, error) {
	target := model.Details{
		Set: true, 
	}

	exercise, err := r.PgRepo.Increase(ctx, input, target)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) UpdateEachExercise(ctx context.Context, input model.UpdateExerciseInput) (*model.EachExercise, error) {
	userDetails, err := r.PgRepo.GetUserByEmail(ctx, input.UserEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetUserByEmail")
	}

	requestedExercise, err := r.PgRepo.GetExerciseByNameAndWorkoutPlanID(input.EachExercise.Name, *userDetails.UserWorkoutPlanID)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetExerciseByNameAndWorkoutPlanID")
	}

	requestedExercise.Name = input.EachExercise.Name
	requestedExercise.Weight = input.EachExercise.Weight
	requestedExercise.Unit = input.EachExercise.Unit
	requestedExercise.Reps = input.EachExercise.Reps
	requestedExercise.Sets = input.EachExercise.Sets

	err = r.PgRepo.UpdateExercise(requestedExercise)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.UpdateExercise")
	}

	return requestedExercise, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	return r.PgRepo.GetAllEachExercise(ctx)
}
