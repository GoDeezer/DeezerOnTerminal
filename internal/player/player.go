package player

import (
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/godeezer/lib/deezer"
)

type Player struct {
	DeezerClient *deezer.Client
	ReDraw       chan struct{}

	PlayerQueue *PlayerQueue
	CurrentSong deezer.Song
	Progress    float64

	SampleRate beep.SampleRate
	Streamer   beep.StreamSeeker
	Ctrl       *beep.Ctrl
	Resampler  *beep.Resampler
	Volume     *effects.Volume
}

func NewPlayer(client *deezer.Client, redraw chan struct{}) *Player {
	return &Player{
		DeezerClient: client,
		ReDraw:       redraw,
		PlayerQueue:  NewPlayerQueue(client),
	}
}

func (self *Player) SetCurrentSong(song deezer.Song) {
	self.CurrentSong = song
	self.Progress = 0.0
}

func (self *Player) Play() {
	go func() {
		reader, err := self.DeezerClient.Download(self.CurrentSong, deezer.MP3128)
		if err != nil {
			panic(err)
		}

		streamer, format, err := mp3.Decode(reader)
		if err != nil {
			panic(err)
		}
		defer streamer.Close()

		done := make(chan struct{})
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/3))
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- struct{}{}
		})))

		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				self.Progress += 0.01
				self.ReDraw <- struct{}{}
			case <-done:
				return
			}
		}
	}()
}
