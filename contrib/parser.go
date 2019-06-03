package contrib

import (
	"github.com/hellgate75/vmklet/model"
)

func NewParser(format model.FormatType) (model.Parser, error) {
	return model.Parser{
		Format: format,
	}, nil
}
