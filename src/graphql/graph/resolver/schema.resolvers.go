package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/model"
)

func (r *mutationResolver) HydrateMuscleGroups(ctx context.Context) (string, error) {
	for _, eachMuscleGroup := range graph.MuscleGroupData {
		rows, err := r.DB.Model(&model.MuscleGroup{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachMuscleGroup.Name)
		}

		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachMuscleGroup.Name {
				count += 1
				fmt.Printf("%v , exists in database!\n", eachMuscleGroup.Name)
			}
		}

		if count == 0 {
			muscleGroup := model.MuscleGroup{
				Name: eachMuscleGroup.Name,
			}
			r.DB.Create(muscleGroup)
		}
	}

	return "Muscle group table has been hydrated", nil
}

func (r *mutationResolver) HydrateSpecificParts(ctx context.Context) (string, error) {
	for _, eachSpecificMuscleGroup := range graph.SpecificMuscleGroupData {
		rows, err := r.DB.Model(&model.SpecificParts{}).Select("name").Rows()
		if err != nil {
			fmt.Printf("%v, selecting database\n", eachSpecificMuscleGroup.Name)
		}

		defer rows.Close()

		var name string
		var count int

		for rows.Next() {
			rows.Scan(&name)
			if name == eachSpecificMuscleGroup.Name {
				fmt.Printf("%v, exists in database!\n", eachSpecificMuscleGroup.Name)
			}
		}

		if count == 0 {
			specificMuscleGroup := model.SpecificParts{
				Name:        eachSpecificMuscleGroup.Name,
				MuscleGroup: eachSpecificMuscleGroup.MuscleGroup,
			}

			r.DB.Create(specificMuscleGroup)
		}
	}

	return "Specific Parts has been hydrated", nil
}

func (r *queryResolver) GetMuscleSpecifics(ctx context.Context, input *model.MuscleSpecificInput) ([]string, error) {
	allSpecifics := []string{}
	r.DB.Model(&model.SpecificParts{}).Where("name = ?", input.Name).Scan(&allSpecifics)

	return allSpecifics, nil
}
