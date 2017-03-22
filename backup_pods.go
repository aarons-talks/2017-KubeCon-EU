package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

func backupPods(cl *kubernetes.Clientset) error {
	pods, err := cl.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for i, pod := range pods.Items {
		log.Printf("-----\npod %d:\n%#v\n\n", i, pod)
	}
	return nil
}
