package main

import (
	"iPadel-GO/initializers"
	"iPadel-GO/server"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
	initializers.SyncDatabase()

}

func main() {

	server.Init()
}
