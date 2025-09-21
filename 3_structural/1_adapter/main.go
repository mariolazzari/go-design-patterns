package main

import "fmt"

// Client
type Publisher struct {
}

func (p *Publisher) publishContentOnPlatform(platform Platform) {
	fmt.Println("Publisher is ready to publish your content.")
	platform.postMedia()
}

// Client Interface
type Platform interface {
	postMedia()
}

// Compatible Service
type Instagram struct {
}

func (i *Instagram) postMedia() {
	fmt.Println("Instagram has published your post.")
}

// Incompatible Service
type TikTok struct {
}

func (t *TikTok) scheduleMedia() {
	fmt.Println("TikTok is ready to schedule your post.")
}

// Adapter
type TikTokAdapter struct {
	tikTok *TikTok
}

func (t *TikTokAdapter) postMedia() {
	t.tikTok.scheduleMedia()
	fmt.Println("Adapter has posted the TikTok content.")
}

func main() {

	publisher := &Publisher{}
	instagram := &Instagram{}

	publisher.publishContentOnPlatform(instagram)

	tikTok := &TikTok{}
	tikTokAdapter := &TikTokAdapter{
		tikTok: tikTok,
	}

	publisher.publishContentOnPlatform(tikTokAdapter)
}
