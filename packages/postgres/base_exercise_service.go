package postgres

import (
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BaseExercise struct {
	db      gorm.DB
	Message string
	// everything in this struct is avaliable to Hydrate function
}

func (b *BaseExercise) Hydrate() (string, error) {
	data := []model.BaseExercise{
		// Compound Movements
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bench Press", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Overhead Press", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Bent Over Rows", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Convental Deadlifts", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Romanian Deadlift", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Sumo Deadlifts", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Squats", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Barbell Hip Thrust", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		{ID: uuid.New().String(), Name: "Wide Squats", MuscleGroup: "", SpecificParts: "", Level: 3, AvoidGiven: nil, MovementType: "Compound"},
		// {ID: uuid.New().String(), Name: "NEW ADDED", MuscleGroup: "", Level: 3, AvoidGiven: "Sciatic Nerve damage", MovementType: string(models.Compound)},
		// "Isolated" Movements
		//chest
		{ID: uuid.New().String(), Name: "Chest Flys", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Incline Bench", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dips", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dumbell Pullovers", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		// back
		{ID: uuid.New().String(), Name: "Single Arm Rows", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Cable Rows", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Pull Ups", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Diverging Rows", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Lateral Pull Downs", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		// arms
		{ID: uuid.New().String(), Name: "Bicep Curls", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Tricep Extensions", MuscleGroup: "", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Curl Machine", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hammer Curls", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Skull Crushes", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		// shoulders
		{ID: uuid.New().String(), Name: "Reverse Pec Deck", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Face Pulls", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Dumbell Lat Raises", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Cable Lat Raises", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Front Raises", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Upright Rows", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		// legs
		{ID: uuid.New().String(), Name: "Leg Extensions", MuscleGroup: "", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hamstring Curls", MuscleGroup: "", Level: 3, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Leg Press", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Hack Squat", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Bulgarian Split Squat", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
		{ID: uuid.New().String(), Name: "Glute-Ham Raises", MuscleGroup: "", Level: 1, AvoidGiven: nil, MovementType: "Isolated"},
	}

	// TODO: refactor badly..
	for _, eachExercise := range data {
		rows, err := b.db.Model(&model.BaseExercise{}).Select("name", "avoid_given").Rows()
		if err != nil {
			return "unable to hydrate base exercise table", errors.Wrapf(err, "db.Model.Select.Rows")
		}

		defer rows.Close()
		var name string
		var avoidGiven []*model.AvoidGiven
		var count int
		for rows.Next() {
			err := rows.Scan(&name, &avoidGiven)
			if err != nil {
				return "unable to hydrate base exercise table", errors.Wrapf(err, "rows.Scan")
			}
			if name == eachExercise.Name {
				count += 1
				fmt.Printf("%v ,exists in database!\n", eachExercise.Name)
			}
		}
		if count == 0 {
			b.db.Create(eachExercise)
		}
	}
	return "successfully hydrated base exercise table", nil
}
