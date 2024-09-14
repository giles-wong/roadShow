package tools

import (
	"errors"
	"github.com/giles-wong/roadShow/global"
)

func GetAppSecret(appKey string) (string, error) {

	switch appKey {
	case "admin-pc":
		return global.App.Config.SignConf.Admin, nil
	case "user-pc":
		return global.App.Config.SignConf.User, nil
	default:
		return "", errors.New("appKey not found")
	}
}
