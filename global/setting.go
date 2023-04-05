package global

import (
	"database/sql"
	"github.com/dsxg666/notebook/pkg/setting"
	"go.uber.org/zap"
)

var (
	ServerSetting   *setting.ServerSettingS
	EmailSetting    *setting.EmailSettingS
	DatabaseSetting *setting.DatabaseSettingS
	SugarLogger     *zap.SugaredLogger
	MySqlConn       *sql.DB
)