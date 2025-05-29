
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type ImageListOptions struct {
		Project    string `choices:"centos-cloud|ubuntu-os-cloud|windows-cloud|windows-sql-cloud|suse-cloud|suse-sap-cloud|rhel-cloud|rhel-sap-cloud|cos-cloud|debian-cloud"`
		MaxResults int
		PageToken  string
	}
	shellutils.R(&ImageListOptions{}, "image-list", "List images", func(cli *google.SRegion, args *ImageListOptions) error {
		images, err := cli.GetImages(args.Project, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(images, 0, 0, 0, nil)
		return nil
	})

	type ImageShowOptions struct {
		ID string
	}
	shellutils.R(&ImageShowOptions{}, "image-show", "Show image", func(cli *google.SRegion, args *ImageShowOptions) error {
		image, err := cli.GetImage(args.ID)
		if err != nil {
			return err
		}
		printObject(image)
		return nil
	})

	type ImageCreateOptions struct {
		NAME   string
		Desc   string
		BUCKET string
		FILE   string
	}

	shellutils.R(&ImageCreateOptions{}, "image-create", "Create image", func(cli *google.SRegion, args *ImageCreateOptions) error {
		image, err := cli.CreateImage(args.NAME, args.Desc, args.BUCKET, args.FILE)
		if err != nil {
			return err
		}
		printObject(image)
		return nil
	})
}
