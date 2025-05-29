package aws

func (self *SRegion) ListClusters() ([]string, error) {
	params := map[string]interface{}{
		"maxResults": 100,
	}
	ret := []string{}
	for {
		part := struct {
			ClusterArns []string
			NextToken   string
		}{}
		err := self.ecsRequest("ListClusters", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.ClusterArns...)
		if len(part.ClusterArns) == 0 || len(part.NextToken) == 0 {
			break
		}
		params["nextToken"] = part.NextToken
	}
	return ret, nil
}

func (self *SRegion) ListServices(cluster string) ([]string, error) {
	params := map[string]interface{}{
		"cluster":    cluster,
		"maxResults": 100,
	}
	ret := []string{}
	for {
		part := struct {
			ServiceArns []string
			NextToken   string
		}{}
		err := self.ecsRequest("ListServices", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.ServiceArns...)
		if len(part.ServiceArns) == 0 || len(part.NextToken) == 0 {
			break
		}
		params["nextToken"] = part.NextToken
	}
	return ret, nil
}

func (self *SRegion) ListTasks(cluster string) ([]string, error) {
	params := map[string]interface{}{
		"cluster":    cluster,
		"maxResults": 100,
	}
	ret := []string{}
	for {
		part := struct {
			TaskArns  []string
			NextToken string
		}{}
		err := self.ecsRequest("ListTasks", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.TaskArns...)
		if len(part.TaskArns) == 0 || len(part.NextToken) == 0 {
			break
		}
		params["nextToken"] = part.NextToken
	}
	return ret, nil
}
