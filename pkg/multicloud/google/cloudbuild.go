
package google

import "yunion.io/x/pkg/errors"

type SCloudbuildBuild struct {
	Id     string
	Status string
	LogUrl string
}

type SCloudbuildMetadata struct {
	Build SCloudbuildBuild
}

type SCloudbuildOperation struct {
	Name     string
	Metadata SCloudbuildMetadata
}

func (region *SRegion) GetCloudbuildOperation(name string) (*SCloudbuildOperation, error) {
	operation := SCloudbuildOperation{}
	err := region.cloudbuildGet(name, &operation)
	if err != nil {
		return nil, errors.Wrap(err, "region.cloudbuildGet")
	}
	return &operation, nil
}
