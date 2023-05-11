package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

// Use cookie to store session id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middle
func AuthSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get(userkey)
		if sessionID == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "This page needs login",
			})
			return
		}
		ctx.Next()
	}
}

// Save Session for User
func SaveSession(ctx *gin.Context, userID string) {
	session := sessions.Default(ctx)
	session.Set(userkey, userID)
	session.Save()
}

// Clear Session for User
func ClearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

// Get Session for User
func GetSession(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	sessionID := session.Get(userkey)
	if sessionID == nil {
		return "0"
	}
	return sessionID.(string)
}

// Check Session for User
func CheckSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	sessionID := session.Get(userkey)
	return sessionID != nil
}
