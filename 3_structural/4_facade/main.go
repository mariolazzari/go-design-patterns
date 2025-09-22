package main

import "fmt"

// Subsystem Class
type Post interface {
	publish()
}

type InstagramPost struct {
}

func (i *InstagramPost) publish() {
	fmt.Println("Publishing your post to Instagram")
}

type TikTokPost struct {
}

func (t *TikTokPost) publish() {
	fmt.Println("Publishing your video to TikTok")
}

type TwitterPost struct {
}

func (t *TwitterPost) publish() {
	fmt.Println("Publishing your tweet to Twitter")
}

// Facade
type Publisher struct {
	instagram *InstagramPost
	tikTok    *TikTokPost
	twitter   *TwitterPost
}

func newPublisher() *Publisher {
	fmt.Println("Initializing Publisher...")
	publisher := &Publisher{
		instagram: &InstagramPost{},
		tikTok:    &TikTokPost{},
		twitter:   &TwitterPost{},
	}
	fmt.Println("...Publisher is ready")
	return publisher
}

func (p *Publisher) postToInstagram() {
	p.instagram.publish()
}

func (p *Publisher) postToTikTok() {
	p.tikTok.publish()
}

func (p *Publisher) postToTwitter() {
	p.twitter.publish()
}

func main() {
	publisher := newPublisher()

	publisher.postToInstagram()
	publisher.postToTikTok()
	publisher.postToTwitter()
}
