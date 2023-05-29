package initialize

import (
	"ecommerce/internal/config/database"
)

func mongodb() {
	database.ConnectMongoDBEcommerce()
}
