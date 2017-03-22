package main

import (
	"log"

	"k8s.io/client-go/kubernetes/typed/core/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

func runWatchLoop(podIface v1.PodInterface, store storage) error {
	watcher, err := podIface.Watch(apiv1.ListOptions{})
	if err != nil {
		return err
	}
	defer watcher.Stop()
	watchCh := watcher.ResultChan()
	// loop forever. inside of each loop iteration, we watch on watchCh until it's closed
	for {
		// receive on watchCh until it is closed
		for evt := range watchCh {
			if err := store.Append(evt); err != nil {
				log.Printf("Error appending event %#v (%s)", evt, err)
			}
		}
	}
	return nil
}
