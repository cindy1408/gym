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

	exercise, err := r.increase(ctx, input, "rep")
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) IncreaseSet(ctx context.Context, input model.IncreaseInput) (*model.EachExercise, error) {

	exercise, err := r.increase(ctx, input, "set")
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return exercise, nil
}

func (r *mutationResolver) UpdateEachExercise(ctx context.Context, input model.UpdateExerciseInput) (*model.EachExercise, error) {
	var requestedExercise *model.EachExercise

	var userDetails *model.User
	r.DB.Model(&model.User{}).Where("email = ?", input.UserEmail).Scan(&userDetails)

	r.DB.Model(&model.EachExercise{}).Where("name = ? AND user_workout_plan_id = ?", input.EachExercise.Name, userDetails.UserWorkoutPlanID).Scan(&requestedExercise)

	requestedExercise.Name = input.EachExercise.Name
	requestedExercise.Weight = input.EachExercise.Weight
	requestedExercise.Unit = input.EachExercise.Unit
	requestedExercise.Reps = input.EachExercise.Reps
	requestedExercise.Sets = input.EachExercise.Sets

	r.DB.Model(&model.EachExercise{}).Where("name = ? AND user_workout_plan_id = ?", input.EachExercise.Name, userDetails.UserWorkoutPlanID).Updates(&requestedExercise)

	return requestedExercise, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	allEachExercises := []*model.EachExercise{}
	r.DB.Table("each_exercises").Scan(&allEachExercises)
	return allEachExercises, nil
}

func (r *mutationResolver) increase(ctx context.Context, input model.IncreaseInput, target string) (*model.EachExercise, error) {

	userDetails, err := postgres.GetUserByEmail(ctx, r.DB, input.UserEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetUserByEmail")
	}

	requestedExercise, err := postgres.GetExerciseByNameAndWorkoutPlanID(r.DB, input.ExerciseName, *userDetails.UserWorkoutPlanID)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetExerciseByNameAndWorkoutPlanID")
	}

	if target == "set" {
		requestedExercise.Sets = requestedExercise.Sets + 1
	} else if target == "rep" {
		requestedExercise.Reps = requestedExercise.Reps + 1
	} else {
		return nil, errors.New("target needs to be either set or rep")
	}

	r.DB.Model(&model.EachExercise{}).Where("name = ? AND user_workout_plan_id = ?", input.ExerciseName, userDetails.UserWorkoutPlanID).Updates(&requestedExercise)

	// Grab the data from database and return it
	eachExercise, err := postgres.GetExerciseByID(r.DB, requestedExercise.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return eachExercise, nil
}
