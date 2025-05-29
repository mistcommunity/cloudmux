package generic

import (
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type Operator[T cloudprovider.ICloudResource] struct {
	objects []T
}

func NewOperator[T cloudprovider.ICloudResource](nf func() ([]T, error)) (*Operator[T], error) {
	objs, err := nf()
	if err != nil {
		return nil, errors.Wrap(err, "new resources")
	}

	return &Operator[T]{
		objects: objs,
	}, nil
}

func (o Operator[T]) Iter(f func(T) error, continueOnErr bool) error {
	return Iter(o.objects, f, continueOnErr)
}
