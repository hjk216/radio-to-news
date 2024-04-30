package main

import (
	"fmt"
	"strconv"
)



func SerializeDate(date string) bool {
	if len(date) > 10 || len(date) < 10 {
		fmt.Println("Error: String is not right size.")
		return false
	}

	_, year_err := strconv.Atoi(date[0:3])
	_, month_err := strconv.Atoi(date[5:6])
	_, day_err := strconv.Atoi(date[8:9])

	if year_err != nil {
        fmt.Println("Error:", year_err)
        return false
    }
	if month_err != nil {
        fmt.Println("Error:", month_err)
        return false
    }
	if day_err != nil {
        fmt.Println("Error:", day_err)
        return false
    }
	return true
}
