package postgres

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/packages/api/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (p PgRepo) ValidateUserWorkoutPlan(userEmail string, gymDay string) (string, bool) {
	fmt.Println(gymDay, userEmail)
	var total int64
	p.db.Model(&model.UserWorkoutPlan{}).Where("user_email = ? AND gym_day = ?", userEmail, gymDay).Count(&total)

	if total == 0 {
		return "", false
	}

	var userWorkoutPlanID, userEmailDB, gymDayDB string

	row, err := p.db.Model(&model.UserWorkoutPlan{}).Select("id", "user_email", "gym_day").Rows()
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

func (p PgRepo) AssignUserWorkoutPlan(input model.AddUserWorkoutInput) (*model.UserWorkoutPlan, error) {
	rows, err := p.db.Model(&model.UserWorkoutPlan{}).Select("user_email", "gym_day").Rows()
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
		p.db.Create(userWorkoutPlan)
	}
	return userWorkoutPlan, nil
}

func (p PgRepo) GetUserWorkoutPlansByEmail(ctx context.Context, email string) ([]*model.UserWorkoutPlan, error) {
	userWorkouts := []*model.UserWorkoutPlan{}
	result := p.db.Model(&model.UserWorkoutPlan{}).Where("email = ?", email).Scan(&userWorkouts)

	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to find user workout plans by email")
	}

	return userWorkouts, nil
}

func (p PgRepo) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	userWorkoutPlans := []*model.UserWorkoutPlan{}
	result := p.db.Table("user_workout_plan").Scan(&userWorkoutPlans)
	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(result.Error, "unable to find user workout plans by email")
	}

	return userWorkoutPlans, nil
}
