package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *mutationResolver) AddExercise(ctx context.Context, input *model.AddExerciseInput) (string, error) {
	validated := r.ValidateUser(input.UserEmail)

	if !validated {
		return "Please create an account first", nil
	}

	// validate user workout plan
	workoutPlanID, validated := r.ValidateUserWorkoutPlan(input.UserEmail, input.GymDay)
	if !validated {
		return "You need to create a workout plan", nil
	}

	for _, eachExercise := range input.EachExercise {
		addUserExercise := model.EachExercise{
			ID:                uuid.NewString(),
			UserWorkoutPlanID: workoutPlanID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Unit:              eachExercise.Unit,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}
		r.DB.Create(&addUserExercise)
	}

	return "Your account has been successfully created", nil
}

func (r *mutationResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (string, error) {
	if input.Name == "" ||
		input.MuscleGroup == "" ||
		input.SpecificParts == "" ||
		input.Level == 0 ||
		input.MovementType == "" {
		return "at least one of the required field is missing", nil
	}
	newExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}
	r.baseExercises = append(r.baseExercises, &newExercise)

	return fmt.Sprintf("base exercise %v has been added", input.Name), nil
}

func (r *mutationResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}

	r.DB.Debug().Model(&model.BaseExercise{}).Where("name = ?", input.Name).Updates(updatedExercise)

	return &updatedExercise, nil
}

func (r *mutationResolver) HydrateBaseExercise(ctx context.Context) (string, error) {
	for _, eachBaseExercise := range graph.BaseExerciseData {
		rows, err := r.DB.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()
		if err != nil {
			fmt.Printf("%v , selecting database\n", eachBaseExercise.Name)
		}
		defer rows.Close()

		var name, avoidGiven string
		var count int

		for rows.Next() {
			rows.Scan(&name, &avoidGiven)
			if avoidGiven == "" {
				if name == eachBaseExercise.Name {
					count += 1
					fmt.Printf("%v , exists in database!\n", eachBaseExercise.Name)
				}
			} else if name == eachBaseExercise.Name && &avoidGiven == eachBaseExercise.AvoidGiven {
				fmt.Printf("%v , exists in database!\n", eachBaseExercise.Name)
			}
		}
		if count == 0 {
			r.DB.Create(eachBaseExercise)
		}
	}

	return "Base exercise table has been hydrated!", nil
}

func (r *mutationResolver) HydrateMuscleGroups(ctx context.Context) (string, error) {
	for _, eachMuscleGroup := range graph.MuscleGroupData {
		rows, err := r.DB.Model(&model.MuscleGroup{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachMuscleGroup.Name)
		}
		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachMuscleGroup.Name {
				count += 1
				fmt.Printf("%v , exists in database!\n", eachMuscleGroup.Name)
			}
		}
		if count == 0 {
			muscleGroup := model.MuscleGroup{
				Name: eachMuscleGroup.Name,
			}
			r.DB.Create(muscleGroup)
		}
	}
	return "Muscle group table has been hydrated", nil
}

func (r *mutationResolver) HydrateSpecificParts(ctx context.Context) (string, error) {
	for _, eachSpecificMuscleGroup := range graph.SpecificMuscleGroupData {
		rows, err := r.DB.Model(&model.SpecificParts{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachSpecificMuscleGroup.Name)
		}
		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachSpecificMuscleGroup.Name {
				fmt.Printf("%v, exists in database!\n", eachSpecificMuscleGroup.Name)
			}
		}
		if count == 0 {
			specificMuscleGroup := model.SpecificParts{
				Name:        eachSpecificMuscleGroup.Name,
				MuscleGroup: eachSpecificMuscleGroup.MuscleGroup,
			}
			r.DB.Create(specificMuscleGroup)
		}
	}
	return "Specific Parts has been hydrated", nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (string, error) {
	if input.FirstName == "" {
		return "first name is missing", nil
	}
	if input.LastName == "" {
		return "last name is missing", nil
	}
	if input.Password == "" {
		return "password is missing", nil
	}

	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return "invalid email address", nil
	}

	hashedPw := Hasher(input.Password)

	newUser := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPw,
	}

	result, err := r.AddUserToDB(&newUser)
	if err != nil {
		return "", errors.Wrapf(err, "r.AddUserToDB, %v")
	}

	return result, nil
}

func (r *mutationResolver) AddUserWorkout(ctx context.Context, input model.AddUserWorkoutInput) (string, error) {
	// check if the user email exists!
	validated := r.ValidateUser(input.UserEmail)

	if !validated {
		return "You need to sign up first", nil
	}

	// assign user email to userworkout plan
	rows, err := r.DB.Model(&model.UserWorkoutPlan{}).Select("user_email", "gym_day").Rows()
	if err != nil {
		return "", errors.Wrapf(err, "error with user workout plan %v")
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
			return fmt.Sprintf("User %v already has %v, please add exercise instead", existingUserEmail, existingUserGymDay), nil
		}
	}

	if count == 0 {
		userWorkoutPlan = &model.UserWorkoutPlan{
			ID:        uuid.New().String(),
			UserEmail: input.UserEmail,
			GymDay:    input.GymDay,
		}
		r.DB.Create(userWorkoutPlan)

		// update user table
		r.DB.Model(&model.User{}).Where("email = ?", input.UserEmail).Update("user_workout_plan_id", userWorkoutPlan.ID)
	}

	if input.Exercises != nil {
		return r.AddEachExercisesToDB(userWorkoutPlan.ID, input.Exercises)
	}

	return fmt.Sprintf("User %v has been added", input.UserEmail), nil
}

func (r *mutationResolver) IncreaseRep(ctx context.Context, input model.IncreaseRepInput) (*model.UserWorkoutPlan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	return r.baseExercises, nil
}

func (r *queryResolver) GetAllAvailableBaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	allBaseExercises := []*model.BaseExercise{}
	r.DB.Table("base_exercises").Scan(&allBaseExercises)
	return allBaseExercises, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserWorkoutPlansByEmail(ctx context.Context, input string) (*model.UserWorkoutPlan, error) {
	return nil, nil
}

func (r *queryResolver) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	return r.userWorkoutPlans, nil
}

func (r *queryResolver) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	allUsers := []*model.User{}
	r.DB.Table("users").Scan(&allUsers)

	return allUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
