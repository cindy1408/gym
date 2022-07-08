package model 

type BaseExercise struct {
	Name          string  `json:"name"`
	MuscleGroup   string  `json:"muscleGroup"`
	SpecificParts string  `json:"specificParts"`
	Level         int     `json:"level"`
	AvoidGiven    *string `json:"avoidGiven"`
	MovementType  string  `json:"movementType"`
}

type BaseExerciseInput struct {
	Name          string  `json:"name"`
	MuscleGroup   string  `json:"muscleGroup"`
	SpecificParts string  `json:"specificParts"`
	Level         int     `json:"level"`
	AvoidGiven    *string `json:"avoidGiven"`
	MovementType  string  `json:"movementType"`
}
