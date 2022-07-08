package data

import (
	"github.com/cindy1408/gym/src/api/graphql/graph/model"
)

var BaseExerciseData = []model.BaseExercise{
	// Compound Movements
	{Name: "Bench Press", MuscleGroup: "Chest", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Overhead Press", MuscleGroup: "Arms", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Bent Over Rows", MuscleGroup: "Back", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Convental Deadlifts", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Romanian Deadlift", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Sumo Deadlifts", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Squats", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Barbell Hip Thrust", MuscleGroup: "Glute", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{Name: "Wide Squats", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	// "Isolated" Movements
	//chest
	{Name: "Chest Flys", MuscleGroup: "Chest", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Incline Bench", MuscleGroup: "Chest", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Dips", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Dumbell Pullovers", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// back
	{Name: "Single Arm Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Cable Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Pull Ups", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Diverging Rows", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Lateral Pull Downs", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// arms
	{Name: "Bicep Curls", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Tricep Extensions", MuscleGroup: "Arms", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Curl Machine", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Hammer Curls", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Skull Crushes", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// shoulders
	{Name: "Reverse Pec Deck", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Face Pulls", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Dumbell Lat Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Cable Lat Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Front Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Upright Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// legs
	{Name: "Leg Extensions", MuscleGroup: "Legs", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Hamstring Curls", MuscleGroup: "Legs", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Leg Press", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Hack Squat", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Bulgarian Split Squat", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{Name: "Glute-Ham Raises", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
}
