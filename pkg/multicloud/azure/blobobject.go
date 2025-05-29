
package azure

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/storage"

	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SObject struct {
	container *SContainer

	cloudprovider.SBaseCloudObject
}

func (o *SObject) GetIBucket() cloudprovider.ICloudBucket {
	return o.container.storageaccount
}

func (o *SObject) GetAcl() cloudprovider.TBucketACLType {
	return o.container.getAcl()
}

func (o *SObject) SetAcl(aclStr cloudprovider.TBucketACLType) error {
	return nil
}

func (o *SObject) getBlobName() string {
	if len(o.Key) <= len(o.container.Name)+1 {
		return ""
	} else {
		return o.Key[len(o.container.Name)+1:]
	}
}

func (o *SObject) getBlobRef() (*storage.Blob, error) {
	blobName := o.getBlobName()
	if len(blobName) == 0 {
		return nil, nil
	}
	contRef, err := o.container.getContainerRef()
	if err != nil {
		return nil, errors.Wrap(err, "src getContainerRef")
	}
	blobRef := contRef.GetBlobReference(blobName)
	return blobRef, nil
}

func (o *SObject) GetMeta() http.Header {
	if o.Meta != nil {
		return o.Meta
	}
	blobRef, err := o.getBlobRef()
	if err != nil {
		log.Errorf("o.getBlobRef fail %s", err)
		return nil
	}
	if blobRef == nil {
		return nil
	}
	err = blobRef.GetMetadata(nil)
	if err != nil {
		log.Errorf("blobRef.GetMetadata fail %s", err)
	}
	err = blobRef.GetProperties(nil)
	if err != nil {
		log.Errorf("blobRef.GetProperties fail %s", err)
	}
	meta := getBlobRefMeta(blobRef)
	o.Meta = meta
	return o.Meta
}

func (o *SObject) SetMeta(ctx context.Context, meta http.Header) error {
	blobRef, err := o.getBlobRef()
	if err != nil {
		return errors.Wrap(err, "o.getBlobRef")
	}
	if blobRef == nil {
		return cloudprovider.ErrNotSupported
	}
	propChanged, metaChanged := setBlobRefMeta(blobRef, meta)
	if propChanged {
		propOpts := storage.SetBlobPropertiesOptions{}
		err := blobRef.SetProperties(&propOpts)
		if err != nil {
			return errors.Wrap(err, "blob.SetProperties")
		}
	}
	if metaChanged {
		metaOpts := storage.SetBlobMetadataOptions{}
		err := blobRef.SetMetadata(&metaOpts)
		if err != nil {
			return errors.Wrap(err, "blob.SetMetadata")
		}
	}
	return nil
}
