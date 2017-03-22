package main

import (
	"encoding/json"
	"fmt"
	"log"

	"k8s.io/client-go/pkg/watch"
)

type eventStorage interface {
	Append(evt watch.Event) error
	Flush() error
}

type loggerEventStorage struct {
	logger log.Logger
}

func (w *loggerEventStorage) Append(evt watch.Event) error {
	w.logger.Printf("received event %#v", evt)
	return nil
}

func (w *loggerEventStorage) Flush() error {
	return nil
}

func encodeEvt(evt watch.Event) (string, error) {
	objStr, err := json.Marshal(evt.Object)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", evt.Type, objStr), nil
}
