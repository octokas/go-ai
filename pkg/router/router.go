package router

// Setup initializes all routes
func Setup() {
	// Register all route groups
	setupHomeRoutes()
	setupV1Routes()
	setupV2Routes()
}
