package contrib

import (
	"errors"
	"fmt"
	"github.com/hellgate75/vmklet/model"
	"github.com/hellgate75/vmklet/streams"
)

func NewStream(streamType model.StreamType) (model.CommandStream, error) {
	switch streamType {
	case model.SSHStream:
		return &streams.SSHCommandStream{Type: streamType}, nil
	}
	return nil, errors.New(fmt.Sprintf("Stream type '%s' not provided!!", streamType.String()))
}
