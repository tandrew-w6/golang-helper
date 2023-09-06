package context

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

func CheckAuthentication(ctx context.Context, keys ...string) (metadata.MD, error) {
	m, isExist := metadata.FromIncomingContext(ctx)
	if !isExist {
		return nil, errors.New(": unauthenticated")
	}

	for _, key := range keys {
		if len(m.Get(key)) == 0 {
			return nil, errors.New(": unauthenticated")
		}
	}

	return m, nil
}

func GetFromMetadata(m metadata.MD, key string) string {
	var (
		result string
	)
	if len(m.Get(key)) > 0 && m.Get(key)[0] != "" {
		result = m.Get(key)[0]
	}
	return result
}

func SetOutgoingContext(ctx context.Context, keys ...string) context.Context {
	var (
		m, _ = CheckAuthentication(ctx)
	)
	for _, key := range keys {
		if len(m.Get(key)) > 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, key, m.Get(key)[0])
		}
	}
	return ctx
}
