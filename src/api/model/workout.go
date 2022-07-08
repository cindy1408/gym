package model

type AddUserWorkoutInput struct {
	UserEmail string               `json:"userEmail"`
	GymDay    string               `json:"gymDay"`
	Exercises []*EachExerciseInput `json:"exercises"`
}

// user workout plan consists of userID, workoutID and name of the day eg. Push day
type UserWorkoutPlan struct {
	Id        string `json:"id"`
	UserEmail string `json:"userEmail"`
	GymDay    string `json:"gymDay"`
}
