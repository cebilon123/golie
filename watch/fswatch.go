package watch

import (
	"io/fs"
	"io/ioutil"
)

//Watcher watches for changes in directory
type Watcher struct {
	Events         chan Event
	Errors         chan error
	watchDirectory string
	stopCh         chan struct{}
}

func (w *Watcher) Close() error {
	go func() {
		close(w.Errors)
		close(w.stopCh)
	}()

	return nil
}

type Event struct {
	Type   EventType
	Caller string
}

type EventType uint8

const (
	CreateEventType EventType = iota
	UpdateEventType
	RemoveEventType
)

//NewWatcher creates new watcher which watches given directory.
func NewWatcher(directory string) *Watcher {
	return &Watcher{
		Events:         make(chan Event, 1),
		Errors:         make(chan error, 1),
		watchDirectory: directory,
	}
}

func (w *Watcher) notify(event Event) {
	select {
	case <-w.stopCh:
		return
	case w.Events <- event:
	}
}

//Watch starts watching of directory. If Close is called, Watch is going to stop watching.
func (w *Watcher) Watch() {
	set := make(map[fs.FileInfo]bool)
	for {
		select {
		//If watcher is closed it should stop watching directory
		case <-w.stopCh:
			return
		default:
			fis, err := ioutil.ReadDir(w.watchDirectory)
			if err != nil {
				return
			}

			//TODO add more advanced watching of changes
			//If there is no file , add it
			for _, fi := range fis {
				if _, ok := set[fi]; !ok {
					set[fi] = true
					w.notify(Event{
						Type:   UpdateEventType,
						Caller: fi.Name(),
					})
				}
			}
		}
	}
}
