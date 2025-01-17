package bootstrap

import (
	"github.com/goravel/framework/foundation"
	"github.com/orangbus/axios/config"
)

func Boot() {
	app := foundation.NewApplication()

	// Bootstrap the application
	app.Boot()

	// Bootstrap the config.
	config.Boot()
}
