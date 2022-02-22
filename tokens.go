package film

import "context"

//go:generate mockery --name TokenSaver
type TokenSaver interface {
	Save(context.Context, string) error
}

//go:generate mockery --name TokenProvider
type TokenProvider interface {
	IsTokenBlacklisted(context.Context, string) (bool, error)
}
