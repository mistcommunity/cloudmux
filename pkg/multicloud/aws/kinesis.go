
package aws

type KinesisStream struct {
	StreamArn               string
	StreamCreationTimestamp int
	StreamModeDetails       struct {
		StreamMode string
	}
	StreamName   string
	StreamStatus string
	Shards       []struct {
		ShardId string
	}
}

func (self *SRegion) ListStreams() ([]KinesisStream, error) {
	params := map[string]interface{}{
		"MaxResults": "10000",
	}
	ret := []KinesisStream{}
	for {
		part := struct {
			StreamSummaries []KinesisStream
			NextToken       string
		}{}
		err := self.kinesisRequest("ListStreams", "/listStreams", params, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.StreamSummaries...)
		if len(part.StreamSummaries) == 0 || len(part.NextToken) == 0 {
			break
		}
		params["NextToken"] = part.NextToken
	}
	return ret, nil
}

func (self *SRegion) DescribeStream(name string) (*KinesisStream, error) {
	params := map[string]interface{}{
		"StreamName": name,
	}
	ret := struct {
		StreamDescription KinesisStream
	}{}
	err := self.kinesisRequest("DescribeStream", "/describeStream", params, &ret)
	if err != nil {
		return nil, err
	}
	return &ret.StreamDescription, nil
}
