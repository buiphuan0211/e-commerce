package initialize

import (
	"ecommerce/internal/config"
	"ecommerce/pkg/admin/errorcode"
)

// Init ...
func Init() {
	// Config ...
	config.Init()

	// ErrorCode ...
	errorcode.Init()

	// zookeeper ...
	zookeeper()

	// Mongodb ...
	mongodb()
}
