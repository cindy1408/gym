package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func AddUserToDB(db *gorm.DB, newUser *model.User) (string, error) {
	rows, err := db.Model(&model.User{}).Select("email").Rows()
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
			db.Model(&model.User{}).First(&model.User{Email: email}).Scan(&existUser)

			return fmt.Sprintf("%v , exists in database!\n", newUser.Email), nil
		}
	}

	if count == 0 {
		db.Create(&newUser)
	}

	return fmt.Sprintf("user %v has been successfully created", newUser.FirstName), nil
}

func ValidateUser(db *gorm.DB, email string) bool {
	rows, err := db.Model(&model.User{}).Select("email").Rows()
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

func UpdateUser(db *gorm.DB, input model.AddUserWorkoutInput, workoutPlanID string) bool {
	db.Model(&model.User{}).Where("email = ?", input.UserEmail).Update("user_workout_plan_id", workoutPlanID)

	return true
}

func GetAllUsers(ctx context.Context, db *gorm.DB) ([]*model.User, error) {
	allUsers := []*model.User{}
	db.Table("users").Scan(&allUsers)
	
	return allUsers, nil 
}

func GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*model.User, error) {
	var userDetails *model.User
	db.Model(&model.User{}).Where("email = ?", email).Scan(&userDetails)

	return userDetails, nil 
}