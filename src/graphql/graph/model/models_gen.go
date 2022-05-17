// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddUserWorkoutInput struct {
	UserEmail string               `json:"userEmail"`
	GymDay    string               `json:"gymDay"`
	Exercises []*EachExerciseInput `json:"exercises"`
}

type AvoidGiven struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BaseExercise struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	MuscleGroup   string  `json:"muscleGroup"`
	SpecificParts string  `json:"specificParts"`
	Level         int     `json:"level"`
	AvoidGiven    *string `json:"avoidGiven"`
	MovementType  string  `json:"movementType"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type EachExercise struct {
	ID              string `json:"id"`
	WorkoutPerDayID string `json:"workoutPerDayID"`
	Name            string `json:"name"`
	Weight          int    `json:"weight"`
	Unit            string `json:"Unit"`
	Sets            int    `json:"Sets"`
	Reps            int    `json:"Reps"`
}

type EachExerciseInput struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Unit   string `json:"Unit"`
	Sets   int    `json:"Sets"`
	Reps   int    `json:"Reps"`
}

type MuscleGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SpecificParts struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID                string  `json:"id"`
	FirstName         string  `json:"firstName"`
	LastName          string  `json:"lastName"`
	Email             string  `json:"email"`
	Password          string  `json:"password"`
	UserWorkoutPlanID *string `json:"UserWorkoutPlanId"`
}

type UserWorkoutPlan struct {
	UserID          string  `json:"userId"`
	Name            string  `json:"name"`
	WorkoutPerDayID *string `json:"workoutPerDayId"`
}

type WorkoutPerDay struct {
	ID         string `json:"id"`
	GymDay     string `json:"gymDay"`
	ExerciseID string `json:"exerciseID"`
}

type BaseExerciseInput struct {
	Name          string  `json:"name"`
	MuscleGroup   string  `json:"muscleGroup"`
	SpecificParts string  `json:"specificParts"`
	Level         int     `json:"level"`
	AvoidGiven    *string `json:"avoidGiven"`
	MovementType  string  `json:"movementType"`
}

type IncreaseRepInput struct {
	UserID       string `json:"userID"`
	GymDay       string `json:"gymDay"`
	ExerciseName string `json:"exerciseName"`
}

type MuscleSpecificInput struct {
	Name string `json:"name"`
}
