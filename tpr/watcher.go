package tpr

import (
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
		return req.Watch()
	}
}

// func watchFilterer(t *store, ns string) func(watch.Event) (watch.Event, bool) {
// 	return func(in watch.Event) (watch.Event, bool) {
// 		encodedBytes, err := runtime.Encode(t.codec, in.Object)
// 		if err != nil {
// 			glog.Errorf("couldn't encode watch event object (%s)", err)
// 			return watch.Event{}, false
// 		}
// 		finalObj := t.singularShell("", "")
// 		if err := decode(t.codec, nil, encodedBytes, finalObj); err != nil {
// 			glog.Errorf("couldn't decode watch event bytes (%s)", err)
// 			return watch.Event{}, false
// 		}
// 		if !t.hasNamespace {
// 			if err := removeNamespace(finalObj); err != nil {
// 				glog.Errorf("couldn't remove namespace from %#v (%s)", finalObj, err)
// 				return watch.Event{}, false
// 			}
// 		}
// 		return watch.Event{
// 			Type:   in.Type,
// 			Object: finalObj,
// 		}, true
// 	}
// }
