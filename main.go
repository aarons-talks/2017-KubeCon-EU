package main

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Error configuring the Kube client (%s)", err)
		os.Exit(1)
	}

	cl, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Printf("Error creating a new Kube client (%s)", err)
		os.Exit(1)
	}

	store := &writerStorage{writer: os.Stdout}
	podIface := cl.Core().Pods("")
	log.Fatal(runWatchLoop(podIface, store))
}
