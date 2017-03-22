package tpr

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
)

const (
	backupTPRName = "Backup"
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
