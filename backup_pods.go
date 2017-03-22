package main

import (
	"encoding/json"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

// just fetch all pods, marshal them into JSON, and print them out
func backupPods(cl *kubernetes.Clientset) error {
	pods, err := cl.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return err
	}
	for i, pod := range pods.Items {
		b, err := json.Marshal(&pod)
		if err != nil {
			log.Printf("error formatting pod %d, continuing (%s)", i, err)
			continue
		}
		log.Printf("-----\nPod %d\n%s\n\n", i, string(b))
	}
	return nil
}
