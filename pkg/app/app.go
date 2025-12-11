package app

import "context"

type App interface {
	Run(ctx context.Context) error
	Close(ctx context.Context) error
}
