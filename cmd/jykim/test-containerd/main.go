package jykim

import (
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"my.com/lib"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("사용법: %s <네임스페이스> <컨테이너 ID>", os.Args[0])
	}
	namespace := os.Args[1]
	containerId := os.Args[2]

	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		fmt.Printf("에러:", err)
	}
	defer client.Close()
	lib.GetContainerInfo(client, namespace, containerId)
}
