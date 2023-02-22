package utils

import (
	"flag"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetKubeConfig() *rest.Config {
	// 声明kubeconfig配置文件
	var kubeconfig *string
	// 获取home环境变量，拿到$HOME/.kube/config配置文件
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		//否则就根据kubeconfig传到获得config的路径
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig\nfile")
	}
	flag.Parse()

	// 将kubeconfig格式化为rest.config类型

	//config, err := clientcmd.BuildConfigFromFlags("", "kubeconfi")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}
