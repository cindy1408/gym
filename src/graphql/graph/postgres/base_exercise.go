package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func UpdateBaseExercise(ctx context.Context, db *gorm.DB, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}

	result := db.Debug().Model(&model.BaseExercise{}).Where("name = ?", input.Name).Updates(updatedExercise)

	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to update base exercise")
	}

	return &updatedExercise, nil
}

func HydrateBaseExercise(ctx context.Context, db *gorm.DB) (string, error) {
	for _, eachBaseExercise := range graph.BaseExerciseData {
		rows, err := db.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()
		if err != nil {
			fmt.Printf("%v , selecting database\n", eachBaseExercise.Name)
		}
		defer rows.Close()

		var name, avoidGiven string
		var count int

		for rows.Next() {
			rows.Scan(&name, &avoidGiven)
			if avoidGiven == "" {
				if name == eachBaseExercise.Name {
					count += 1
					fmt.Printf("%v , exists in database!\n", eachBaseExercise.Name)
				}
			} else if name == eachBaseExercise.Name && &avoidGiven == eachBaseExercise.AvoidGiven {
				fmt.Printf("%v , exists in database!\n", eachBaseExercise.Name)
			}
		}
		if count == 0 {
			db.Create(eachBaseExercise)
		}
	}

	return "Base exercise table has been hydrated!", nil
}

func AddBaseExercise(ctx context.Context, db *gorm.DB, baseExercise *model.BaseExerciseInput) (string, error) {
	newExercise := model.BaseExercise{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	}
	db.Create(newExercise)

	return fmt.Sprintf("base exercise %v has been added", baseExercise.Name), nil
}

func ValidateBaseExercise(ctx context.Context, db *gorm.DB, name string) bool {
	rows, err := db.Model(&model.BaseExercise{}).Select("name").Rows()
	if err != nil {
		fmt.Println("issue with base exercise table")
	}

	defer rows.Close()

	var existingBaseExercise string
	var exists int
	for rows.Next() {
		exists++
		rows.Scan(&existingBaseExercise)
		if existingBaseExercise == name {
			return true
		}
	}

	return exists != 0
}

func GetAllBaseExercise(ctx context.Context, db *gorm.DB) ([]*model.BaseExercise, error) {
	allBaseExercises := []*model.BaseExercise{}
	db.Table("base_exercises").Scan(&allBaseExercises)
	return allBaseExercises, nil
}

func GetBaseExerciseByName(ctx context.Context, db *gorm.DB, name string) (*model.BaseExercise, error) {
	if name == "" {
		return nil, errors.Wrapf(nil, "base exercise name is empty")
	}
	var baseExercise *model.BaseExercise
	result := db.Where("name", name).Find(&model.BaseExercise{}).Scan(&baseExercise)

	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to find base exercise")
	}

	return baseExercise, nil 
}

func Increase(ctx context.Context, db *gorm.DB, input model.IncreaseInput, target model.Details) (*model.EachExercise, error) {
	userDetails, err := GetUserByEmail(ctx, db, input.UserEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetUserByEmail")
	}

	requestedExercise, err := GetExerciseByNameAndWorkoutPlanID(db, input.ExerciseName, *userDetails.UserWorkoutPlanID)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetExerciseByNameAndWorkoutPlanID")
	}

	if target == model.Set {
		requestedExercise.Sets = requestedExercise.Sets + 1
	} else if target == model.Rep {
		requestedExercise.Reps = requestedExercise.Reps + 1
	} else {
		return nil, errors.New("target needs to be either set or rep")
	}

	err = UpdateExercise(db, requestedExercise)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.UpdateExercise")
	}

	// Grab the data from database and return it
	eachExercise, err := GetExerciseByID(db, requestedExercise.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return eachExercise, nil
}

func DeleteBaseExerciseByName(db *gorm.DB, name string) (error) {
	result := db.Where("name = ?", name).Delete(&model.BaseExercise{})

	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil 
}