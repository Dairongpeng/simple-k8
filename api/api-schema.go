package api

import "simple-k8/go-common/api-base"

var ApiV1Schema = apibase.Route{
	Path: "/api/v1",
	SubRoutes: []apibase.Route{
		UserOperationEasyMatrixAPIRoutes,
		KubeApis,
	},
}
