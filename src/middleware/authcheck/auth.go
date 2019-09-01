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
		if obj == nil {
			ctx.GenResError(e.ERROR_AUTH_CHECK_TOKEN_FAIL, "鉴权失败")
			ctx.Abort()
		}
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
func StaffRequire(ctx *context.Context) {
	var code int
	var token string

	code = e.SUCCESS

	headers := ctx.Request.Header["Hualvmall_staff_authorization"]

	if len(headers) > 0 {
		token = headers[0]
	}

	if token == "" {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		obj, err := auth.ParseToken(token)
		fmt.Printf("%+v\n", obj)
		if obj == nil {
			ctx.GenResError(e.ERROR_AUTH_CHECK_TOKEN_FAIL, "鉴权失败")
			ctx.Abort()
		}
		// 获取 uid 信息

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
