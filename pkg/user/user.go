package user

import (
	"faceit/pkg/query"
	"faceit/pkg/types"
	"time"
)

type User struct {
	Id        string
	FirstName *string
	LastName  *string
	Nickname  *string
	Email     string
	Country   *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UpdateUserInput struct {
	FirstName types.Option[string]
	LastName  types.Option[string]
	Nickname  types.Option[string]
	Password  types.Option[string]
	Email     types.Option[string]
	Country   types.Option[string]
}

type FilterQuery struct {
	query.Paged
	query.Sorted
	IdEquals types.Option[string]

	FirstNameEquals types.Option[string]
	FirstNameLike   *string

	LastNameEquals types.Option[string]
	LastNameLike   *string

	NicknameEquals types.Option[string]
	NicknameLike   *string

	EmailEquals types.Option[string]
	EmailLike   *string

	CountryEquals types.Option[string]
	CountryLike   *string
	CountryIn     []string

	CreatedAtEquals      types.Option[time.Time]
	CreatedAtLessThan    *time.Time
	CreatedAtGreaterThan *time.Time

	UpdatedAtEquals      types.Option[time.Time]
	UpdatedAtLessThan    *time.Time
	UpdatedAtGreaterThan *time.Time
}
