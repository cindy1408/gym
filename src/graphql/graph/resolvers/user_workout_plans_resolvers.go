package resolvers

import (
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
)

func (r *mutationResolver) ValidateUserWorkoutPlan(userEmail string, gymDay string) (string, bool) {

	fmt.Println(gymDay, userEmail)
	var total int64
	r.DB.Model(&model.UserWorkoutPlan{}).Where("user_email = ? AND gym_day = ?", userEmail, gymDay).Count(&total)

	if total == 0 {
		return "", false
	}

	var userWorkoutPlanID, userEmailDB, gymDayDB string

	row, err := r.DB.Model(&model.UserWorkoutPlan{}).Select("id", "user_email", "gym_day").Rows()
	if err != nil {
		fmt.Println("issue with user workout plan database")
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&userWorkoutPlanID, &userEmailDB, &gymDayDB)
		if userEmail == userEmailDB && gymDay == gymDayDB {
			return userWorkoutPlanID, true
		}
	}

	return "", false
}
