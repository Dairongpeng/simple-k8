package impl

import (
	"github.com/kataras/iris/context"
	"simple-k8/go-common/api-base"
	"simple-k8/log"
)

const (
	ENABLE  = 0
	DISABLE = 1
)

var (
	VerifyCode     bool
	VerifyIdentity bool
)

func Login(ctx context.Context) apibase.Result {
	log.Debugf("[Login] Login from API ")

	return true
}

func LogOut(ctx context.Context) apibase.Result {
	log.Debugf("[LogOut] LogOut from API ")

	userId, err := ctx.Values().GetInt("userId")
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	apibase.DeleteUserCache(userId)
	ctx.RemoveCookie("em_token")
	return true
}

func UserInfo(ctx context.Context) apibase.Result {
	return true
}

func Register(ctx context.Context) (rlt apibase.Result) {

	return 0
}

func ResetPwdByAdmin(ctx context.Context) apibase.Result {
	log.Debugf("[AddUser] AddUser from API ")
	return nil
}

func ModifyInfoById(ctx context.Context) apibase.Result {
	log.Debugf("[ModifyUserById] ModifyUserById from API ")

	return nil
}

func ModifyPwdById(ctx context.Context) apibase.Result {
	log.Debugf("[ModifyUserById] ModifyUserById from API ")
	return nil
}
