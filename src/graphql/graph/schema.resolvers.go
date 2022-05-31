package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/model"
)

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
	for _, eachBaseExercise := range BaseExerciseData {

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
	for _, eachMuscleGroup := range MuscleGroupData {
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
	for _, eachSpecificMuscleGroup := range SpecificMuscleGroupData {
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

	newUser := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
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
	// TODO:
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
