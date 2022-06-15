package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/resolver"
	"github.com/pkg/errors"
)

var DB resolver.Resolver

func (r DB) AddUserToDB(ctx context.Context, newUser *model.User) (string, error) {
	rows, err := r.DB.Model(&model.User{}).Select("email").Rows()
	// rows, err := r.DB.Model(&model.User{}).Select("email").Rows()
	if err != nil {
		return "issue with user database", errors.Wrapf(err, "model.User{}, %v")
	}
	defer rows.Close()

	var email string
	var count int
	for rows.Next() {
		rows.Scan(&email)
		if email == newUser.Email {
			count++

			var existUser model.User
			r.DB.Model(&model.User{}).First(&model.User{Email: email}).Scan(&existUser)

			return fmt.Sprintf("%v , exists in database!\n", newUser.Email), nil
		}
	}

	if count == 0 {
		r.DB.Create(&newUser)
	}
	return fmt.Sprintf("user %v has been successfully created", newUser.FirstName), nil
}

func (r *Resolver) ValidateUser(email string) bool {
	rows, err := r.DB.Model(&model.User{}).Select("email").Rows()
	if err != nil {
		fmt.Println("issue with user table")
	}
	defer rows.Close()

	var existingEmail string
	var exists int
	for rows.Next() {
		exists++
		rows.Scan(&existingEmail)
		if existingEmail == email {
			return true
		}
	}

	return exists != 0
}

func (r *Resolver) UpdateUser(input model.AddUserWorkoutInput, workoutPlanID string) bool {
	r.DB.Model(&model.User{}).Where("email = ?", input.UserEmail).Update("user_workout_plan_id", workoutPlanID)

	return true
}
