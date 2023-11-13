package main
import (
        "strings"
        "context"
        "fmt"
        "k8s.io/client-go/kubernetes"
        k8srest "k8s.io/client-go/rest"
	"github.com/prometheus/common/expfmt"
)

func main() {
	config, err := k8srest.InClusterConfig()
	if err != nil {
		fmt.Printf("error getting in-cluster config: %v\n", err)
		return
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

	metric := metricFamilies["apiserver_request_total"]
	for _, m := range metric.Metric {
		fmt.Printf("Labels:\n")
		for _, l := range m.Label {
			fmt.Printf("%s: %s\n", l.GetName(), l.GetValue())
		}
		fmt.Printf("Value: %v\n", m.GetCounter().GetValue())
	}
}


