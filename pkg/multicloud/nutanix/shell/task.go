
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"
)

func init() {
	type TaskListOptions struct {
	}
	shellutils.R(&TaskListOptions{}, "task-list", "list task", func(cli *nutanix.SRegion, args *TaskListOptions) error {
		tasks, err := cli.GetTasks()
		if err != nil {
			return err
		}
		printList(tasks, 0, 0, 0, []string{})
		return nil
	})

	type TaskIdOptions struct {
		ID string
	}

	shellutils.R(&TaskIdOptions{}, "task-show", "show task", func(cli *nutanix.SRegion, args *TaskIdOptions) error {
		task, err := cli.GetTask(args.ID)
		if err != nil {
			return err
		}
		printObject(task)
		return nil
	})

}
