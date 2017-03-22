package main

import (
	"log"
	"strings"

	"github.com/arschles/2017-KubeCon-EU/tpr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/watch"
)

func runWatchLoop(
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
			log.Printf("received event %#v", evt)
			backup, ok := evt.Object.(*tpr.Backup)
			if !ok {
				log.Printf("event was not a *Backup, skipping")
				continue
			}
			log.Printf("backing up all %s resources", backup.ResourceType)
			switch strings.ToLower(backup.ResourceType) {
			case "pod", "pods":
				if err := backupPods(cl); err != nil {
					log.Printf("error backing up all pods (%s)", err)
				}
			default:
				log.Printf("%s resources are not supported", backup.ResourceType)
			}
		}
		log.Printf("Watch channel closed, reopening...")
	}
	return nil
}
