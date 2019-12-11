package main

import (
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	if _, err := cli.Info(context.Background()); err != nil {
		log.Fatalln("Failed to test docker client")
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.Names[0])
	}
}
