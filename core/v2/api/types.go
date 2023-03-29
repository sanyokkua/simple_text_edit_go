package api

import (
	"context"
)

type ContextProvider func() (ctx context.Context)
type Destination string
type Button string
