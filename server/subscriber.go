package server

import "gopkg.in/fsnotify.v1"

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

	return nil
}

func (s *Subscriber) eventFilter() {
	for {
		select {
		case ev := <-s.Watcher.Events:
			s.Events <- ev
		}
	}
}
