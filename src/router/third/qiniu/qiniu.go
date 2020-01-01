package qiniu

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/pkg/setting"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func GetToken(ctx *context.Context) {
	accessKey := setting.QiniuSetting.AccessKey
	secretKey := setting.QiniuSetting.SecretKey
	bucket := setting.QiniuSetting.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	values := map[string]interface{}{"token": upToken, "succMsg": "获取token成功"}

	ctx.GenResSuccess(values)

}
