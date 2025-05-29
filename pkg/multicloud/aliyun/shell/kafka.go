
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type KafkaListOptions struct {
		Ids []string
	}
	shellutils.R(&KafkaListOptions{}, "kafka-list", "List kafka", func(cli *aliyun.SRegion, args *KafkaListOptions) error {
		kafkas, err := cli.GetKafkas(args.Ids)
		if err != nil {
			return err
		}
		printList(kafkas, 0, 0, 0, []string{})
		return nil
	})

	type KafkaIdOptions struct {
		ID string
	}

	shellutils.R(&KafkaIdOptions{}, "kafka-delete", "Delete kafka", func(cli *aliyun.SRegion, args *KafkaIdOptions) error {
		return cli.DeleteKafka(args.ID)
	})

	shellutils.R(&KafkaIdOptions{}, "kafka-release", "Release kafka", func(cli *aliyun.SRegion, args *KafkaIdOptions) error {
		return cli.ReleaseKafka(args.ID)
	})

	type KafkaTopicListOptions struct {
		KafkaIdOptions
		Page     int
		PageSize int
	}

	shellutils.R(&KafkaTopicListOptions{}, "kafka-topic-list", "List kafka topic", func(cli *aliyun.SRegion, args *KafkaTopicListOptions) error {
		topics, _, err := cli.GetKafkaTopics(args.ID, args.Page, args.PageSize)
		if err != nil {
			return err
		}
		printList(topics, 0, 0, 0, nil)
		return nil
	})

}
