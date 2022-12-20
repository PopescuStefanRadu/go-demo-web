package user

import (
	"context"
	"faceit/pkg/errors"
	"faceit/pkg/query"
	"faceit/pkg/types"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var upperCaseValidator *regexp.Regexp
var lowerCaseValidator *regexp.Regexp
var digitValidator *regexp.Regexp
var specialCharacterValidator *regexp.Regexp

func init() {
	var err error
	upperCaseValidator, err = regexp.Compile("[[:upper:]]]")
	if err != nil {
		panic(err)
	}
	lowerCaseValidator, err = regexp.Compile("[[:lower:]]")
	if err != nil {
		panic(err)
	}
	digitValidator, err = regexp.Compile("\\d")
	if err != nil {
		panic(err)
	}
	specialCharacterValidator, err = regexp.Compile("[[:punct:]]")
	if err != nil {
		panic(err)
	}
}

type Service struct {
	Repo Repository
}

func (s *Service) AddUser(ctx context.Context, user User) (User, error) {
	return s.Repo.AddUser(ctx, user)
}

func (s *Service) UpdateUser(ctx context.Context, update *UpdateUserInput) (*User, error) {
	return s.Repo.UpdateUser(ctx, update)
}

func (s *Service) UpdatePassword(ctx context.Context, newPassword string) (*User, error) {
	if !(upperCaseValidator.MatchString(newPassword) &&
		lowerCaseValidator.MatchString(newPassword) &&
		digitValidator.MatchString(newPassword) &&
		specialCharacterValidator.MatchString(newPassword)) {
		return nil, errors.LibErr{Msg: "password is not strong enough", Fault: errors.FaultClient}
	}
	password, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		return nil, fmt.Errorf("could not generate password: %w", err)
	}
	return s.Repo.UpdateUser(ctx, &UpdateUserInput{Password: types.NewOption(string(password))})
}

func (s *Service) DeleteUser(ctx context.Context, id string) (*User, error) {
	return s.Repo.DeleteUser(ctx, id)
}

func (s *Service) FindUsers(ctx context.Context, filter *FilterQuery) (*query.PaginatedResponse[User], error) {
	return s.Repo.FindUsers(ctx, filter)
}
