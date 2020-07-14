package impl

import (
	"fmt"
	"github.com/kataras/iris/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"log"
	"simple-k8/go-common/api-base"
	"simple-k8/go-common/utils"
	"simple-k8/model"
	"strconv"
)

// pod列表
func ListPods(ctx context.Context) apibase.Result {
	nameSpace := ctx.FormValue("nameSpace")
	page := ctx.FormValue("page")
	size := ctx.FormValue("size")

	pager, err := strconv.ParseInt(page, 10, 32)
	index, err := strconv.ParseInt(size, 10, 32)

	clientSet := utils.K8sClient()
	//获取PODS，不传namespace默认查全部
	pods, err := clientSet.CoreV1().Pods(nameSpace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	i := 0
	products := []map[string]interface{}{}

	for _, d := range pods.Items {
		products = append(products, map[string]interface{}{
			"name":      d.Name,
			"nameSpace": d.Namespace,
			"id":        i,
			"time":      d.CreationTimestamp,
		})
		i++
	}
	arr := products[(int32(pager)-1)*int32(index) : (int32(pager)-1)*int32(index)+int32(index)]

	return map[string]interface{}{
		"products": arr,
		"total":    len(products),
	}

}

// deployment列表
func ListDeployments(ctx context.Context) apibase.Result {
	nameSpace := ctx.FormValue("nameSpace")
	page := ctx.FormValue("page")
	size := ctx.FormValue("size")

	pager, err := strconv.ParseInt(page, 10, 32)
	index, err := strconv.ParseInt(size, 10, 32)

	clientSet := utils.K8sClient()
	deploymentsClient := clientSet.AppsV1().Deployments(nameSpace)
	deployments, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	j := 0
	products := []map[string]interface{}{}
	for _, d := range deployments.Items {
		// 名称 -> 容量
		//fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		products = append(products, map[string]interface{}{
			"deploymentName": d.Name,
			"podName":        d.Name + "-*",
			"nameSpace":      d.Namespace,
			"replicas":       d.Spec.Replicas,
			"id":             j,
			"time":           d.CreationTimestamp,
		})
		j++
	}
	arr := products[(int32(pager)-1)*int32(index) : (int32(pager)-1)*int32(index)+int32(index)]

	return map[string]interface{}{
		"products": arr,
		"total":    len(products),
	}
}

func DeploymentDetail(ctx context.Context) apibase.Result {
	nameSpace := ctx.FormValue("nameSpace")
	deploymentName := ctx.FormValue("deploymentName")
	clientSet := utils.K8sClient()
	deploymentsClient := clientSet.AppsV1().Deployments(nameSpace)
	deployment, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	products := []map[string]interface{}{}
	products = append(products, map[string]interface{}{
		"deployName": deployment.Name,
		"podName":    deployment.Name + "-*",
		"nameSpace":  deployment.Namespace,
		"replicas":   deployment.Spec.Replicas,
		"time":       deployment.CreationTimestamp,
	})

	return map[string]interface{}{
		"products": products,
	}
}

// 扩缩容
func UpdateDeploymentReplicas(ctx context.Context) apibase.Result {
	deploymentName := ctx.FormValue("deploymentName")
	nameSpace := ctx.FormValue("nameSpace")
	replicas := ctx.FormValue("replicas")
	num, err := strconv.ParseInt(replicas, 10, 32)
	if err != nil {
		panic(err)
	}
	var re int32
	clientSet := utils.K8sClient()
	deploymentsClient := clientSet.AppsV1().Deployments(nameSpace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("失败: %v", getErr))
		}
		re = *result.Spec.Replicas
		result.Spec.Replicas = int32Ptr(int32(num))
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("扩容失败: %v", retryErr))
	}
	fmt.Println("更新成功")
	deName := model.Kube.InsertKubeRecord(re, int32(num), nameSpace, deploymentName)
	if deName == deploymentName {
		fmt.Println("扩容记录写入数据库成功")
	} else {
		panic(fmt.Errorf("写入数据库记录失败"))
	}

	return map[string]interface{}{
		"status": "true",
	}
}

func int32Ptr(i int32) *int32 { return &i }

func Test(ctx context.Context) apibase.Result {
	id_ := ctx.FormValue("id")
	type_ := ctx.FormValue("type")
	fmt.Println(type_ + "=======================")
	fmt.Println(id_ + "=======================")
	if type_ == "cluster" {
		return map[string]interface{}{
			"url": "hello" + id_,
		}
	}
	if type_ == "services" {
		return map[string]interface{}{
			"url": "world" + id_,
		}
	}
	return nil
}
