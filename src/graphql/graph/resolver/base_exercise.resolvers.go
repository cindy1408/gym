package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"gorm.io/gorm"
)

func (r *BaseExerciseResolver) GetBaseExerciseByName(ctx context.Context, input string) (*model.BaseExercise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *BaseExerciseResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	allBaseExercises := []*model.BaseExercise{}
	r.DB.Table("base_exercises").Scan(&allBaseExercises)
	return allBaseExercises, nil
}

func (r *BaseExerciseResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}

	r.DB.Debug().Model(&model.BaseExercise{}).Where("name = ?", input.Name).Updates(updatedExercise)

	return &updatedExercise, nil
}

func (r *BaseExerciseResolver) HydrateBaseExercise(ctx context.Context) (string, error) {
	for _, eachBaseExercise := range graph.BaseExerciseData {
		rows, err := r.DB.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()
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
			r.DB.Create(eachBaseExercise)
		}
	}

	return "Base exercise table has been hydrated!", nil
}

func (r *BaseExerciseResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (string, error) {
	if input.Name == "" ||
		input.MuscleGroup == "" ||
		input.SpecificParts == "" ||
		input.Level == 0 ||
		input.MovementType == "" {
		return "at least one of the required field is missing", nil
	}
	newExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}
	r.baseExercises = append(r.baseExercises, &newExercise)

	return fmt.Sprintf("base exercise %v has been added", input.Name), nil
}

type BaseExerciseResolver struct {
	DB            *gorm.DB
	baseExercises []*model.BaseExercise
}

func NewBaseExerciseResolver(db *gorm.DB) *BaseExerciseResolver {
	return &BaseExerciseResolver{
		DB: db,
	}
}

// Query returns generated.QueryResolver implementation.
// func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// type queryResolver struct{ *Resolver }
