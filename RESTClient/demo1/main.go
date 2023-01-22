package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		panic(err)
	}
	// 下面的不加会出问题 但是还不知道啥意思
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	pod := v1.Pod{}
	err = restClient.Get().Namespace("kube-system").Resource("pods").Name("etcd-master").Do(context.TODO()).Into(&pod)
	if err != nil {
		print(err)
	} else {
		println(pod.Name)
	}

}
