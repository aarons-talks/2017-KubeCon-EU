package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/watch"
)

func runWatchLoop(store storage, openWatcher func() (watch.Interface, error)) error {
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
			log.Printf("got event %#v", evt)
			if err := store.Append(evt); err != nil {
				log.Printf("Error appending event %#v (%s)", evt, err)
			}
		}
		log.Printf("Watch channel closed, reopening...")
	}
	return nil
}

func openPodsWatcher(cl *kubernetes.Clientset, ns string) func() (watch.Interface, error) {
	return func() (watch.Interface, error) {
		return cl.Core().Pods(namespace).Watch(apiv1.ListOptions{})
	}
}
