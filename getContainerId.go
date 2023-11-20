package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
)

func main() {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	containers, err := client.Containers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID())
	}
}
