package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

func backupPods(cl *kubernetes.Clientset) error {
	pods, err := cl.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for i, pod := range pods {
		log.Printf("pod %d:\n%#v", i, pod)
	}
	return nil
}
