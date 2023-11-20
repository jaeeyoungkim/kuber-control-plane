package main

import (
	"context"
	"fmt"
	"strings"

	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"k8s.io/client-go/kubernetes"
	k8srest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type apiServerRequestTotal struct {
	code        string
	verb        string
	Driver      string
	Mount       string
	Total       int64
	Used        int64
	UsedPercent float64
}

func getApiserverRequestTotal(metric *io_prometheus_client.MetricFamily) {
	for _, m := range metric.Metric {
		fmt.Printf("Labels:\n")
		for _, l := range m.Label {
			fmt.Printf("%s: %s\n", l.GetName(), l.GetValue())
		}
		fmt.Printf("Value: %v\n", m.GetCounter().GetValue())
		fmt.Printf("\n")
	}
}

func main() {
	config, err := k8srest.InClusterConfig()
	if err != nil {
		fmt.Printf("error getting in-cluster config: %v\n", err)
		kubeconfig := "/Users/jaeyoung/.kube/config"
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("error getting out-of-cluster config: %v\n, err")
			return
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error creating clientset: %v\n", err)
		return
	}

	bytes, err := clientset.RESTClient().Get().AbsPath("/metrics").DoRaw(context.Background())
	if err != nil {
		fmt.Printf("error getting metrics: %v\n", err)
		return
	}

	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Printf("error parsing metrics: %v\n", err)
		return
	}

	for key := range metricFamilies {
		fmt.Println(key)
	}

	var apiserverRequestTotalMetrics *io_prometheus_client.MetricFamily
	apiserverRequestTotalMetrics = metricFamilies["apiserver_request_total"]
	_ = apiserverRequestTotalMetrics
	// getApiserverRequestTotal(apiserverRequestTotalMetrics)
	// for _, m := range metric.Metric {
	// 	fmt.Printf("Labels:\n")
	// 	for _, l := range m.Label {
	// 		fmt.Printf("%s: %s\n", l.GetName(), l.GetValue())

	// 	}
	// 	fmt.Printf("Value: %v\n", m.GetCounter().GetValue())
	// 	fmt.Printf("\n")
	// }
}
