package main

import (
	"fmt"
)

// Observer Interface
type Observer interface {
	Update(media []string)
}

// Subject
type MediaLibrary struct {
	media     []string
	observers []Observer
}

func (m *MediaLibrary) AddMedia(media string) {
	m.media = append(m.media, media)
	m.Notify()
}

func (m *MediaLibrary) Attach(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *MediaLibrary) Detach(observer Observer) {
	for i, o := range m.observers {
		if o == observer {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			break
		}
	}
}

func (m *MediaLibrary) Notify() {
	for _, o := range m.observers {
		o.Update(m.media)
	}
}

// Concrete Observer
type MediaPlayer struct{}

func (m *MediaPlayer) Update(media []string) {
	fmt.Printf("Playing media file: %s\n", media[len(media)-1])
}

func main() {
	library := &MediaLibrary{}
	player1 := &MediaPlayer{}
	player2 := &MediaPlayer{}

	library.Attach(player1)
	library.AddMedia("video.mp4")
	library.AddMedia("audio.mp3")

	library.Detach(player1)

	library.Attach(player2)
	library.AddMedia("video2.mp4")
	library.AddMedia("audio2.mp4")
}
