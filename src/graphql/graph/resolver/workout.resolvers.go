package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
)

func (r *mutationResolver) AddUserWorkout(ctx context.Context, input model.AddUserWorkoutInput) (string, error) {
	// check if the user email exists!
	validated := postgres.ValidateUser(r.DB, input.UserEmail)

	if !validated {
		return "You need to sign up first", nil
	}

	// assign user email to userworkout plan
	userWorkoutPlan, err := postgres.AssignUserWorkoutPlan(r.DB, input)
	if err != nil {
		return "there was an error", nil
	}

	// update user table
	result := postgres.UpdateUser(r.DB, input, userWorkoutPlan.ID)
	if !result {
		return "there was an error for update user", nil
	}

	if input.Exercises != nil {
		return postgres.AddEachExercises(r.DB, userWorkoutPlan.ID, input.Exercises)
	}

	return fmt.Sprintf("User %v has been added", input.UserEmail), nil
}

func (r *queryResolver) GetUserWorkoutPlansByEmail(ctx context.Context, input string) (*model.UserWorkoutPlan, error) {
	userWorkout := model.UserWorkoutPlan{}
	r.DB.Model(&model.UserWorkoutPlan{}).Where("email = ?", input).Scan(&userWorkout)
	return &userWorkout, nil 
}

func (r *queryResolver) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	userWorkoutPlans := []*model.UserWorkoutPlan{}
	r.DB.Table("user_workout_plan").Scan(&userWorkoutPlans)
	return userWorkoutPlans, nil
}
