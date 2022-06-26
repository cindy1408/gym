package postgres

import (
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func AddEachExercises(db *gorm.DB, userWorkoutID string, eachExercises []*model.EachExerciseInput) (string, error) {
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
	return "user exercises has been added", nil
}

func GetExerciseByID(db *gorm.DB, id string) (*model.EachExercise, error) {
	var exercise *model.EachExercise
	db.Model(&model.EachExercise{}).Where("id = ?", id).Scan(&exercise)

	return exercise, nil
}

func GetExerciseByNameAndWorkoutPlanID(db *gorm.DB, exerciseName string, userWorkoutID string) (*model.EachExercise, error) {
	var requestedExercise *model.EachExercise
	db.Model(&model.EachExercise{}).Where("name = ? AND user_workout_plan_id = ?", exerciseName, userWorkoutID).Scan(&requestedExercise)

	return requestedExercise, nil
}

func UpdateExercise(db *gorm.DB, updateExercise *model.EachExercise) error {
	result := db.Model(&model.EachExercise{}).Where("id = ?", updateExercise.ID).Updates(&updateExercise)

	if result.RowsAffected == 0 {
		return errors.Wrapf(result.Error, "issue with updating db")
	}

	return nil
}
