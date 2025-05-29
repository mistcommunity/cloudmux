
package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type ProjectListOptions struct {
		MaxResults int
		PageToken  string
	}
	shellutils.R(&ProjectListOptions{}, "project-list", "List projects", func(cli *google.SRegion, args *ProjectListOptions) error {
		projects, err := cli.GetClient().GetProjects()
		if err != nil {
			return err
		}
		printList(projects, 0, 0, 0, nil)
		return nil
	})
}
