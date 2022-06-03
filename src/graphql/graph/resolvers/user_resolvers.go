package resolvers

import (
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
)

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
