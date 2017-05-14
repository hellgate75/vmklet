package contrib

import "vmklet/model"

func NewParser(format model.FormatType) (model.Parser, error) {
	return model.Parser{
		Format: format.String(),
	}, nil
}
