package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/data"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
)

type BaseExercises interface {
	UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error)
	HydrateBaseExercise(ctx context.Context) (string, error)
	AddBaseExercise(ctx context.Context, baseExercise *model.BaseExerciseInput) (string, error)
	ValidateBaseExercise(ctx context.Context, name string) bool
	GetAllBaseExercise(ctx context.Context) ([]*model.BaseExercise, error)
	GetBaseExerciseByName(ctx context.Context, name string) (*model.BaseExercise, error)
	Increase(ctx context.Context, input model.IncreaseInput, target model.Details) (*model.EachExercise, error)
	DeleteBaseExerciseByName(ctx context.Context, name string) error
}

func (p PgRepo) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}

	result := p.db.Debug().Model(&model.BaseExercise{}).Where("name = ?", input.Name).Updates(updatedExercise)

	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to update base exercise")
	}

	return &updatedExercise, nil
}

func (p PgRepo) HydrateBaseExercise(ctx context.Context) (string, error) {
	for _, eachBaseExercise := range data.BaseExerciseData {
		// TODO: ISSUE HERE
		rows, err := p.db.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()
		fmt.Println("here3")
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
			p.db.Create(eachBaseExercise)
		}
	}
	
	return "Base exercise table has been hydrated!", nil
}

func (p PgRepo) AddBaseExercise(ctx context.Context, baseExercise *model.BaseExerciseInput) (string, error) {

	p.db.Create(model.BaseExercise{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	})

	return fmt.Sprintf("base exercise %v has been added", baseExercise.Name), nil
}

func (p PgRepo) ValidateBaseExercise(ctx context.Context, name string) bool {
	rows, err := p.db.Model(&model.BaseExercise{}).Select("name").Rows()
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

func (p PgRepo) GetAllBaseExercise(ctx context.Context) ([]*model.BaseExercise, error) {
	allBaseExercises := []*model.BaseExercise{}
	fmt.Println("here")
	// TODO: ISSUE IS HERE
	p.db.Table("base_exercises").Scan(&allBaseExercises)
	return allBaseExercises, nil
}

func (p PgRepo) GetBaseExerciseByName(ctx context.Context, name string) (*model.BaseExercise, error) {
	if name == "" {
		return nil, errors.Wrapf(nil, "base exercise name is empty")
	}
	var baseExercise *model.BaseExercise
	result := p.db.Where("name", name).Find(&model.BaseExercise{}).Scan(&baseExercise)

	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to find base exercise")
	}

	return baseExercise, nil
}

func (p PgRepo) Increase(ctx context.Context, input model.IncreaseInput, target model.Details) (*model.EachExercise, error) {
	userDetails, err := p.GetUserByEmail(ctx, input.UserEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.GetUserByEmail")
	}

	requestedExercise, err := p.GetExerciseByNameAndWorkoutPlanID(input.ExerciseName, *userDetails.UserWorkoutPlanID)
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

	err = p.UpdateExercise(requestedExercise)
	if err != nil {
		return nil, errors.Wrapf(err, "postgres.UpdateExercise")
	}

	// Grab the data from database and return it
	eachExercise, err := p.GetExerciseByID(requestedExercise.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "r.increase")
	}

	return eachExercise, nil
}

func (p PgRepo) DeleteBaseExerciseByName(name string) error {
	result := p.db.Where("name = ?", name).Delete(&model.BaseExercise{})

	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
