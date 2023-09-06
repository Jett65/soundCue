package main

import (
	// "fmt"
	// "github.com/Jett65/soundQueue/sound"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
    f, err := os.Open("Tockafeller Skank.mp3")
    if err != nil {
        log.Fatal(err)
    }

    streamer, format, err := mp3.Decode(f)
    if err != nil {
        log.Fatal(err)
    }

    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

    speaker.Play(streamer)
}
