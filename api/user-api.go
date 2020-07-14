package api

import (
	"simple-k8/api/impl"
	"simple-k8/go-common/api-base"
)

var UserOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "user",
	SubRoutes: []apibase.Route{{
		Path: "modifyInfo",
		POST: impl.ModifyInfoById,
	}, {
		Path: "modifyPwd",
		POST: impl.ModifyPwdById,
	}, {
		Path: "register",
		POST: impl.Register,
	}, {
		Path: "list",
		POST: impl.UserInfo,
	}, {
		Path: "login",
		POST: impl.Login,
	}, {
		Path: "logout",
		POST: impl.LogOut,
	}},
}
