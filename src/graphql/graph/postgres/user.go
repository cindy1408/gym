package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
)

func (p PgRepo) AddUserToDB(newUser *model.User) (string, error) {
	rows, err := p.db.Model(&model.User{}).Select("email").Rows()
	if err != nil {
		return "issue with user database", errors.Wrapf(err, "model.User{}")
	}

	defer rows.Close()

	var email string
	var count int
	for rows.Next() {
		rows.Scan(&email)
		if email == newUser.Email {
			count++

			var existUser model.User
			p.db.Model(&model.User{}).First(&model.User{Email: email}).Scan(&existUser)

			return fmt.Sprintf("%v , exists in database!\n", newUser.Email), nil
		}
	}

	if count == 0 {
		p.db.Create(&newUser)
	}

	return fmt.Sprintf("user %v has been successfully created", newUser.FirstName), nil
}

func (p PgRepo) ValidateUser(email string) bool {
	rows, err := p.db.Model(&model.User{}).Select("email").Rows()
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

func (p PgRepo) UpdateUser(input model.AddUserWorkoutInput, workoutPlanID string) bool {
	p.db.Model(&model.User{}).Where("email = ?", input.UserEmail).Update("user_workout_plan_id", workoutPlanID)

	return true
}

func (p PgRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	allUsers := []*model.User{}
	p.db.Table("users").Scan(&allUsers)

	return allUsers, nil
}

func (p PgRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var userDetails *model.User
	p.db.Model(&model.User{}).Where("email = ?", email).Scan(&userDetails)

	return userDetails, nil
}
