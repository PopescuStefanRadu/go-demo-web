package repository

import (
	"context"
	"database/sql"
	"faceit/pkg/errors"
	"faceit/pkg/query"
	"faceit/pkg/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
)

const (
	colId        = "id"
	colFirstName = "first_name"
	colLastName  = "last_name"
	colNickname  = "nickname"
	colEmail     = "email"
	colCountry   = "country"
	colCreatedAt = "created_at"
	colUpdatedAt = "updated_at"
)

var allCols = []string{colId, colFirstName, colLastName, colNickname, colEmail, colCountry, colCreatedAt, colUpdatedAt}

type Config struct {
	TargetVersion     uint
	MigrationLocation string
	DatabaseUrl       string
}

type UserDBRepo struct {
	Config
	sql.
}

func (r UserDBRepo) Migrate(_ context.Context) error {
	m, err := migrate.New("file://../../database/migration.sql", "pgx://sql:sql@localhost:5432/faceit?search_path=public")
	if err != nil {
		return errors.LibErr{
			Fault: errors.FaultLib,
			Err:   err,
			Msg:   "could not create migration",
		}
	}
	if r.TargetVersion == 0 {
		if err := m.Up(); err != nil {
			return errors.LibErr{
				Fault: errors.FaultUnknown,
				Err:   err,
				Msg:   "could not migrate db",
			}
		}
		return nil
	}
	if err := m.Migrate(r.TargetVersion); err != nil {
		return errors.LibErr{Fault: errors.FaultUnknown, Err: err, Msg: "could not migrate db"}
	}
	return nil
}

func (UserDBRepo) FindUsers(ctx context.Context, filter *user.FilterQuery) (*query.PaginatedResponse[user.User], error) {
	selectQuery := createQuery(filter)
	selectQuery.RunWith()
	panic("implement me")
}

func createQuery(filter *user.FilterQuery) sq.SelectBuilder {
	selectBuilder := sq.Select(allCols...)
	if filter.IdEquals.IsPresent {
		selectBuilder = selectBuilder.Where(sq.Eq{colId: filter.IdEquals.Val})
	}
}

func (UserDBRepo) DeleteUser(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (UserDBRepo) AddUser(ctx context.Context, addUserInput user.User) error {
	//TODO implement me
	panic("implement me")
}

func (UserDBRepo) UpdateUser(ctx context.Context, input *user.UpdateUserInput) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}
