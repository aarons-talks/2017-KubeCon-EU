package main

import (
	"log"
	"os"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	namespace = "default"
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

	store := &writerStorage{
		buf:       nil,
		maxBufLen: 2,
		mut:       new(sync.RWMutex),
		writer:    os.Stdout,
	}
	log.Printf("watching namespace %s", namespace)
	log.Fatal(runWatchLoop(store, openPodsWatcher(cl, namespace)))
}
