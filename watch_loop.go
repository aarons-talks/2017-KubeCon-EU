package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/watch"
)

func runWatchLoop(
	logger *log.Logger,
	cl *kubernetes.Clientset,
	openWatcher func() (watch.Interface, error),
) error {
	// loop forever. inside of each loop iteration, open a new watcher, get its watch chan, and
	// listen on that chan
	for {
		watcher, err := openWatcher()
		if err != nil {
			return err
		}
		defer watcher.Stop()
		watchCh := watcher.ResultChan()
		// receive on watchCh until it is closed
		for evt := range watchCh {
			logger.Printf("received event %#v", evt)
		}
		log.Printf("Watch channel closed, reopening...")
	}
	return nil
}
