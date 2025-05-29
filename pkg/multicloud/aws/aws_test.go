package aws

import (
	"testing"
)

func TestAwsClientConfig(t *testing.T) {
	cfg := NewAwsClientConfig("", "", "", "")
	if cfg.getAssumeRoleName() != DefaultAssumeRoleName {
		t.Errorf("getAssumeRoleName != %s", DefaultAssumeRoleName)
	}
	cfg = cfg.SetAssumeRole("newRole")
	if cfg.getAssumeRoleName() != "newRole" {
		t.Errorf("getAssumeRoleName != newRole")
	}
}
