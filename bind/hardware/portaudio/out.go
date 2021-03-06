// Package portaudio is for modular binding of ontomix to audio interface via PortAudio
package portaudio

import (
	"github.com/gordonklaus/portaudio"

	"github.com/go-ontomix/ontomix/bind/sample"
	"github.com/go-ontomix/ontomix/bind/spec"
)

var outPortaudioStream *portaudio.Stream

func ConfigureOutput(s spec.AudioSpec) {
	var err error
	outSpec = &s
	portaudio.Initialize()
	outPortaudioStream, err = portaudio.OpenDefaultStream(0, s.Channels, s.Freq, 0, outPortaudioStreamCallback)
	noErr(err)
	noErr(outPortaudioStream.Start())
}

func TeardownOutput() {
	//	noErr(out.Stop())
	//	noErr(out.Close())
	portaudio.Terminate()
}

/*
 *
 private */

var (
	outSpec *spec.AudioSpec
)

func outPortaudioStreamCallback(out [][]float32) {
	var smp []sample.Value
	for s := range out[0] {
		smp = sample.OutNext()
		for c := 0; c < outSpec.Channels; c++ {
			out[c][s] = float32(smp[c])
		}
	}
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
