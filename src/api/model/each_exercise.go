package model

type Details string 

var (
	Set Details 
	Rep Details 
)

type EachExercise struct {
	ID                string `json:"id"`
	UserWorkoutPlanID string `json:"userWorkoutPlanID"`
	Name              string `json:"name"`
	Weight            int    `json:"weight"`
	Unit              string `json:"unit"`
	Sets              int    `json:"sets"`
	Reps              int    `json:"reps"`
}

type AddExerciseInput struct {
	UserEmail    string               `json:"userEmail"`
	GymDay       string               `json:"gymDay"`
	EachExercise []*EachExerciseInput `json:"eachExercise"`
}

type UpdateExerciseInput struct {
	UserEmail    string             `json:"userEmail"`
	GymDay       string             `json:"gymDay"`
	EachExercise *EachExerciseInput `json:"eachExercise"`
}

type EachExerciseInput struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Unit   string `json:"unit"`
	Sets   int    `json:"sets"`
	Reps   int    `json:"reps"`
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

type MuscleGroup struct {
	Name string `json:"name"`
}

type SpecificParts struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscleGroup"`
}

type AvoidGiven struct {
	Name string `json:"name"`
}
