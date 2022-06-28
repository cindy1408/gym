package postgres_test

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Postgres", func() {
	var (
		ctx context.Context
		err          error

		db           *gorm.DB
		postgresRepo postgres.Repo

		baseExerciseName string
		baseExercise     *model.BaseExercise
		baseExerciseErr  error
	)

	BeforeEach(func() {
		ctx = context.Background()
	})

	AfterEach(func() {
		err = postgres.DeleteBaseExerciseByName(db, baseExerciseName)
		Expect(err).NotTo(HaveOccurred())

		// NEED TO CLOSE THE CONNECTION TO THE DATABASE
	})

	When("getting base exercise by name does not exist", func() {
		BeforeEach(func() {
			baseExerciseName = "does not exist"

			baseExercise, baseExerciseErr = postgresRepo.GetBaseExerciseByName(ctx, db, baseExerciseName)
			Expect(baseExerciseErr).NotTo(HaveOccurred())
		})

		It("should not return error nil", func() {
			Expect(baseExerciseErr).To(HaveOccurred())
		})

		It("should return nil Base Exercise", func() {
			Expect(baseExercise).To(BeNil())
		})
	})

})
