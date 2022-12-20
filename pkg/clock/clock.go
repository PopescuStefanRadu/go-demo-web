package clock

import "time"

//go:generate mockgen -source clock.go -destination mock/clock.go

type Clock interface {
	Now() time.Time
}
