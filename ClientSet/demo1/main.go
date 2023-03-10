package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	CoreV1 := clientset.CoreV1()
	pod, err := CoreV1.Pods("default").Get(context.TODO(), "test", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Status)
}
