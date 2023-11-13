package main
import (
        "strings"
        "context"
        "fmt"
        "k8s.io/client-go/kubernetes"
        k8srest "k8s.io/client-go/rest"
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
        fmt.Println(reader)
}
