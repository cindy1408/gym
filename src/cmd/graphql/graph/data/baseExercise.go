package data

import (
	"fmt"

	"github.com/cindy1408/gym/src/cmd/graphql/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseExerciseMutation struct {
	db gorm.DB
}

func (b *BaseExerciseMutation) HydrateBaseExercise() {
	data := []model.BaseExercise{
		// Compound Movements
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"}, 
		{ID: uuid.New().String(), Name: "Overhead Press", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Convental Deadlifts", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Romanian Deadlift", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Sumo Deadlifts", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Squats", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Barbell Hip Thrust", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Wide Squats", MuscleGroup: []*model.MuscleGroup{}, SpecificParts: []*model.Body{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Compound"},
		// {ID: uuid.New().String(), Name: "NEW ADDED", MuscleGroup: []*model.MuscleGroup{}, Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(models.Compound)},
		// "Isolated" Movements
		//chest
		{ID: uuid.New().String(), Name: "Chest Flys", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Incline Bench", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dips", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dumbell Pullovers", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		// back
		{ID: uuid.New().String(), Name: "Single Arm Rows", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Cable Rows", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Pull Ups", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Diverging Rows", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Lateral Pull Downs", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		// arms
		{ID: uuid.New().String(), Name: "Bicep Curls", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Tricep Extensions", MuscleGroup: []*model.MuscleGroup{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Curl Machine", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hammer Curls", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Skull Crushes", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		// shoulders
		{ID: uuid.New().String(), Name: "Reverse Pec Deck", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Face Pulls", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dumbell Lat Raises", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Cable Lat Raises", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Front Raises", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Upright Rows", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		// legs
		{ID: uuid.New().String(), Name: "Leg Extensions", MuscleGroup: []*model.MuscleGroup{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hamstring Curls", MuscleGroup: []*model.MuscleGroup{}, Level: 3, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Leg Press", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hack Squat", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Bulgarian Split Squat", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Glute-Ham Raises", MuscleGroup: []*model.MuscleGroup{}, Level: 1, AvoidGiven: []*model.AvoidGiven{}, MovementType: "Isolated"},
	}

	// TODO: refactor badly..
	for _, eachExercise := range data {
		rows, _ := b.db.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()

		defer rows.Close()
		var name string
		var avoidGiven []*model.AvoidGiven
		var count int
		for rows.Next() {
			rows.Scan(&name, &avoidGiven)
			if name == eachExercise.Name {
				count += 1
				fmt.Printf("%v ,exists in database!\n", eachExercise.Name)
			}
		}
		if count == 0 {
			b.db.Create(eachExercise)
		}
	}
}
