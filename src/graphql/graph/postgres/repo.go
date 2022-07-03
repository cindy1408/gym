package postgres

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

type Repo interface {
	GetBaseExerciseByName(
		ctx context.Context, 
		db *gorm.DB, 
		name string) (*model.BaseExercise, error)
	GetAllBaseExercise(
		ctx context.Context, 
		db *gorm.DB) ([]*model.BaseExercise, error)
	ValidateBaseExercise(
		ctx context.Context, 
		db *gorm.DB, 
		name string)
	AddBaseExercise(
		ctx context.Context, 
		db *gorm.DB, 
		baseExercise *model.BaseExerciseInput)
	HydrateBaseExercise(
		ctx context.Context, 
		db *gorm.DB) (string, error)
	UpdateBaseExercise(
		ctx context.Context, 
		db *gorm.DB, 
		input *model.BaseExerciseInput) (*model.BaseExercise, error)
	Increase(
		ctx context.Context, 
		db *gorm.DB, 
		input model.IncreaseInput, 
		target string) (*model.EachExercise, error)
	DeleteBaseExerciseByName(
		db *gorm.DB, 
		name string) (error)
}

// PgRepo is the postgres implementation of postgres.Repo
type PgRepo struct {
	pgClient *gorm.DB
}

// NewRepo creates a new PgRepo
func NewRepo(pgClient *gorm.DB) PgRepo {
	return PgRepo{
		pgClient: pgClient,
	}
}