package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"github.com/google/uuid"
)

func (r *mutationResolver) AddExercise(ctx context.Context, input *model.AddExerciseInput) (string, error) {
	validated := postgres.ValidateUser(input.UserEmail)

	if !validated {
		return "Please create an account first", nil
	}

	// validate user workout plan
	workoutPlanID, validated := postgres.ValidateUserWorkoutPlan(input.UserEmail, input.GymDay)
	if !validated {
		return "You need to create a workout plan", nil
	}

	for _, eachExercise := range input.EachExercise {
		addUserExercise := model.EachExercise{
			ID:                uuid.NewString(),
			UserWorkoutPlanID: workoutPlanID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Unit:              eachExercise.Unit,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}
		r.DB.Create(&addUserExercise)
	}

	return "Your account has been successfully created", nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	panic(fmt.Errorf("not implemented"))
}
