package resolvers

import (
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/pkg/errors"
)

func (r mutationResolver) AddUserToDB(newUser *model.User) (string, error) {
	rows, err := r.DB.Model(&model.User{}).Select("email").Rows()
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

func (r *mutationResolver) ValidateUser(email string) bool {
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
