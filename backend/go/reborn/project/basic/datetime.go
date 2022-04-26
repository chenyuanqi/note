package main

import (
	"fmt"
	"time"
)

func dump(params ...interface{}) {
	fmt.Println(params...)
}

func main() {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := "2022-12-31"
	// func Parse(layout, value string) (Time, error)
	// func (t Time) Format(layout string) string
	t, _ := time.Parse(layoutISO, date)
	dump(t)                  // 2022-12-31 00:00:00 +0000 UTC
	dump(t.Format(layoutUS)) // December 31, 2022
}
