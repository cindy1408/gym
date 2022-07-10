package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/packages/api/graphql/graph/model"
	"github.com/cindy1408/gym/packages/data"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo interface {
	GetBaseExerciseByName(
		ctx context.Context,
		name string) (*model.BaseExercise, error)
	GetAllBaseExercise(
		ctx context.Context) ([]*model.BaseExercise, error)
	ValidateBaseExercise(
		ctx context.Context,
		name string)
	AddBaseExercise(
		ctx context.Context,
		baseExercise *model.BaseExerciseInput)
	HydrateBaseExercise(
		ctx context.Context) (string, error)
	UpdateBaseExercise(
		ctx context.Context,
		input *model.BaseExerciseInput) (*model.BaseExercise, error)
	Increase(
		ctx context.Context,
		input model.IncreaseInput,
		target string) (*model.EachExercise, error)
	DeleteBaseExerciseByName(
		name string) error
	HydrateMuscleGroups(ctx context.Context) error
}

// PgRepo is the postgres implementation of postgres.Repo
type PgRepo struct {
	db *gorm.DB
}

// NewRepo creates a new PgRepo
func NewRepo(DB *gorm.DB) PgRepo {
	return PgRepo{
		db: DB,
	}
}

func (p *PgRepo) Init() error {
	err := p.db.AutoMigrate(
		&model.BaseExercise{},
		&model.User{},
		&model.MuscleGroup{},
		&model.SpecificParts{},
		&model.UserWorkoutPlan{},
		&model.EachExercise{},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewDatabase() (*gorm.DB, error) {
	databaseURL := "host=localhost user=postgres password=password dbname=gym port=5432 sslmode=disable"
	return gorm.Open(postgres.New(postgres.Config{
		DSN: databaseURL,
	}), &gorm.Config{})
}

func (p PgRepo) HydrateMuscleGroups(ctx context.Context) error {
	for _, eachMuscleGroup := range data.MuscleGroupData {
		rows, err := p.db.Model(&model.MuscleGroup{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachMuscleGroup.Name)
		}
		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachMuscleGroup.Name {
				count += 1
				fmt.Printf("%v , exists in database!\n", eachMuscleGroup.Name)
			}
		}

		if count == 0 {
			p.db.Create(eachMuscleGroup)
		}
	}
	
	return nil
}

func (p PgRepo) HydrateSpecificParts(ctx context.Context) error {
	for _, eachSpecificMuscleGroup := range data.SpecificMuscleGroupData {
		rows, err := p.db.Model(&model.SpecificParts{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachSpecificMuscleGroup.Name)
		}
		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachSpecificMuscleGroup.Name {
				count += 1
				fmt.Printf("%v, exists in database!\n", eachSpecificMuscleGroup.Name)
			}
		}

		if count == 0 {
			p.db.Create(eachSpecificMuscleGroup)
		} else {
			return errors.New("HydrateSpecificParts")
		}
	}

	return nil
}

func (p PgRepo) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) ([]string, error) {
	allSpecifics := []string{}
	p.db.Model(&model.SpecificParts{}).Where("name = ?", input.Name).Scan(&allSpecifics)

	return allSpecifics, nil
}

