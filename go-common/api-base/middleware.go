package apibase

import (
	"errors"
	"fmt"
	"net/http"
	_ "runtime/debug"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/context"
)

const (
	JWT_SECRET = "SGVsbG8gV29ybGQK"
)

type user struct {
	UserId   int
	UserName string
}

var (
	userCache = map[int]user{}
)

func SetUserCache(userId int, username string) {
	userCache[userId] = user{UserId: userId, UserName: username}
}

func DeleteUserCache(userId int) {
	delete(userCache, userId)
}

func getUserCache(userId int) *user {
	userInfo, ok := userCache[userId]
	if ok {
		return &userInfo
	}
	return nil
}

func ApiValidateCookies(ctx context.Context) bool {

	if strings.HasSuffix(ctx.Path(), "register") ||
		strings.HasSuffix(ctx.Path(), "login") ||
		strings.HasSuffix(ctx.Path(), "license") ||
		strings.HasSuffix(ctx.Path(), "upload") ||
		strings.HasSuffix(ctx.Path(), "dt_agent_health_check") ||
		strings.HasSuffix(ctx.Path(), "dt_agent_error") ||
		strings.HasSuffix(ctx.Path(), "dt_agent_host_resource") ||
		strings.HasSuffix(ctx.Path(), "dt_agent_performance") ||
		strings.HasSuffix(ctx.Path(), "callback") ||
		strings.HasSuffix(ctx.Path(), "identity") ||
		strings.Contains(ctx.Path(), "Captcha") {
		return true
	}

	reject := func(result interface{}, format string, args ...interface{}) {
		do_panic := false
		if b, ok := result.(bool); ok {
			if b {
				do_panic = true
			}
		} else if err, ok := result.(error); ok {
			if err != nil {
				do_panic = true
			}
		}
		if do_panic {
			panic(fmt.Sprintf(format, args...))
		}
	}

	defer func() {
		if r := recover(); r != nil {
			Feedback(ctx, &AccessDeniedError{errors.New(r.(string))})
		}
	}()

	token := ctx.GetCookie("em_token")

	// validate
	tk, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if tk.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("Unexpected auth method: %s", tk.Method.Alg())
		}
		return []byte(JWT_SECRET), nil
	})
	reject(err, "Validate em_token failure: %s", err)
	reject(!tk.Valid, "Invalid em_token: %s", token)

	var expiration time.Time
	var info *user
	if claims, ok := tk.Claims.(jwt.MapClaims); ok {
		if id, ok := claims["user_id"].(float64); ok {
			info = getUserCache(int(id))
			reject(info == nil, "Invalid user_id (%s) in em_token", id)
			reject(int(id) != info.UserId, "Unmatched user_id in cookies")
		} else {
			reject(true, "Missing 'user_id' in em_token")
		}
		if claim_username, ok := claims["user_name"].(string); ok {
			reject(claim_username != info.UserName, "Unmatched user_name in cookies")
		} else {
			reject(true, "Missing 'user_name' in em_token")
		}
		exp, ok := claims["exp"].(float64)
		reject(!ok, "Missing 'exp' in em_token")
		expiration = time.Unix(int64(exp), 0)
		reject(time.Now().After(expiration), "em_token is expired")
	} else {
		reject(true, "Unable to map em_token")
	}

	// shift to next
	ctx.Values().Set("userId", info.UserId)
	ctx.Values().Set("username", info.UserName)
	ctx.Values().Set("expiration", expiration)

	return true
}

func CreateToken(username string, userId int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user_name"] = username
	claims["user_id"] = userId
	claims["exp"] = time.Now().Unix() + 259200
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))

	SetUserCache(userId, username)

	return tokenString
}

var (
	SetCookieExpiration = time.Duration(259200) * time.Second
)

func SetCookie(ctx context.Context, name, value string) {
	c := &http.Cookie{}
	c.Name = name
	c.Value = value
	c.HttpOnly = true
	c.Expires = time.Now().Add(SetCookieExpiration)
	c.MaxAge = int(SetCookieExpiration.Seconds())
	ctx.SetCookie(c)

	ctx.Header(name, value)
}
