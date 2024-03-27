package lifecycle

import (
	config "github.com/we-and/weand_backend_common/config"
	commoninstance "github.com/we-and/weand_backend_common/instance"

	"gorm.io/gorm"
)

func RestartConnections(c *commoninstance.CommonInstance, appConfig *config.AppConfig) bool {
	if c == nil {
		return false
	}
	if appConfig == nil {
		return false
	}
	map_ := *(c.DbConnections)
	c.CloseConnections(map_.ConnectionsMap)

	// /CREATE database connections
	connections := map[string]*gorm.DB{}
	c.GenerateLiveDB(*appConfig, &connections)
	c.GenerateTestDB(*appConfig, &connections)
	if c.IsLocalDevEnvironment() {
		c.GenerateLocalDB(*appConfig, &connections)
	}

	dbConnections := commoninstance.DbConnections{
		ConnectionsMap: connections,
	}
	c.DbConnections = &dbConnections
	return true
}
