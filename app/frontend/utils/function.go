package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	if ctx.Value(UserIdKey) == nil {
		return 0
	}
	return uint32(ctx.Value(UserIdKey).(float64))
}
