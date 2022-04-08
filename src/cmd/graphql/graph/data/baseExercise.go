package data

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseExerciseMutation struct {
	db gorm.DB
}

func (b *BaseExerciseMutation) HydrateBaseExercise() {

	data := []model.BaseExercise{
		// Compound Movements
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: "Chest", SpecificParts: model.SpecificPartsFrontDeltoid, Level: 3, AvoidGiven: "Shoulder Pain", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: "Chest", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Chest Pain", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Overhead Press", MuscleGroup: "Shoulders", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Back Pain", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: "Back", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Lower Back Pain", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: "Back", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Convental Deadlifts", MuscleGroup: "Hamstrings", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Romanian Deadlift", MuscleGroup: "Hamstrings", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Sumo Deadlifts", MuscleGroup: "Hamstrings", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Squats", MuscleGroup: "Quads", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Barbell Hip Thrust", MuscleGroup: "Glutes", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "", MovementType: string(model.Compound)},
		{ID: uuid.New().String(), Name: "Wide Squats", MuscleGroup: "Glutes", SpecificParts: model.SpecificParts.FrontDeltoid, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Compound)},
		// {ID: uuid.New().String(), Name: "NEW ADDED", MuscleGroup: "Glutes", Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(models.Compound)},
		// Isolated Movements
		//chest
		{ID: uuid.New().String(), Name: "Chest Flys", MuscleGroup: "Chet", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Incline Bench", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Dips", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Dumbell Pullovers", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		// back
		{ID: uuid.New().String(), Name: "Single Arm Rows", MuscleGroup: "Back", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Cable Rows", MuscleGroup: "Back", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Pull Ups", MuscleGroup: "Back", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Diverging Rows", MuscleGroup: "Back", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Lateral Pull Downs", MuscleGroup: "Back", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		// arms
		{ID: uuid.New().String(), Name: "Bicep Curls", MuscleGroup: "Arms", Level: 1, AvoidGiven: "n/a", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Tricep Extensions", MuscleGroup: "Chest", Level: 3, AvoidGiven: "Shoulders Pain", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Curl Machine", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Hammer Curls", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Skull Crushes", MuscleGroup: "Chest", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		// shoulders
		{ID: uuid.New().String(), Name: "Reverse Pec Deck", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Face Pulls", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Dumbell Lat Raises", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Cable Lat Raises", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Front Raises", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Upright Rows", MuscleGroup: "Deltoid", Level: 1, AvoidGiven: "", MovementType: string(model.Isolated)},
		// legs
		{ID: uuid.New().String(), Name: "Leg Extensions", MuscleGroup: "Quads", Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Hamstring Curls", MuscleGroup: "Hamstrings", Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Leg Press", MuscleGroup: "Arms", Level: 1, AvoidGiven: "n/a", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Hack Squat", MuscleGroup: "Arms", Level: 1, AvoidGiven: "n/a", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Bulgarian Split Squat", MuscleGroup: "Arms", Level: 1, AvoidGiven: "n/a", MovementType: string(model.Isolated)},
		{ID: uuid.New().String(), Name: "Glute-Ham Raises", MuscleGroup: "Glutes", Level: 1, AvoidGiven: "n/a", MovementType: string(model.Isolated)},
	}
	// TODO: refactor badly..
	for _, eachExercise := range data {
		rows, _ := b.db.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()

		defer rows.Close()
		var name, avoidGiven string
		var count int
		for rows.Next() {
			rows.Scan(&name, &avoidGiven)
			if name == eachExercise.Name && avoidGiven == eachExercise.AvoidGiven {
				count += 1
				fmt.Printf("%v ,exists in database!\n", eachExercise.Name)
			}
		}
		if count == 0 {
			b.db.Create(eachExercise)
		}
	}
}
