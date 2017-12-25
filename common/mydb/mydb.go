package mydb

import (
	"database/sql"

	mylog "../common_log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MapSrvCfg = make(map[uint32](*ServerCfg))
)

type ServerCfg struct {
	Svrid       uint32
	Name        string
	Group       uint32
	SvrType     uint32
	GameType    uint32
	GameSubType uint8
	Svrip       string
	Svrlanip    string
	Svrport     uint32
	Svrlanport  uint32
	Websvrport  uint32
	Phpport     uint32
	OpenRobot   uint8
	State       int
	Onlines     int
	Robots      int
	Report_time string
}

func (this *ServerCfg) GetSrvid() uint32 {
	return this.Svrid
}

func (this *ServerCfg) GetSvrip() string {
	return this.Svrip
}

func (this *ServerCfg) GetSvrPort() uint32 {
	return this.Svrport
}
func newSrvCfg() (srv *ServerCfg) {
	return &ServerCfg{
		Svrid:       0,
		Name:        "",
		Group:       0,
		SvrType:     0,
		GameType:    0,
		GameSubType: 0,
		Svrip:       "",
		Svrport:     0,
		Svrlanip:    "",
		Svrlanport:  0,
		Websvrport:  0,
		Phpport:     0,
		OpenRobot:   0,
		State:       0,
		Onlines:     0,
		Robots:      0,
		Report_time: "",
	}
}

func (dstsrv *ServerCfg) Assignment(src *ServerCfg) {
	dstsrv.Svrid = src.Svrid
	dstsrv.Name = src.Name
	dstsrv.Group = src.Group
	dstsrv.SvrType = src.SvrType
	dstsrv.GameType = src.GameType
	dstsrv.GameSubType = src.GameSubType
	dstsrv.Svrip = src.Svrip
	dstsrv.Svrport = src.Svrport
	dstsrv.Svrlanip = src.Svrlanip
	dstsrv.Svrlanport = src.Svrlanport
	dstsrv.Websvrport = src.Websvrport
	dstsrv.Phpport = src.Phpport
	dstsrv.OpenRobot = src.OpenRobot
}

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.2.241:3306)/chess?charset=utf8")
	if nil != err {
		mylog.LOG_ERROR("fail to open database", err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from serverinfo")
	if err != nil {
		mylog.LOG_ERROR("Query failed", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		cfg := newSrvCfg()
		rows.Columns()
		err = rows.Scan(&cfg.Svrid, &cfg.Name, &cfg.Group, &cfg.SvrType, &cfg.GameType, &cfg.GameSubType, &cfg.Svrip, &cfg.Svrlanip, &cfg.Svrport, &cfg.Svrlanport, &cfg.Websvrport, &cfg.Phpport, &cfg.OpenRobot, &cfg.State, &cfg.Onlines, &cfg.Robots, &cfg.Report_time)
		if err != nil {
			mylog.LOG_ERROR("read error", err.Error())
		}
		MapSrvCfg[cfg.Svrid] = cfg
	}
}
