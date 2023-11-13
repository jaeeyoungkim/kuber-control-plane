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
        fmt.Printf("error")
        }
        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
                panic(err.Error())
        }

        bytes, err := clientset.RESTClient().Get().AbsPath("/metrics").DoRaw(context.Background())

        reader := strings.NewReader(string(bytes))

	var parser expfmt.TextParser
	parsed, err := parser.TextToMetricFamilies(reader)
	if err != nil {
		fmt.Println("Error parsing metrics:", err)
		return
	}

	for metricName, metricFamily := range parsed {
		for _, metric := range metricFamily.GetMetric() {
			fmt.Printf("Metric name: %s\n", metricName)
			fmt.Printf("Metric value: %v\n", metric.GetCounter().GetValue())

			for _, label := range metric.GetLabel() {
				fmt.Printf("Label: %s Value: %s\n", label.GetName(), label.GetValue())
			}
			fmt.Println()
		}
	}
}


