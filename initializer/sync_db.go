package initializers

import (
	"github.com/TAhirr01/firstmodule/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
