package lib

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func GetContainerInfo(client *containerd.Client, namespace, containerId string) {
	ctx := namespaces.WithNamespace(context.Background(), namespace)

	container, err := client.LoadContainer(ctx, containerId)
	if err != nil {
		fmt.Printf("컨테이너 를 찾을 수 없습니다 %v", err)
	}

	info, err := container.Info(ctx)
	if err != nil {
		fmt.Printf("컨테이너 정보를 가져오는데 실패했습니다: %v", err)
	}

	fmt.Printf("컨테이너 ID: %s\n", info.ID)
	fmt.Printf("컨테이너 이미지: %s\n", info.Image)
	fmt.Printf("컨테이너 레이블: %s\n", info.ID)
	for k, v := range info.Labels {
		fmt.Printf("\t%s: %s\n", k, v)
	}
}
