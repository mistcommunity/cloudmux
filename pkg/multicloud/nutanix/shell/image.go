
package shell

import (
	"os"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type ImageListOptions struct {
	}
	shellutils.R(&ImageListOptions{}, "image-list", "list hosts", func(cli *nutanix.SRegion, args *ImageListOptions) error {
		images, err := cli.GetImages()
		if err != nil {
			return err
		}
		printList(images, 0, 0, 0, []string{})
		return nil
	})

	type ImageIdOptions struct {
		ID string
	}

	shellutils.R(&ImageIdOptions{}, "image-show", "show host", func(cli *nutanix.SRegion, args *ImageIdOptions) error {
		image, err := cli.GetImage(args.ID)
		if err != nil {
			return err
		}
		printObject(image)
		return nil
	})

	type ImageUploadOptions struct {
		STOREG_ID string
		NAME      string
		FILE      string
	}

	shellutils.R(&ImageUploadOptions{}, "image-upload", "upload host", func(cli *nutanix.SRegion, args *ImageUploadOptions) error {
		fi, err := os.Open(args.FILE)
		if err != nil {
			return err
		}
		defer fi.Close()

		stat, _ := fi.Stat()
		image, err := cli.CreateImage(
			args.STOREG_ID,
			&cloudprovider.SImageCreateOption{
				ImageName: args.NAME,
			},
			stat.Size(),
			fi,
			nil,
		)
		if err != nil {
			return err
		}
		printObject(image)
		return nil
	})

}
