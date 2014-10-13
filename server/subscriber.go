package server

import (
	"fmt"
	"os"

	"gopkg.in/fsnotify.v1"
)

// Subscriber is a fsnotify.Watcher wrapper for recursive watching
type Subscriber struct {
	*fsnotify.Watcher
	Events chan fsnotify.Event
}

func NewSubscriber() *Subscriber {
	w, _ := fsnotify.NewWatcher()
	s := &Subscriber{
		Watcher: w,
	}

	s.Events = make(chan fsnotify.Event, 1024)
	go s.eventFilter()

	return s
}

func (s *Subscriber) RecursiveAdd(name string) error {
	err := s.Add(name)
	if err != nil {
		return err
	}

	// FIXME: this cause "too many open files"
	// filepath.Walk(name, s.visit)

	return nil
}

func (s *Subscriber) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
	}

	// FIXME: this cause "too many open files"
	// if f.IsDir() {
	// 	err = s.Add(path)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	return err
}

func (s *Subscriber) eventFilter() {
	for {
		select {
		case ev := <-s.Watcher.Events:
			fmt.Println(ev)
			if ev.Op == fsnotify.Create {
				s.Add(ev.Name)
			}
			s.Events <- ev
		}
	}
}
