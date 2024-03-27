package instance

import config "github.com/we-and/weand_backend_common/config"

type ProjectConfig struct {
	DEV_FOLDER_KEY string
	ENV_KEY        string
	PROTOCOL_KEY   string
	PROJECT_NAME   string
}

type CommonInstance struct {
	ProjConfig    ProjectConfig
	ServerLongId  string
	ServerId      string
	ServerAddress string
	Version       string
	TLSFolder     string
	isHTTPS       bool
	AppConfig     config.AppConfig
	DbConnections *DbConnections
}
