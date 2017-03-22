package main

import (
	"github.com/arschles/2017-KubeCon-EU/tpr"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/rest"
)

func newDynamicClient(cfg rest.Config) (*dynamic.Client, error) {
	cfg.ContentConfig.GroupVersion = &unversioned.GroupVersion{
		Group:   tpr.GroupName,
		Version: tpr.Version,
	}
	cfg.APIPath = "apis"
	return dynamic.NewClient(&cfg)
}
