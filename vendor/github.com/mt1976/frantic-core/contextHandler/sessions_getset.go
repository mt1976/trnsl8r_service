package contextHandler

import (
	"context"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
)

// sessionKey      = new(cfg.GetSecuritySessionKey_Session())
// userKeyKey      = new(cfg.GetSecuritySessionKey_UserKey())
// userCodeKey     = new(cfg.GetSecuritySessionKey_UserCode())
// tokenKey        = new(cfg.GetSecuritySessionKey_Token())
// expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())

func GetUserCode(ctx context.Context) string {
	value := ctx.Value(userCodeKey)
	if value == nil {
		logHandler.WarningLogger.Printf("User code requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetUserKey(ctx context.Context) string {
	value := ctx.Value(userKeyKey)
	if value == nil {
		logHandler.WarningLogger.Printf("User key requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSessionID(ctx context.Context) string {
	value := ctx.Value(sessionIDKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session ID requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSessionToken(ctx context.Context) any {
	value := ctx.Value(tokenKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session token requested but not found in context, returning nil")
		return nil
	}
	return value
}

func GetSessionExpiry(ctx context.Context) time.Time {
	value := ctx.Value(expiryPeriodKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session expiry requested but not found in context, returning zero time")
		return time.Time{}
	}
	return value.(time.Time)
}

func GetSessionIdentifier() string {
	return sessionIDKey.name
}

// Setters

func SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func SetUserKey(ctx context.Context, userKey string) context.Context {
	return context.WithValue(ctx, userKeyKey, userKey)
}

func SetUserCode(ctx context.Context, userCode string) context.Context {
	return context.WithValue(ctx, userCodeKey, userCode)
}

func SetSessionToken(ctx context.Context, token any) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func SetSessionExpiry(ctx context.Context, expiry time.Time) context.Context {
	return context.WithValue(ctx, expiryPeriodKey, expiry)
}
