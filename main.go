package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	namespace string = "default"
	sleepTime int    = 5
)

func init() {
	// load environment values
	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime, _ = strconv.Atoi(os.Getenv("SLEEP_TIME"))
	}
	if len(os.Getenv("NAMESPACE")) != 0 {
		namespace = os.Getenv("NAMESPACE")
	}
}

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get pods in the namespaces
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		// print message
		fmt.Printf("Existen %d pods en el namespace %s\n", len(pods.Items), namespace)

		// sleep time
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
