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
	"github.com/google/uuid"
)

// type MutationResolver struct {
// 	DB *gorm.DB
// }

func (r *mutationResolver) CreateBaseExercise(ctx context.Context, input *model.BaseExerciseInput) (*model.BaseExercise, error) {
	newExercise := model.BaseExercise{
		ID:            uuid.New().String(),
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
				ID:            baseExercise.ID,
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
	for _, eachBaseExercise := range Data {

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

	newUser := &model.User{
		ID:        uuid.New().String(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	_ = r.DB.Model(&model.User{})
	// if err != nil {
	// 	fmt.Printf("%v , selecting database\n", newUser.Email)
	// }
	// defer rows.Close()
	// var email string

	// for rows.Next() {
	// 	rows.Scan(&email)
	// 	if email == newUser.Email {
	// 		fmt.Printf("%v , exists in database!\n", newUser.Email)
	// 	} else {
	// 		r.DB.Create(newUser)
	// 	}
	// }
	return newUser, nil
}

func (r *mutationResolver) AddUserWorkout(ctx context.Context, input model.AddUserWorkoutInput) (*model.WorkoutPerDay, error) {
	id, err := r.Query().GetUserIDByUserEmail(ctx, input.UserEmail)
	if err != nil {
		return nil, fmt.Errorf("error has occurred: ", err)
	}

	for _, eachExercise := range input.Exercises {
		eachExercise := model.EachExercise{
			Name:   eachExercise.Name,
			Weight: eachExercise.Weight,
			Unit:   eachExercise.Unit,
			Sets:   eachExercise.Sets,
			Reps:   eachExercise.Reps,
		}
		r.exercises = append(r.exercises, &eachExercise)
	}

	userWorkoutDay := model.WorkoutPerDay{
		ID:     uuid.New().String(),
		GymDay: input.GymDay,
	}

	r.userWorkoutDays = append(r.userWorkoutDays, &userWorkoutDay)

	userWorkoutPlan := model.UserWorkoutPlan{
		UserID: id,
		Name:   input.GymDay,
		// Cycle:  r.userWorkoutDays,
	}
	r.userWorkoutPlans = append(r.userWorkoutPlans, &userWorkoutPlan)

	return &userWorkoutDay, nil
}

func (r *mutationResolver) IncreaseRep(ctx context.Context, input model.IncreaseRepInput) (*model.UserWorkoutPlan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BaseExercises(ctx context.Context) ([]*model.BaseExercise, error) {
	return r.baseExercises, nil
}

func (r *queryResolver) GetAllAvaliableBaseExercises(ctx context.Context) ([]string, error) {
	var allBaseExerciseNames []string
	for _, eachBaseExercise := range r.baseExercises {
		allBaseExerciseNames = append(allBaseExerciseNames, eachBaseExercise.Name)
	}
	return allBaseExerciseNames, nil
}

func (r *queryResolver) GetAllEachExercise(ctx context.Context) ([]*model.EachExercise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllWorkoutDay(ctx context.Context) ([]*model.WorkoutPerDay, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserWorkoutPlansByEmail(ctx context.Context, input string) (*model.UserWorkoutPlan, error) {
	var userID string
	var userWorkoutPlanResult model.UserWorkoutPlan
	if input == "" {
		return nil, fmt.Errorf("empty user email")
	}
	for _, user := range r.users {
		if user.Email == input {
			userID = user.ID
		}
	}
	for _, workoutPlan := range r.userWorkoutPlans {
		if userID == workoutPlan.UserID {
			userWorkoutPlanResult = model.UserWorkoutPlan{
				UserID: workoutPlan.UserID,
				Name:   workoutPlan.Name,
				// Cycle:  workoutPlan.Cycle,
			}
		}
	}
	return &userWorkoutPlanResult, nil
}

func (r *queryResolver) GetAllUserWorkoutPlans(ctx context.Context) ([]*model.UserWorkoutPlan, error) {
	return r.userWorkoutPlans, nil
}

func (r *queryResolver) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserIDByUserEmail(ctx context.Context, input string) (string, error) {
	if input == "" {
		return "", errors.New("you must insert an email address")
	}
	for _, user := range r.users {
		if user.Email == input {
			return user.ID, nil
		}
	}
	return "", errors.New("please enter a valid email")
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
