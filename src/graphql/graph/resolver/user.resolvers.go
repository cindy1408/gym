package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"net/mail"

	"github.com/cindy1408/gym/src/graphql/graph/model"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (string, error) {
	if input.FirstName == "" {
		return "first name is missing", nil
	}

	if input.LastName == "" {
		return "last name is missing", nil
	}

	if input.Password == "" {
		return "password is missing", nil
	}

	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return "invalid email address", nil
	}

	hashedPw := Hasher(input.Password)

	newUser := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPw,
	}

	postgres := postgres.Postgres{}
	result, err := postgres.AddUserToDB(&newUser)
	if err != nil {
		return "", errors.Wrapf(err, "r.AddUserToDB")
	}

	return result, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	postgres := postgres.Postgres{}
	allUsers, err := postgres.GetAllUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "postgres.GetAllUsers")
	}

	return allUsers, nil
}
