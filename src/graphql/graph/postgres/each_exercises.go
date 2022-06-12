package postgres

import (
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
)

func (r Resolver) AddEachExercisesToDB(userWorkoutID string, eachExercises []*model.EachExerciseInput) (string, error) {
	for _, eachExercise := range eachExercises {
		addExercise := &model.EachExercise{
			ID:                uuid.New().String(),
			UserWorkoutPlanID: userWorkoutID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}

		r.DB.Create(&addExercise)
	}
	return "User's workout and exercises has been updated", nil
}
