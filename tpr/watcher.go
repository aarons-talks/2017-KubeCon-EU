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
		// get the resource client for the Backup TPR in the given namespace, then start
		// watching for backups
		iface, err := cl.Resource(resource, ns).Watch(&Backup{})
		if err != nil {
			return nil, err
		}
		// run all watch events through a filter before returning them.
		// see the below watchFilterer function for info on what we're doing here
		return watch.Filter(iface, watchFilterer()), nil
	}
}

// this is a function that translates watch events for Backup TPRs so that consumers of
// the watch channel can cast objects in each event directly to *Backups
func watchFilterer() func(watch.Event) (watch.Event, bool) {
	return func(in watch.Event) (watch.Event, bool) {
		// event objects for TPRs come in as *runtime.Unstructured - a 'bucket' of unknown bytes
		// that Kubernetes converts to a map[string]interface{} for us
		unstruc, ok := in.Object.(*runtime.Unstructured)
		if !ok {
			log.Printf("Not an unstructured")
			return in, false
		}
		// marshal the map into JSON, then unmarshal it back into a *Backup so that we can return
		// it. If we fail anywhere, then indicate to the watch stream not to process this event.
		// the consumer of the watch stream won't care about the event.
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
