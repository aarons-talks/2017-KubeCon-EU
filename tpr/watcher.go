package tpr

import (
	"encoding/json"
	"log"

	"k8s.io/client-go/pkg/watch"
	"k8s.io/client-go/rest"
)

// NewBackupWatcher returns a function that watches all Backup TPRs in a given namespace
func NewBackupWatcher(restClient rest.Interface, ns string) func() (watch.Interface, error) {
	return func() (watch.Interface, error) {
		req := restClient.Get().AbsPath(
			"apis",
			groupName,
			tprVersion,
			"watch",
			"namespaces",
			ns,
			backupURLName,
		)
		watchIface, err := req.Watch()
		if err != nil {
			return nil, err
		}
		return watch.Filter(watchIface, watchFilterer(ns)), nil
	}
}

func watchFilterer(ns string) func(watch.Event) (watch.Event, bool) {
	return func(in watch.Event) (watch.Event, bool) {
		b, err := json.Marshal(in.Object)
		if err != nil {
			log.Printf("Error marshaling %#v (%s)", in.Object, err)
			return in, false
		}
		backup := new(Backup)
		if err := json.Unmarshal(b, backup); err != nil {
			log.Printf("Error unmarshaling %s (%s)", string(b), err)
			return in, false
		}
		in.Object = backup
		return in, true
	}
}
