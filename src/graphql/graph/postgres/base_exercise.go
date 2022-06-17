package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/model"
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

	db.Debug().Model(&model.BaseExercise{}).Where("name = ?", input.Name).Updates(updatedExercise)

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

func ValidateBaseExercise(ctx context.Context, db *gorm.DB, name string) (bool, error) {
	// TODO! 
	return false, nil 
}