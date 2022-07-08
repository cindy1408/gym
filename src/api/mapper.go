package api

import (
	e "github.com/cindy1408/gym/src/api/graphql/graph/model"
	i "github.com/cindy1408/gym/src/api/model"
)

// converts graphql model to src/model

func ExternalBaseExerciseToInternalMapper(baseExercise e.BaseExercise) i.BaseExercise {
	return i.BaseExercise{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	}
}

func ExternalBaseExerciseInputToInternalMapper(baseExercise e.BaseExerciseInput) i.BaseExerciseInput {
	return i.BaseExerciseInput{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	}
}

func ExternalEachExerciseToInternalMapper(eachExercise e.EachExercise) i.EachExercise {
	return i.EachExercise{
		ID:                eachExercise.ID,
		UserWorkoutPlanID: eachExercise.UserWorkoutPlanID,
		Name:              eachExercise.Name,
		Weight:            eachExercise.Weight,
		Unit:              eachExercise.Unit,
		Sets:              eachExercise.Sets,
		Reps:              eachExercise.Reps,
	}
}

func ExternalEachExerciseInputToInternalMapper(eachExercise e.EachExerciseInput) i.EachExerciseInput {
	return i.EachExerciseInput{
		Name:   eachExercise.Name,
		Weight: eachExercise.Weight,
		Unit:   eachExercise.Unit,
		Sets:   eachExercise.Sets,
		Reps:   eachExercise.Reps,
	}
}

func ExternalMuscleSpecificInputToInternalMapper(muscle e.MuscleSpecificInput) i.MuscleSpecificInput {
	return i.MuscleSpecificInput{
		Name: muscle.Name, 
		MuscleGroup: muscle.MuscleGroup,
	}
}


func InternalBaseExerciseToExternalMapper(baseExercise i.BaseExercise) e.BaseExercise {
	return e.BaseExercise{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	}
}

func InternalEachExerciseToExternalMapper(eachExercise i.EachExercise) e.EachExercise {
	return e.EachExercise{
		ID:                eachExercise.ID,
		UserWorkoutPlanID: eachExercise.UserWorkoutPlanID,
		Name:              eachExercise.Name,
		Weight:            eachExercise.Weight,
		Unit:              eachExercise.Unit,
		Sets:              eachExercise.Sets,
		Reps:              eachExercise.Reps,
	}
}

func InternalBaseExerciseInputToExernalMapper(baseExercise i.BaseExerciseInput) e.BaseExerciseInput {
	return e.BaseExerciseInput{
		Name:          baseExercise.Name,
		MuscleGroup:   baseExercise.MuscleGroup,
		SpecificParts: baseExercise.SpecificParts,
		Level:         baseExercise.Level,
		AvoidGiven:    baseExercise.AvoidGiven,
		MovementType:  baseExercise.MovementType,
	}
}