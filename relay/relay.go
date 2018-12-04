package relay

import (
	"email-bot/datastructures"
	"email-bot/lg"
	"time"
)

//Relay keeps listening for signals coming in from scrapes.
//Relay sends each signal on to be processed depending on the signals name.
type Relay struct {
	incoming <-chan *datastructures.Chunk
}

// Listen starts listening for signals and processing them depending
// on their name in a new goroutine.
func (r *Relay) Listen() {
	go r.startListening()
}

func (r *Relay) startListening() {
	for chunk := range r.incoming {
		switch chunk.Name {
		case "visit-link":
			links := chunk.Value.([]string)
			lg.Debug("linbk: %v", links)
		default:
			time.Sleep(time.Millisecond * 3000)
			lg.Debug("No signals received for %v milliseconds", 3000)
		}
	}
}

func NewRelay(incomingChannel chan *datastructures.Chunk) *Relay {
	return &Relay{
		incoming: incomingChannel,
	}
}
