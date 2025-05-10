package router

func initGlobalRoutes(config *Config) {
	globalApi := config.Server.Group("/api/v1")

	globalApiFile := globalApi.Group("/")
	globalApiFile.GET("")
}
