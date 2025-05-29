package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type KinesisListOptions struct {
	}
	shellutils.R(&KinesisListOptions{}, "kinesis-stream-list", "List kinesis stream", func(cli *aws.SRegion, args *KinesisListOptions) error {
		streams, err := cli.ListStreams()
		if err != nil {
			return err
		}
		printList(streams, 0, 0, 0, []string{})
		return nil
	})

	type KinesisNameOptions struct {
		NAME string
	}
	shellutils.R(&KinesisNameOptions{}, "kinesis-stream-show", "Show kinesis stream", func(cli *aws.SRegion, args *KinesisNameOptions) error {
		stream, err := cli.DescribeStream(args.NAME)
		if err != nil {
			return err
		}
		printObject(stream)
		return nil
	})

}
