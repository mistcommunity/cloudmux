package compute

type TNicType string

const (
	HOST_TYPE_ALIYUN         = "aliyun"
	HOST_TYPE_AWS            = "aws"
	HOST_TYPE_AZURE          = "azure"
	HOST_TYPE_ESXI 			 = "esxi" // # VMWare vSphere ESXi
	HOST_TYPE_GOOGLE         = "google"
	HOST_TYPE_NUTANIX        = "nutanix"
	HOST_TYPE_PROXMOX        = "proxmox"

	// # possible status
	HOST_ONLINE  = "online"
	HOST_OFFLINE = "offline"

	NIC_TYPE_IPMI   = TNicType("ipmi")
	NIC_TYPE_ADMIN  = TNicType("admin")
	NIC_TYPE_NORMAL = TNicType("")

	BAREMETAL_READY   = "ready"
	BAREMETAL_RUNNING = "running"
	BAREMETAL_UNKNOWN = "unknown"

	HOST_STATUS_RUNNING = BAREMETAL_RUNNING
	HOST_STATUS_READY   = BAREMETAL_READY
	HOST_STATUS_UNKNOWN = BAREMETAL_UNKNOWN
)
