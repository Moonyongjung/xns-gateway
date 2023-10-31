package types

type Channels struct {
	HttpServerChan chan []byte
	ErrChan        chan error
}

func makeChannels() Channels {
	return Channels{
		HttpServerChan: make(chan []byte),
		ErrChan:        make(chan error),
	}
}
