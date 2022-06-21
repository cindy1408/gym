// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddUserWorkoutInput struct {
	UserEmail string               `json:"userEmail"`
	GymDay    string               `json:"gymDay"`
	Exercises []*EachExerciseInput `json:"exercises"`
}

type AvoidGiven struct {
	Name string `json:"name"`
}

type BaseExercise struct {
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
	ID                string `json:"id"`
	UserWorkoutPlanID string `json:"userWorkoutPlanID"`
	Name              string `json:"name"`
	Weight            int    `json:"weight"`
	Unit              string `json:"Unit"`
	Sets              int    `json:"Sets"`
	Reps              int    `json:"Reps"`
}

type EachExerciseInput struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Unit   string `json:"Unit"`
	Sets   int    `json:"Sets"`
	Reps   int    `json:"Reps"`
}

type MuscleGroup struct {
	Name string `json:"name"`
}

type SpecificParts struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscleGroup"`
}

type User struct {
	FirstName         string  `json:"firstName"`
	LastName          string  `json:"lastName"`
	Email             string  `json:"email"`
	Password          string  `json:"password"`
	UserWorkoutPlanID *string `json:"UserWorkoutPlanId"`
}

type UserWorkoutPlan struct {
	ID        string `json:"id"`
	UserEmail string `json:"userEmail"`
	GymDay    string `json:"gymDay"`
}

type AddExerciseInput struct {
	UserEmail    string               `json:"userEmail"`
	GymDay       string               `json:"gymDay"`
	EachExercise []*EachExerciseInput `json:"EachExercise"`
}

type BaseExerciseInput struct {
	Name          string  `json:"name"`
	MuscleGroup   string  `json:"muscleGroup"`
	SpecificParts string  `json:"specificParts"`
	Level         int     `json:"level"`
	AvoidGiven    *string `json:"avoidGiven"`
	MovementType  string  `json:"movementType"`
}

type IncreaseInput struct {
	UserEmail    string `json:"userEmail"`
	GymDay       string `json:"gymDay"`
	ExerciseName string `json:"exerciseName"`
}

type MuscleSpecificInput struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"MuscleGroup"`
}

type UpdateExerciseInput struct {
	UserEmail string `json:"userEmail"`
	GymDay    string `json:"gymDay"`
}
