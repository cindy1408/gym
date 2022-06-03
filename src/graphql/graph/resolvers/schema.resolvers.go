package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) AddExercise(ctx context.Context, input *model.AddExerciseInput) (*model.EachExercise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	newExercise := model.BaseExercise{
		Name:          input.Name,
		MuscleGroup:   input.MuscleGroup,
		SpecificParts: input.SpecificParts,
		Level:         input.Level,
		AvoidGiven:    input.AvoidGiven,
		MovementType:  input.MovementType,
	}
	r.baseExercises = append(r.baseExercises, &newExercise)

	return &newExercise, nil
}

func (r *mutationResolver) UpdateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	updatedExercise := model.BaseExercise{}

	for i, baseExercise := range r.baseExercises {
		if input.Name == baseExercise.Name {
			updatedExercise = model.BaseExercise{
				Name:          input.Name,
				MuscleGroup:   input.MuscleGroup,
				SpecificParts: input.SpecificParts,
				Level:         input.Level,
				AvoidGiven:    input.AvoidGiven,
				MovementType:  input.MovementType,
			}
			r.baseExercises[i] = &updatedExercise
			return &updatedExercise, nil
		}
	}
	return nil, errors.New("unable to find exercise name in database")
}

func (r *mutationResolver) HydrateBaseExercise(ctx context.Context) ([]*model.BaseExercise, error) {
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

	return r.baseExercises, nil
}

func (r *mutationResolver) HydrateMuscleGroups(ctx context.Context) ([]*model.MuscleGroup, error) {
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
	return r.muscleGroups, nil
}

func (r *mutationResolver) HydrateSpecificParts(ctx context.Context) ([]*model.SpecificParts, error) {
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
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	if input.FirstName == "" {
		return nil, errors.New("first name is missing")
	}
	if input.LastName == "" {
		return nil, errors.New("last name is missing")
	}
	if input.Password == "" {
		return nil, errors.New("password is missing")
	}
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}

	hashedPw := Hasher(input.Password)

	newUser := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPw,
	}

	rows, err := r.DB.Model(&model.User{}).Select("email").Rows()
	if err != nil {
		fmt.Printf("%v , selecting database\n", newUser.Email)
	}
	defer rows.Close()

	var email string
	var count int
	for rows.Next() {
		rows.Scan(&email)
		if email == newUser.Email {
			fmt.Printf("%v , exists in database!\n", newUser.Email)
			count++

			var existUser model.User
			r.DB.Model(&model.User{}).First(&model.User{Email: email}).Scan(&existUser)

			return &existUser, nil
		}
	}

	if count == 0 {
		r.DB.Create(&newUser)
	}

	return &newUser, nil
}

func (r *mutationResolver) AddUserWorkout(ctx context.Context, input model.AddUserWorkoutInput) (*model.UserWorkoutPlan, error) {
	// check if the user email exists!
	validated := r.ValidateUser(input.UserEmail)

	if !validated {
		fmt.Println("You need to sign up first")
		return nil, nil
	}

	// assign user email to userworkout plan
	rows, err := r.DB.Model(&model.UserWorkoutPlan{}).Select("user_email", "gym_day").Rows()
	if err != nil {
		fmt.Println("error with user workout plan")
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
			fmt.Println("User already has this gym day, please add exercise instead")
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

	if input.Exercises == nil {
		return userWorkoutPlan, nil
	}

	// if exercise exists, add exercise in each_exercise table
	for _, eachExercise := range input.Exercises {
		addExercise := &model.EachExercise{
			ID:                uuid.New().String(),
			UserWorkoutPlanID: userWorkoutPlan.ID,
			Name:              eachExercise.Name,
			Weight:            eachExercise.Weight,
			Sets:              eachExercise.Sets,
			Reps:              eachExercise.Reps,
		}

		r.DB.Create(&addExercise)
	}

	return nil, nil
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
