package migrations

import (
	"api/app/models"
	"api/pkg/migrate"

	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Links struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Links{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Links{})
	}

	migrate.Add("2022_06_20_105547_add_links_table", up, down)
}
