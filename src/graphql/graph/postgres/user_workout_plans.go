package postgres

import (
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ValidateUserWorkoutPlan(db *gorm.DB, userEmail string, gymDay string) (string, bool) {
	fmt.Println(gymDay, userEmail)
	var total int64
	db.Model(&model.UserWorkoutPlan{}).Where("user_email = ? AND gym_day = ?", userEmail, gymDay).Count(&total)

	if total == 0 {
		return "", false
	}

	var userWorkoutPlanID, userEmailDB, gymDayDB string

	row, err := db.Model(&model.UserWorkoutPlan{}).Select("id", "user_email", "gym_day").Rows()
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

func AssignUserWorkoutPlan(db *gorm.DB, input model.AddUserWorkoutInput) (*model.UserWorkoutPlan, error) {
	rows, err := db.Model(&model.UserWorkoutPlan{}).Select("user_email", "gym_day").Rows()
	if err != nil {
		return nil, errors.Wrapf(err, "error with user workout plan")
	}
	defer rows.Close()

	var existingUserEmail string
	var existingUserGymDay string
	var userWorkoutPlan *model.UserWorkoutPlan
	var count int

	for rows.Next() {
		rows.Scan(&existingUserEmail, &existingUserGymDay)
		if existingUserEmail == input.UserEmail && existingUserGymDay == input.GymDay {
			count++
			return nil, nil
		}
	}

	if count == 0 {
		userWorkoutPlan = &model.UserWorkoutPlan{
			ID:        uuid.New().String(),
			UserEmail: input.UserEmail,
			GymDay:    input.GymDay,
		}
		db.Create(userWorkoutPlan)
	}
	return userWorkoutPlan, nil
}
