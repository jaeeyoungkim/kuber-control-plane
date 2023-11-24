package main

import (
	"context"
	"encoding/json"
	"fmt"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"k8s.io/client-go/kubernetes"
	k8srest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
)

const (
	ConfigPath = "/USers/irostub/.kube/config"
)

type CommonResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type ApiServerRequestTotalBucket struct {
	//variable	type	JsonSerializeRule
	Code        string  `json:"code"`
	Component   string  `json:"component"`
	DryRun      string  `json:"dryRun"`
	Group       string  `json:"group"`
	Resource    string  `json:"resource"`
	Scope       string  `json:"scope"`
	Subresource string  `json:"subresource"`
	Verb        string  `json:"verb"`
	Version     string  `json:"version"`
	Counter     float64 `json:"counter"`
}

func getApiserverRequestTotalMetrics() []ApiServerRequestTotalBucket {
	client, err, done := getKubernetesClient()
	if !done {
		return nil
	}

	bytes, err := client.RESTClient().Get().AbsPath("/metrics").DoRaw(context.Background())
	if err != nil {
		fmt.Printf("error getting metrics: %v\n", err)
		return nil
	}

	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Printf("error parsing metrics: %v\n", err)
		return nil
	}

	var apiserverRequestTotalMetricFamily *io_prometheus_client.MetricFamily
	apiserverRequestTotalMetricFamily = metricFamilies["apiserver_request_total"]
	metric := apiserverRequestTotalMetricFamily.GetMetric()
	var datas []ApiServerRequestTotalBucket
	for _, m := range metric {
		label := m.GetLabel()
		data := ApiServerRequestTotalBucket{}

		for _, l := range label {
			name := l.GetName()
			value := l.GetValue()
			switch name {
			case "code":
				data.Code = value
			case "component":
				data.Component = value
			case "dryRun":
				data.DryRun = value
			case "group":
				data.Group = value
			case "resource":
				data.Resource = value
			case "scope":
				data.Scope = value
			case "subresource":
				data.Subresource = value
			case "verb":
				data.Verb = value
			case "version":
				data.Version = value
			}
		}
		data.Counter = m.GetCounter().GetValue()
		datas = append(datas, data)
	}
	return datas
}

func getKubernetesClient() (*kubernetes.Clientset, error, bool) {
	//load config from default config
	config, err := k8srest.InClusterConfig()
	if err != nil {
		fmt.Printf("error getting out-of-cluster config: %v\n", err)
		kubeConfig := ConfigPath
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

func main() {
	metrics := getApiserverRequestTotalMetrics()
	c := CommonResponse[[]ApiServerRequestTotalBucket]{
		Data:    metrics,
		Message: "this is ok",
	}
	marshal, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))
}
