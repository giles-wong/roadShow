package tools

import (
	"errors"
	"github.com/giles-wong/roadShow/global"
)

func GetAppSecret(appKey string) (string, error) {

	switch appKey {
	case "admin-pc":
		return global.App.Config.Signature.AdminApi, nil
	case "user-pc":
		return global.App.Config.Signature.UserApi, nil
	default:
		return "", errors.New("appKey not found")
	}
}
