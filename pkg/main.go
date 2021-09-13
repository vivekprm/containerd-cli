package main

import (
	"context"
	"fmt"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"log"
	"os"
)

func main() {
	// if err := redisExample(); err != nil {
	// 	log.Fatal(err)
	// }
	if err := importImages(); err != nil {
		log.Fatal(err)
	}
}
func importImages() error {
	cli, err := containerd.New("/run/k3s/containerd/containerd.sock")
	// cli, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer cli.Close()
	file, err := os.Open("qotm.tar")
	if err != nil {
		fmt.Println("extractTarFile: error in opening the tar ", err.Error())
		return err
	}
	defer file.Close()

	ctx := namespaces.WithNamespace(context.Background(), "k8s.io")
	images, err := cli.Import(ctx, file)
	if err != nil {
		fmt.Println("error in importing the images ", err.Error())
		return err
	}
	fmt.Println(images[0])
	return nil
}
func redisExample() error {
	client, err := containerd.New("/run/k3s/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "example")
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	if err != nil {
		return err
	}
	log.Printf("Successfully pulled %s image\n", image.Name())

	return nil
}
