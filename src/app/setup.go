// Package app performs basic app functions like setup, loading config and global definitions
package app

// Setup initializes application components.
func Setup() {
	setupLoggerOrPanic()
}

func setupLoggerOrPanic() {
	if err := initLogger(); err != nil {
		panic("failed to setup logger: " + err.Error())
	}
}

func initLogger() error {
	if Verbose || Build == BuildDev {
		return setupDevLogger()
	}
	return setupProdLogger()
}
