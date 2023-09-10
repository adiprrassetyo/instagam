package times

import (
	"log"
	"time"
)

func Now(time_loc string) time.Time {
	location, err := time.LoadLocation(time_loc)
	if err != nil {
		log.Println("Error LoadLocation: ", err)
		return time.Now()
	}
	return time.Now().In(location)
}
