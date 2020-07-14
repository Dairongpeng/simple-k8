package api

import (
	"simple-k8/api/impl"
	"simple-k8/go-common/api-base"
)

var KubeApis = apibase.Route{
	Path: "kube",
	SubRoutes: []apibase.Route{{
		Path: "pod/list",
		GET:  impl.ListPods,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "pod列表",
				Query: apibase.ApiParams{
					"nameSpace": apibase.ApiParam{"string", "'xxx' or ''", "", false},
					"page":      apibase.ApiParam{"string", "页面容量", "", true},
					"size":      apibase.ApiParam{"string", "页面下标", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$[*].podName":   apibase.ApiReturn{"string", "pod名"},
						"$[*].nameSpace": apibase.ApiReturn{"string", "名称空间"},
						"$[*].id":        apibase.ApiReturn{"int", "唯一下标"},
						"$[*].time":      apibase.ApiReturn{"string", "创建时间"},
						"$.total":        apibase.ApiReturn{"int", "全部记录数"},
					},
				}},
			},
		},
	}, {
		Path: "deployment/list",
		GET:  impl.ListDeployments,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "控制器列表",
				Query: apibase.ApiParams{
					"nameSpace": apibase.ApiParam{"string", "'xxx' or ''", "", false},
					"page":      apibase.ApiParam{"string", "页面容量", "", true},
					"size":      apibase.ApiParam{"string", "页面下标", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$[*].deploymentName": apibase.ApiReturn{"string", "控制器名"},
						"$[*].podName":        apibase.ApiReturn{"string", "pod名"},
						"$[*].nameSpace":      apibase.ApiReturn{"string", "名称空间"},
						"$[*].replicas":       apibase.ApiReturn{"int", "控制器pod容量"},
						"$[*].id":             apibase.ApiReturn{"int", "唯一下标"},
						"$[*].time":           apibase.ApiReturn{"string", "创建时间"},
						"$.total":             apibase.ApiReturn{"int", "全部记录数"},
					},
				}},
			},
		},
	}, {
		Path: "deployment/detail",
		GET:  impl.DeploymentDetail,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "控制器详情",
				Query: apibase.ApiParams{
					"nameSpace":      apibase.ApiParam{"string", "'xxx' or ''", "", true},
					"deploymentName": apibase.ApiParam{"string", "'xxx' or ''", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$[*].deploymentName": apibase.ApiReturn{"string", "控制器名"},
						"$[*].podName":        apibase.ApiReturn{"string", "pod名"},
						"$[*].nameSpace":      apibase.ApiReturn{"string", "名称空间"},
						"$[*].replicas":       apibase.ApiReturn{"int", "控制器pod容量"},
						"$[*].time":           apibase.ApiReturn{"string", "创建时间"},
					},
				}},
			},
		},
	}, {
		Path: "deployment/update/scale",
		GET:  impl.UpdateDeploymentReplicas,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "扩缩容",
				Query: apibase.ApiParams{
					"deploymentName": apibase.ApiParam{"string", "deployment控制器名", "", true},
					"nameSpace":      apibase.ApiParam{"string", "名称空间", "", true},
					"replicas":       apibase.ApiParam{"int", "扩缩容大小", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$[*].status": apibase.ApiReturn{"string", "状态"},
					},
				}},
			},
		},
	}, {
		Path: "test",
		GET:  impl.Test,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "控制器测试",
				Query: apibase.ApiParams{
					"type": apibase.ApiParam{"string", "'cluster' or 'services'", "", true},
					"id":   apibase.ApiParam{"string", "服务器组ID", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$[*].url": apibase.ApiReturn{"string", "返回值测试===="},
					},
				}},
			},
		},
	}},
}
