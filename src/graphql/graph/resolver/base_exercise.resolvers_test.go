package resolver_test

import (
	"context"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("BaseExercise.Resolvers", func() {
	var (
		ctx context.Context
		err error

		db           *gorm.DB
		postgresRepo postgres.Repo
		mockBaseExercise := mocks_postgres.mockBaseExercise

		baseExerciseName string
		baseExercise     *model.BaseExercise
		baseExerciseErr  error
	)

	BeforeEach(func() {
		ctx = context.Background()
		// start a new database
	})

	AfterEach(func() {
		err = postgres.DeleteBaseExerciseByName(db, baseExerciseName)
		Expect(err).NotTo(HaveOccurred())
		// close the database
	})

	When("getting a base exercise that does not exist", func() {
		BeforeEach(func() {
			baseExerciseName = "does not exist"

			baseExercise, baseExerciseErr = postgresRepo.GetBaseExerciseByName(ctx, db, baseExerciseName)
			Expect(baseExerciseErr).NotTo(HaveOccurred())
		})

		It("should return an error", func() {
			Expect(baseExerciseErr).To(HaveOccurred())
		})

		It("should return nil Base Exercise", func() {
			Expect(baseExercise).To(BeNil())
		})
	})
})
