package main

import (
	"KubernetesDevelopment/utils"
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	config := utils.GetKubeConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// 指定命名空间
	namespace := "default"
	pod := ""

	list, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("获取default命名空间的pod")
	for _, item := range list.Items {
		fmt.Println(item.Name)
		pod = item.Name
	}

	podInfo, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("获取pod的PodIP")
	fmt.Println(podInfo.Status.PodIP)

	// 删除某个pod
	err = clientset.CoreV1().Pods(namespace).Delete(context.TODO(), pod, v1.DeleteOptions{})
	if err != nil {
		fmt.Println("删除失败", err)
	}

	// 更新pod(完整yaml)
	// k8s.io/api/core/v1
	//clientset.CoreV1().Pods(namespace).Update(context.TODO(),&v12.Pod{},v1.UpdateOptions{})

	// 获取deployment副本数
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(), "test-nginx", v1.GetOptions{})
	if err != nil {
		fmt.Println("查询", err)
	}
	fmt.Println("查询test-nginx deployment的副本数")
	fmt.Println(scale.Status.Replicas)

	// 创建deployment
	// k8s.io/api/apps/v1
	//clientset.AppsV1().Deployments(namespace).Create(context.TODO(), &v12.Deployment{}, v1.CreateOptions{})

}
