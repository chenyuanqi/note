//Package category 模型
package it_main

import (
	"mlj/pkg/database"
	"time"
)

type OrderStatistics struct {
	Platform   string    `json:"platform,omitempty"`
	Date       string    `json:"date,omitempty"`
	Computed   float64   `json:"computed,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

func (statis *OrderStatistics) Create() {
	database.DBItMain.Create(&statis)
}

func (statis *OrderStatistics) Save() (rowsAffected int64) {
	result := database.DBItMain.Save(&statis)
	return result.RowsAffected
}

func (statis *OrderStatistics) Delete() (rowsAffected int64) {
	result := database.DBItMain.Delete(&statis)
	return result.RowsAffected
}

func (statis OrderStatistics) TableName() string {
	return "ho_order_statistics"
}
