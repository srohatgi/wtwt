package weather

import (
	"log"
)

// GetLongLat returns longitude and latitude of the passed in ip address
func GetLongLat(ip uint32) (float64, float64) {
	log.Printf("finding location for ip=%d", ip)
	return 0.0, 0.0
}
