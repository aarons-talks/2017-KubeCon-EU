package tpr

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
)

const (
	// the name of the "Kind" field in new Backup resources
	backupTPRName = "Backup"
	// the value for the path in URLs to fetch Backup resources
	backupURLName = "backups"
	// GroupName is the group name for the backup TPR
	GroupName = "kubeconeu.deis.com"
	// Version is the version of the backup TPR
	Version = "v1alpha1"
)

// Backup is the TPR for backing up a cluster's state
type Backup struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata"`
	ResourceType         string `json:"resource-type"`
	Status               string `json:"status"`
}
