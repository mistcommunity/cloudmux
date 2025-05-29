
package shell

import (
	"os"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
)

func init() {
	type ImageListOptions struct {
		NODE    string
		STORAGE string
	}

	shellutils.R(&ImageListOptions{}, "image-list", "list images", func(cli *proxmox.SRegion, args *ImageListOptions) error {
		images, err := cli.GetImages(args.NODE, args.STORAGE)
		if err != nil {
			return err
		}
		printList(images, 0, 0, 0, []string{})
		return nil
	})

	type ImageUploadOptions struct {
		NODE     string
		STORAGE  string
		FILENAME string
	}

	shellutils.R(&ImageUploadOptions{}, "image-upload", "upload image", func(cli *proxmox.SRegion, args *ImageUploadOptions) error {
		file, err := os.Open(args.FILENAME)
		if err != nil {
			return err
		}
		defer file.Close()
		image, err := cli.UploadImage(args.NODE, args.STORAGE, args.FILENAME, file)
		if err != nil {
			return err
		}
		printObject(image)
		return nil
	})

}
