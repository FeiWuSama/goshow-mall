package mapper

import "context"

type AdminMapper interface {
	HelloWorld(ctx context.Context) string
}
