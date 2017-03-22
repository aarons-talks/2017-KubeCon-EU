package tpr

import (
	"encoding/json"
	"log"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/runtime"
	"k8s.io/client-go/pkg/watch"
)

// NewBackupWatcher returns a function that watches all Backup TPRs in a given namespace
func NewBackupWatcher(cl *dynamic.Client, ns string) func() (watch.Interface, error) {
	return func() (watch.Interface, error) {
		resource := &unversioned.APIResource{
			Name:       backupURLName,
			Namespaced: true,
			Kind:       backupTPRName,
		}
		iface, err := cl.Resource(resource, ns).Watch(&Backup{})
		if err != nil {
			return nil, err
		}
		return watch.Filter(iface, watchFilterer()), nil
	}
}

func watchFilterer() func(watch.Event) (watch.Event, bool) {
	return func(in watch.Event) (watch.Event, bool) {
		unstruc, ok := in.Object.(*runtime.Unstructured)
		if !ok {
			log.Printf("Not an unstructured")
		}
		b, err := json.Marshal(unstruc.Object)
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
