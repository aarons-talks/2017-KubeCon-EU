package main

import (
	"log"
	"os"

	"github.com/arschles/2017-KubeCon-EU/tpr"
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

	dynCl, err := newDynamicClient(*cfg)
	if err != nil {
		log.Printf("Error creating a dynamic client (%s)", err)
		os.Exit(1)
	}
	log.Printf("watching namespace %s for backup TPRs", namespace)
	backupTPRWatchFunc := tpr.NewBackupWatcher(dynCl, namespace)
	if err := runWatchLoop(cl, backupTPRWatchFunc); err != nil {
		log.Fatalf("error running watch loop (%s)", err)
	}
}
