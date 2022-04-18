package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	newExercise := model.BaseExercise{
		ID:            uuid.New().String(),
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}
	r.baseExercise = append(r.baseExercise, &newExercise)

	return &newExercise, nil
}

func (r *mutationResolver) HydrateBaseExercise(ctx context.Context) ([]*model.BaseExercise, error) {
	for _, eachBaseExercise := range Data {
		newExercise := model.BaseExercise{
			ID:            uuid.New().String(),
			Name:          eachBaseExercise.Name,
			MuscleGroup:   eachBaseExercise.MuscleGroup,
			SpecificParts: eachBaseExercise.SpecificParts,
			Level:         eachBaseExercise.Level,
			AvoidGiven:    eachBaseExercise.AvoidGiven,
			MovementType:  eachBaseExercise.MovementType,
		}
		r.baseExercise = append(r.baseExercise, &newExercise)
	}
	return r.baseExercise, nil
}

func (r *queryResolver) BaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	return r.baseExercise, nil
}

func (r *queryResolver) GetAllAvaliableBaseExercises(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllWorkoutDay(ctx context.Context) ([]*model.WorkoutPerDay, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
