package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/cmd/graphql/graph/generated"
	"github.com/cindy1408/gym/src/cmd/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (string, error) {
	muscleGroup := []*model.MuscleGroup{}
	muscleGroup = append(muscleGroup, &model.MuscleGroup{input.MuscleGroup})

	newExercise := &model.BaseExercise{
		ID:            uuid.New().String(),
		Name:          input.Name,
		MuscleGroup:   muscleGroup,
		SpecificParts: []*model.Body{{input.SpecificParts}},
		Level:         input.Level,
		AvoidGiven:    []*model.AvoidGiven{{input.AvoidGiven}},
		MovementType:  input.MovementType,
	}
	fmt.Println(newExercise)
	// r.baseExercise = append(r.baseExercise, newExercise)

	return "exercise was successfully added", nil
}

func (m *mutationResolver) CreateHydrate() (string, error) {
	str, err := m.baseExercise.Hydrate()
	if err != nil {
		return "", errors.Wrapf(err, "m.baseExercise.Hydrate")
	}
	return str, nil 
}

func (r *queryResolver) BaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	panic(fmt.Errorf("not implemented"))
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
