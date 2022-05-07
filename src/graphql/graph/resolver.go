package graph

import (
	"github.com/cindy1408/gym/src/graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// used to pull the interface in
type Resolver struct {
	baseExercises    []*model.BaseExercise
	users            []*model.User
	exercises        []*model.EachExercise
	userWorkoutPlans []*model.UserWorkoutPlan
	userWorkoutDays  []*model.WorkoutPerDay
}
