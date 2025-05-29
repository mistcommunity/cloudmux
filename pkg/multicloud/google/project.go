
package google

import (
	"time"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SProject struct {
	multicloud.SProjectBase
	GoogleTags
	Name           string
	CreateTime     time.Time
	LifecycleState string
	ProjectId      string
	ProjectNumber  string
}

func (cli *SGoogleClient) GetProject(id string) (*SProject, error) {
	project := &SProject{}
	resp, err := cli.managerGet(id)
	if err != nil {
		return nil, errors.Wrap(err, "managerGet")
	}
	err = resp.Unmarshal(project)
	if err != nil {
		return nil, errors.Wrap(err, "resp.Unmarshal")
	}
	return project, nil
}

func (cli *SGoogleClient) GetProjects() ([]SProject, error) {
	nextPageToken := ""
	params := map[string]string{}
	result := []SProject{}
	for {
		if len(nextPageToken) > 0 {
			params["pageToken"] = nextPageToken
		}
		resp, err := cli.managerList("projects", params)
		if err != nil {
			return nil, errors.Wrap(err, "managerList")
		}
		_result := []SProject{}
		if resp.Contains("projects") {
			err = resp.Unmarshal(&_result, "projects")
			if err != nil {
				return nil, errors.Wrap(err, "data.Unmarshal")
			}
		}
		result = append(result, _result...)
		nextPageToken, _ = resp.GetString("nextPageToken")
		if len(nextPageToken) == 0 || len(_result) == 0 {
			break
		}
	}
	return result, nil
}

func (p *SProject) GetName() string {
	return p.Name
}

func (p *SProject) GetId() string {
	return p.ProjectId
}

func (p *SProject) GetGlobalId() string {
	return p.ProjectId
}

func (p *SProject) GetStatus() string {
	return api.EXTERNAL_PROJECT_STATUS_AVAILABLE
}

func (p *SProject) Refresh() error {
	return nil
}

func (p *SProject) IsEmulated() bool {
	return false
}
