package main

import (
	"iPadel-GO/initializers"
	"iPadel-GO/server"
)

func init() {
	initializers.LoadEnvVariables()
	//	initializers.ConnectTODBPostgres()
	initializers.ConnectTODBMSSQL()
	initializers.SyncDatabase()

}

func main() {

	server.Init()
}
