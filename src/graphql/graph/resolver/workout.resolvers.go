package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
)

func (r *mutationResolver) AddUserWorkout(ctx context.Context, input model.AddUserWorkoutInput) (string, error) {
	// check if the user email exists!
	validated := r.PgRepo.ValidateUser(input.UserEmail)

	if !validated {
		return "You need to sign up first", nil
	}

	// assign user email to userworkout plan
	userWorkoutPlan, err := r.PgRepo.AssignUserWorkoutPlan(input)
	if err != nil {
		return "there was an error", nil
	}

	// update user table
	result := r.PgRepo.UpdateUser(input, userWorkoutPlan.ID)
	if !result {
		return "there was an error for update user", nil
	}

	if input.Exercises != nil {
		return r.PgRepo.AddEachExercises(userWorkoutPlan.ID, input.Exercises)
	}

	return fmt.Sprintf("User %v has been added", input.UserEmail), nil
}

func (r *queryResolver) GetUserWorkoutPlansByEmail(ctx context.Context, input string) ([]*model.UserWorkoutPlan, error) {
	return r.PgRepo.GetUserWorkoutPlansByEmail(ctx, input)
}

func (r *queryResolver) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	return r.PgRepo.GetAllUserWorkoutPlans(ctx)
}
