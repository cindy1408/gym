package postgres

import (
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddEachExercisesToDB(db *gorm.DB, userWorkoutID string, eachExercises []*model.EachExerciseInput) (string, error) {
	for _, eachExercise := range eachExercises {
		addExercise := &model.EachExercise{
			ID:                uuid.New().String(),
			UserWorkoutPlanID: userWorkoutID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}

		db.Create(&addExercise)
	}
	return "User's workout and exercises has been updated", nil
}
