
package shell

import (
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type BucketListOptions struct {
		MaxResults int
		PageToken  string
	}

	shellutils.R(&BucketListOptions{}, "bucket-list", "List buckets", func(cli *google.SRegion, args *BucketListOptions) error {
		buckets, err := cli.GetBuckets(args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(buckets, 0, 0, 0, nil)
		return nil
	})

	type BucketCreateOptions struct {
		NAME         string
		StorageClass string `choices:"STANDARD|NEARLINE|COLDLINE"`
		Acl          string `choices:"private|authenticated-read|public-read|public-read-write"`
	}

	shellutils.R(&BucketCreateOptions{}, "bucket-create", "Create buckets", func(cli *google.SRegion, args *BucketCreateOptions) error {
		bucket, err := cli.CreateBucket(args.NAME, args.StorageClass, cloudprovider.TBucketACLType(args.Acl))
		if err != nil {
			return err
		}
		printObject(bucket)
		return nil
	})

	type BucketNameOptions struct {
		NAME string
	}

	shellutils.R(&BucketNameOptions{}, "bucket-show", "Show bucket", func(cli *google.SRegion, args *BucketNameOptions) error {
		bucket, err := cli.GetBucket(args.NAME)
		if err != nil {
			return err
		}
		printObject(bucket)
		return nil
	})

	shellutils.R(&BucketNameOptions{}, "bucket-delete", "Delete bucket", func(cli *google.SRegion, args *BucketNameOptions) error {
		return cli.DeleteBucket(args.NAME)
	})

	shellutils.R(&BucketNameOptions{}, "bucket-acl-list", "Show bucket acls", func(cli *google.SRegion, args *BucketNameOptions) error {
		acls, err := cli.GetBucketAcl(args.NAME)
		if err != nil {
			return err
		}
		printList(acls, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&BucketNameOptions{}, "bucket-iam-show", "Show bucket iam", func(cli *google.SRegion, args *BucketNameOptions) error {
		iam, err := cli.GetBucketIam(args.NAME)
		if err != nil {
			return err
		}
		printObject(iam)
		return nil
	})

	type BucketAclOptions struct {
		BUCKET string
		ACL    string
	}

	shellutils.R(&BucketAclOptions{}, "bucket-acl-set", "Set bucket acl", func(cli *google.SRegion, args *BucketAclOptions) error {
		return cli.SetBucketAcl(args.BUCKET, cloudprovider.TBucketACLType(args.ACL))
	})

	type BucketCORSListOptions struct {
		ID string
	}
	shellutils.R(&BucketCORSListOptions{}, "bucket-cors-list", "List all cors bucket", func(cli cloudprovider.ICloudRegion, args *BucketCORSListOptions) error {
		bucket, err := cli.GetIBucketById(args.ID)
		if err != nil {
			return errors.Wrap(err, "GetIBucketById")
		}
		cors, err := bucket.GetCORSRules()
		if err != nil {
			return errors.Wrap(err, "GetCORSRules")
		}
		printObject(cors)
		return nil
	})

	shellutils.R(&BucketCORSListOptions{}, "bucket-cors-delete", "delete all cors bucket", func(cli cloudprovider.ICloudRegion, args *BucketCORSListOptions) error {
		bucket, err := cli.GetIBucketById(args.ID)
		if err != nil {
			return errors.Wrap(err, "GetIBucketById")
		}
		err = bucket.DeleteCORS()
		if err != nil {
			return errors.Wrap(err, "DeleteCORS")
		}
		return nil
	})
}
