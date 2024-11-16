package router

// Setup initializes all routes
func SetupAll() {
	// Register all route groups
	HomeServer()
	setupV1Routes()
	setupV2Routes()
}

// For individual server starts
func SetupHome() {
	HomeServer()
}

func SetupV1() {
	V1Server()
}

func SetupV2() {
	V2Server()
}
