package aws

type DynamoDB struct {
	AttributeDefinitions []struct {
		AttributeName string `json:"AttributeName"`
		AttributeType string `json:"AttributeType"`
	} `json:"AttributeDefinitions"`
	BillingModeSummary struct {
		BillingMode                       string  `json:"BillingMode"`
		LastUpdateToPayPerRequestDateTime float64 `json:"LastUpdateToPayPerRequestDateTime"`
	} `json:"BillingModeSummary"`
	CreationDateTime          float64 `json:"CreationDateTime"`
	DeletionProtectionEnabled bool    `json:"DeletionProtectionEnabled"`
	ItemCount                 int     `json:"ItemCount"`
	KeySchema                 []struct {
		AttributeName string `json:"AttributeName"`
		KeyType       string `json:"KeyType"`
	} `json:"KeySchema"`
	ProvisionedThroughput struct {
		NumberOfDecreasesToday int `json:"NumberOfDecreasesToday"`
		ReadCapacityUnits      int `json:"ReadCapacityUnits"`
		WriteCapacityUnits     int `json:"WriteCapacityUnits"`
	} `json:"ProvisionedThroughput"`
	TableArn                   string `json:"TableArn"`
	TableID                    string `json:"TableId"`
	TableName                  string `json:"TableName"`
	TableSizeBytes             int    `json:"TableSizeBytes"`
	TableStatus                string `json:"TableStatus"`
	TableThroughputModeSummary struct {
		LastUpdateToPayPerRequestDateTime float64 `json:"LastUpdateToPayPerRequestDateTime"`
		TableThroughputMode               string  `json:"TableThroughputMode"`
	} `json:"TableThroughputModeSummary"`
}

func (self *SRegion) ListTables() ([]string, error) {
	params := map[string]interface{}{
		"Limit": 100,
	}
	ret := []string{}
	for {
		part := struct {
			TableNames             []string
			LastEvaluatedTableName string
		}{}
		err := self.dynamodbRequest("ListTables", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.TableNames...)
		if len(part.TableNames) == 0 || len(part.LastEvaluatedTableName) == 0 {
			break
		}
		params["ExclusiveStartTableName"] = part.LastEvaluatedTableName
	}
	return ret, nil
}

func (self *SRegion) DescribeTable(name string) (*DynamoDB, error) {
	params := map[string]interface{}{
		"TableName": name,
	}
	ret := struct {
		Table DynamoDB
	}{}
	err := self.dynamodbRequest("DescribeTable", params, &ret)
	if err != nil {
		return nil, err
	}
	return &ret.Table, nil
}
