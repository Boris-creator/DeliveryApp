package events

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
)

type Event[T any] struct {
	Name    string
	Payload T
}

type Listeners struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	channels map[string][]chan Event[any]
}

type ListenOptions struct {
	Once bool
}

var UnknownEventError = errors.New("unregistered event type")

var DefaultListeners = New()

func New() *Listeners {
	return &Listeners{channels: make(map[string][]chan Event[any])}
}

func (l *Listeners) AddListener(
	ctx context.Context,
	ev string,
	handler func(Event[any]),
	options ...func(opts *ListenOptions),
) (removeListener func()) {
	var opts ListenOptions
	for _, opt := range options {
		opt(&opts)
	}

	ch := make(chan Event[any])

	l.mu.Lock()
	defer l.mu.Unlock()
	l.channels[ev] = append(l.channels[ev], ch)

	go func() {
		for {
			select {
			case e, ok := <-ch:
				if !ok {
					break
				}

				l.wg.Done()
				handler(e)

				if opts.Once {
					l.removeListener(ev, ch)
				}
			case <-ctx.Done():
				break
			}
		}
	}()

	return func() {
		l.removeListener(ev, ch)
	}
}

func (l *Listeners) Listen() {
	l.wg.Wait()
}

func (l *Listeners) Dispatch(e string, payload any) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	event := Event[any]{
		Name:    e,
		Payload: payload,
	}

	chans, ok := l.channels[event.Name]
	if !ok {
		return fmt.Errorf("%w: %s", UnknownEventError, event.Name)
	}

	l.wg.Add(len(chans))

	for _, ch := range chans {
		ch <- event
	}

	return nil
}

func (l *Listeners) removeListener(ev string, ch chan Event[any]) {
	l.mu.Lock()
	defer l.mu.Unlock()

	close(ch)

	l.channels[ev] = slices.DeleteFunc(l.channels[ev], func(c chan Event[any]) bool {
		return c == ch
	})
	if len(l.channels[ev]) == 0 {
		delete(l.channels, ev)
	}
}
