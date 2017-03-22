package tpr

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
)

const (
	backupURLName = "backups"
	groupName     = "kubeconeu.deis.com"
	tprVersion    = "v1alpha1"
)

// Backup is the TPR for backing up a cluster's state
type Backup struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata"`
	ResourceType         string `json:"resource-type"`
	Status               string `json:"status"`
}
