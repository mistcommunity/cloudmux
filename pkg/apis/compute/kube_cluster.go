package compute

const (
	KUBE_CLUSTER_STATUS_RUNNING  = "running"
	KUBE_CLUSTER_STATUS_CREATING = "creating"
	KUBE_CLUSTER_STATUS_DELETING = "deleting"
	KUBE_CLUSTER_STATUS_ABNORMAL = "abnormal"
	// 升级中
	KUBE_CLUSTER_STATUS_UPDATING = "updating"
	// 升级失败
	KUBE_CLUSTER_STATUS_UPDATING_FAILED = "updating_failed"
	// 伸缩中
	KUBE_CLUSTER_STATUS_SCALING = "scaling"
	// 停止
	KUBE_CLUSTER_STATUS_STOPPED = "stopped"
)
