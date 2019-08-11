package authcheck

import (
	"fmt"
	"staff-mall-center/pkg/auth"
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/e"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthRequired(ctx *context.Context) {
	var code int
	var token string

	code = e.SUCCESS

	headers := ctx.Request.Header["Hualvmall_authorization"]

	if len(headers) > 0 {
		token = headers[0]
	}

	if token == "" {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		obj, err := auth.ParseToken(token)
		fmt.Println(obj)
		ctx.Set("uname", obj.Username)
		ctx.Set("utype", obj.Logintype)
		ctx.Set("uid", obj.UID)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				fmt.Println(err)
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
	}

	if code != e.SUCCESS {
		ctx.GenResError(code, "")
		ctx.Abort()
	}
	ctx.Next()
}
