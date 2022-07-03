package postgres

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (p Postgres) AddEachExercises(userWorkoutID string, eachExercises []*model.EachExerciseInput) (string, error) {
	for _, eachExercise := range eachExercises {
		addExercise := &model.EachExercise{
			ID:                uuid.New().String(),
			UserWorkoutPlanID: userWorkoutID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}

		p.db.Create(&addExercise)
	}
	return "user exercises has been added", nil
}

func (p Postgres) GetExerciseByID(id string) (*model.EachExercise, error) {
	var exercise *model.EachExercise
	p.db.Model(&model.EachExercise{}).Where("id = ?", id).Scan(&exercise)

	return exercise, nil
}

func (p Postgres) GetExerciseByNameAndWorkoutPlanID(exerciseName string, userWorkoutID string) (*model.EachExercise, error) {
	var requestedExercise *model.EachExercise
	p.db.Model(&model.EachExercise{}).Where("name = ? AND user_workout_plan_id = ?", exerciseName, userWorkoutID).Scan(&requestedExercise)

	return requestedExercise, nil
}

func (p Postgres) UpdateExercise(updateExercise *model.EachExercise) error {
	result := p.db.Model(&model.EachExercise{}).Where("id = ?", updateExercise.ID).Updates(&updateExercise)

	if result.RowsAffected == 0 {
		return errors.Wrapf(result.Error, "issue with updating db")
	}

	return nil
}

func (p Postgres) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	allEachExercises := []*model.EachExercise{}
	p.db.Table("each_exercises").Scan(&allEachExercises)
	return allEachExercises, nil
}
