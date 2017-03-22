package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"k8s.io/client-go/pkg/watch"
)

type storage interface {
	Append(evt watch.Event) error
	Flush() error
}

type writerStorage struct {
	buf       []watch.Event
	maxBufLen int
	mut       *sync.RWMutex
	writer    io.Writer
}

func (w *writerStorage) Append(evt watch.Event) error {
	w.mut.Lock()
	defer w.mut.Unlock()
	w.buf = append(w.buf, evt)
	if len(w.buf) >= w.maxBufLen {
		flushErr := w.Flush()
		w.buf = []watch.Event{}
		return flushErr
	}
	return nil
}

func (w *writerStorage) Flush() error {
	w.mut.Lock()
	defer w.mut.Unlock()
	for i, evt := range w.buf {
		var outStr string
		encoded, err := encodeEvt(evt)
		if err != nil {
			outStr = fmt.Sprintf("error encoding event %d (%s)", i, err)
		} else {
			outStr = encoded
		}
		if _, err := fmt.Fprintf(w.writer, "%s %s", evt.Type, outStr); err != nil {
			log.Printf("Error outputting event %d (%s)", i, err)
		}
	}
	w.buf = []watch.Event{}
	return nil
}

func encodeEvt(evt watch.Event) (string, error) {
	objStr, err := json.Marshal(evt.Object)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", evt.Type, objStr), nil
}
