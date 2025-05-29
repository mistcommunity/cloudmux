package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type ContainerClusterListOptions struct {
	}
	shellutils.R(&ContainerClusterListOptions{}, "container-cluster-list", "List container clusters", func(cli *aws.SRegion, args *ContainerClusterListOptions) error {
		clusters, err := cli.ListClusters()
		if err != nil {
			return err
		}
		fmt.Println(clusters)
		return nil
	})

	type ContainerServiceListOptions struct {
		CLUSTER string
	}
	shellutils.R(&ContainerServiceListOptions{}, "container-service-list", "List cluster services", func(cli *aws.SRegion, args *ContainerServiceListOptions) error {
		services, err := cli.ListServices(args.CLUSTER)
		if err != nil {
			return err
		}
		fmt.Println(services)
		return nil
	})

	type ContainerTaskListOptions struct {
		CLUSTER string
	}
	shellutils.R(&ContainerTaskListOptions{}, "container-task-list", "List cluster tasks", func(cli *aws.SRegion, args *ContainerTaskListOptions) error {
		tasks, err := cli.ListTasks(args.CLUSTER)
		if err != nil {
			return err
		}
		fmt.Println(tasks)
		return nil
	})

}
