package utils

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/common/cache"
)

// 从api的Context中获取指定key的信息

// GetUidFromCtx 获取用户id
func GetUidFromCtx(ctx context.Context, key any) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(key).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

// GetUsernameFromCtx 获取用户username
func GetUsernameFromCtx(ctx context.Context) string {
	if val, ok := ctx.Value(cache.JwtFieldUserName).(string); ok {
		return val
	}
	return ""
}
