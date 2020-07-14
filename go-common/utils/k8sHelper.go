package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

// 获取集群内部k8s客户端
func K8sClient() *kubernetes.Clientset {
	// 使用当前上下文环境
	kubeconfig := filepath.Join(
		os.Getenv("KUBECONFIG"),
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	// 根据指定的 config 创建一个新的 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}
