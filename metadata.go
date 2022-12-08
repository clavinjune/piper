package piper

import "context"

type M[IN any] struct {
	Ctx  context.Context
	In   IN
	Data map[string]any
}
