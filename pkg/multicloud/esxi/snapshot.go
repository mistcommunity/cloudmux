
package esxi

import (
	"fmt"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/multicloud"
)

type SVirtualMachineSnapshot struct {
	multicloud.SResourceBase
	multicloud.STagBase
	snapshotTree types.VirtualMachineSnapshotTree
	vm           *SVirtualMachine
}

func NewSnapshot(st types.VirtualMachineSnapshotTree) *SVirtualMachineSnapshot {
	return &SVirtualMachineSnapshot{
		snapshotTree: st,
	}
}

func (s *SVirtualMachineSnapshot) GetId() string {
	return moRefId(s.snapshotTree.Snapshot)
}

func (s *SVirtualMachineSnapshot) GetName() string {
	return s.snapshotTree.Name
}

func (s *SVirtualMachineSnapshot) GetDescription() string {
	return s.snapshotTree.Description
}

func (s *SVirtualMachineSnapshot) GetGlobalId() string {
	return fmt.Sprintf("%s-%d", s.vm.GetGlobalId(), s.snapshotTree.Id)
}

func (s *SVirtualMachineSnapshot) GetStatus() string {
	return api.INSTANCE_SNAPSHOT_READY
}

func (s *SVirtualMachineSnapshot) Refresh() error {
	return nil
}

func (s *SVirtualMachineSnapshot) IsEmulated() bool {
	return false
}

func (s *SVirtualMachineSnapshot) GetProjectId() string {
	return s.vm.GetProjectId()
}

func (s *SVirtualMachineSnapshot) Delete() error {
	cTrue := true
	req := types.RemoveSnapshot_Task{
		This:           s.snapshotTree.Snapshot.Reference(),
		RemoveChildren: false,
		Consolidate:    &cTrue,
	}
	res, err := methods.RemoveSnapshot_Task(s.vm.manager.context, s.vm.manager.client.Client, &req)
	if err != nil {
		return errors.Wrap(err, "RemoveSnapshot_Task")
	}
	return object.NewTask(s.vm.manager.client.Client, res.Returnval).Wait(s.vm.manager.context)
}
