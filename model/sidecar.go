package model

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"simple-k8/go-common/api-base"
	"simple-k8/go-common/db-helper"
	"simple-k8/go-common/utils"
)

type sidecarList struct {
	dbhelper.DbTable
}

var SidecarList = &sidecarList{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_SIDECAR_LIST},
}

const (
	SIDECAR_UNKNOWN_STATUS = -1
	SIDECAR_NOT_INSTALLED  = iota
	SIDECAR_INSTALLING
	SIDECAR_INSTALL_FAILURE
	SIDECAR_INSTALLED
	SIDECAR_UPDATING
	SIDECAR_UPDATE_FAILURE
)

type SimpleSidecarInfo struct {
	ID         string            `db:"id"`
	Name       string            `db:"name"`
	Version    string            `db:"version"`
	Disabled   bool              `db:"disabled"`
	Status     int               `db:"status"`
	OsType     string            `db:"os_type"`
	DeployDate dbhelper.NullTime `db:"deploy_date"`
	AutoDeploy bool              `db:"auto_deployment"`
	UpdateDate dbhelper.NullTime `db:"last_update_date"`
	AutoUpdate bool              `db:"auto_updated"`
}

type SidecarInfo struct {
	SimpleSidecarInfo
	EcsFlag    bool           `db:"is_ecs"`
	OsPlatform string         `db:"os_platform"`
	OsVersion  string         `db:"os_version"`
	Host       string         `db:"host"`
	LocalIp    string         `db:"local_ip"`
	CpuSerial  string         `db:"cpu_serial"`
	CpuCores   uint           `db:"cpu_cores"`
	MemSize    uint64         `db:"mem_size"`
	SwapSize   uint64         `db:"swap_size"`
	ServerHost string         `db:"server_host"`
	ServerPort int            `db:"server_port"`
	SshHost    string         `db:"ssh_host"`
	SshPort    int            `db:"ssh_port"`
	SshUser    string         `db:"ssh_user"`
	SshPwd     string         `db:"ssh_password"`
	CpuUsage   float32        `db:"cpu_usage"`
	MemUsage   int64          `db:"mem_usage"`
	SwapUsage  int64          `db:"swap_usage"`
	Load1      float32        `db:"load1"`
	Uptime     float64        `db:"uptime"`
	DiskUsage  sql.NullString `db:"disk_usage"`
	NetUsage   sql.NullString `db:"net_usage"`
}

type DeployCallbackInfo struct {
	ID          int64  `db:"id"`
	Time        int64  `db:"time"`
	ClientID    string `db:"client_id"`
	InstallType string `db:"install_type"`
	InstallRes  string `db:"install_res"`
	MSG         []byte `db:"msg"`
	RequestUrl  string `db:"request_url"`
	IP          string `db:"ip"`
}

func (l *sidecarList) CreateSidecarByDeploy(id uuid.UUID) (uuid.UUID, error) {
	if _, err := l.InsertWhere(dbhelper.UpdateFields{
		"id":       id,
		"status":   SIDECAR_INSTALLED,
		"disabled": false,
	}); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (l *sidecarList) NewSidecarRecord(name string, version string) (uuid.UUID, error) {
	id := uuid.NewV4()
	if _, err := l.InsertWhere(dbhelper.UpdateFields{
		"id":       id.String(),
		"name":     name,
		"status":   SIDECAR_NOT_INSTALLED,
		"disabled": false,
	}); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (l *sidecarList) GetSidecarInfo(id uuid.UUID) (*SidecarInfo, error) {
	whereCause := dbhelper.WhereCause{}
	info := SidecarInfo{}
	err := l.GetWhere(nil, whereCause.Equal("id", id), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

var _getSidecarListFields = utils.GetTagValues(SimpleSidecarInfo{}, "db")

func (l *sidecarList) GetSidecarList(pagination *apibase.Pagination) ([]SimpleSidecarInfo, int) {
	rows, totalRecords, err := l.SelectWhere(_getSidecarListFields, nil, pagination)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}

	list := []SimpleSidecarInfo{}
	for rows.Next() {
		info := SimpleSidecarInfo{}
		err = rows.StructScan(&info)
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		list = append(list, info)
	}
	return list, totalRecords
}

func (l *sidecarList) UpdateSidecar(id uuid.UUID, updateFields dbhelper.UpdateFields) error {
	return l.UpdateWhere(dbhelper.MakeWhereCause().Equal("id", id), updateFields, false)
}
