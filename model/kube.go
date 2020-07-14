package model

import (
	"simple-k8/go-common/api-base"
	"simple-k8/go-common/db-helper"
)

type kube struct {
	dbhelper.DbTable
}

var Kube = &kube{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_KUBE},
}

type KubeInfo struct {
	ID          int    `db:"id"`
	OldReplicas int    `db:"old_replicas"`
	NewReplicas int    `db:"new_replicas"`
	UpdateTime  string `db:"update_time"`
	Namespace   string `db:"namespace"`
	DeployName  int    `db:"deploy_name"`
}

func (l *kube) InsertKubeRecord(oldReplicas int32, newReplicas int32, namespace string, deployName string) string {
	if _, err := l.InsertWhere(dbhelper.UpdateFields{
		"old_replicas": oldReplicas,
		"new_replicas": newReplicas,
		"namespace":    namespace,
		"deploy_name":  deployName,
	}); err != nil {
		apibase.ThrowDBModelError(err)
	}
	return deployName
}
