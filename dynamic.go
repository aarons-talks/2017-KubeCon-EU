package main

import (
	"github.com/arschles/2017-KubeCon-EU/tpr"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/rest"
)

// constructs a new dynamic.Client from the given rest config. The dynamic client is great
// for interacting with arbitrary Kubernetes resources for which the client doesn't have
// (generated) code to encode and decode the response bodies. Third Party Resources are great
// use cases for this type of client
func newDynamicClient(cfg rest.Config) (*dynamic.Client, error) {
	cfg.ContentConfig.GroupVersion = &unversioned.GroupVersion{
		Group:   tpr.GroupName,
		Version: tpr.Version,
	}
	cfg.APIPath = "apis"
	return dynamic.NewClient(&cfg)
}
