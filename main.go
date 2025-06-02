package main

import (
	"os"
	"path/filepath"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func main() {
	namespace := flag.String("namespace", "default", "kubernetes namespace")
	kubeconfig := flag.String("kubeconfig",filepath.Join(os.Getenv("HOME"), ".kube", "config"), "path to kubeconfig file")
	flag.Parse()


	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

    pods, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod Name: %s\n", pod.Name)
	}
}