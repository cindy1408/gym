package graph

import (
	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
)

var BaseExerciseData = []model.BaseExercise{
	// Compound Movements
	{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: "Chest", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Overhead Press", MuscleGroup: "Arms", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: "Back", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Convental Deadlifts", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Romanian Deadlift", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Sumo Deadlifts", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Squats", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Barbell Hip Thrust", MuscleGroup: "Glute", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	{ID: uuid.New().String(), Name: "Wide Squats", MuscleGroup: "Legs", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
	// "Isolated" Movements
	//chest
	{ID: uuid.New().String(), Name: "Chest Flys", MuscleGroup: "Chest", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Incline Bench", MuscleGroup: "Chest", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Dips", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Dumbell Pullovers", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// back
	{ID: uuid.New().String(), Name: "Single Arm Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Cable Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Pull Ups", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Diverging Rows", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Lateral Pull Downs", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// arms
	{ID: uuid.New().String(), Name: "Bicep Curls", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Tricep Extensions", MuscleGroup: "Arms", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Curl Machine", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Hammer Curls", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Skull Crushes", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// shoulders
	{ID: uuid.New().String(), Name: "Reverse Pec Deck", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Face Pulls", MuscleGroup: "Back", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Dumbell Lat Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Cable Lat Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Front Raises", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Upright Rows", MuscleGroup: "Arms", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	// legs
	{ID: uuid.New().String(), Name: "Leg Extensions", MuscleGroup: "Legs", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Hamstring Curls", MuscleGroup: "Legs", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Leg Press", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Hack Squat", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Bulgarian Split Squat", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	{ID: uuid.New().String(), Name: "Glute-Ham Raises", MuscleGroup: "Legs", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
}
