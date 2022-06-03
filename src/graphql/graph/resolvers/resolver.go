package resolvers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// used to pull the interface in
type Resolver struct {
	DB               *gorm.DB
	baseExercises    []*model.BaseExercise
	// users            []*model.User
	// exercises        []*model.EachExercise
	muscleGroups     []*model.MuscleGroup
	userWorkoutPlans []*model.UserWorkoutPlan
}

func (r *Resolver) Init() error {
	err := r.DB.AutoMigrate(
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

	return r.DB.Error
}

func Hasher(toHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(toHash))
	return hex.EncodeToString(hasher.Sum(nil))
}
