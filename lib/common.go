package lib

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	k8srest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	configPath = "/Users/irostub/.kube/config"
)

func GetKubernetesClient() (*kubernetes.Clientset, error, bool) {
	//load config from default config
	config, err := k8srest.InClusterConfig()
	if err != nil {
		fmt.Printf("error getting out-of-cluster config: %v\n", err)
		kubeConfig := configPath
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			fmt.Printf("error getting out-of-cluster config: %v\n, err")
		}
	}

	//load client from default cluster config
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error creating clientset: %v\n", err)
		return nil, nil, false
	}
	return client, err, true
}
