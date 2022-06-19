package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
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

func (r *mutationResolver) IncreaseRep(ctx context.Context, input model.IncreaseRepInput) (*model.EachExercise, error) {
	var requestedExercise *model.EachExercise
	r.DB.Model(&model.EachExercise{}).Where("name = ? AND user_email = ?", input.ExerciseName, input.UserEmail).Scan(&requestedExercise)

	requestedExercise.Reps = requestedExercise.Reps + 1

	r.DB.Model(&model.EachExercise{}).Where("name = ? AND user_email = ?", input.ExerciseName, input.UserEmail).Updates(&requestedExercise)

	return requestedExercise, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	allEachExercises := []*model.EachExercise{}
	r.DB.Table("each_exercises").Scan(&allEachExercises)
	return allEachExercises, nil
}
