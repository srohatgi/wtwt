package weather

import (
	"fmt"
)

// GetLongLat returns longitude and latitude of the passed in ip address
func GetLongLat(ip uint32) (float64, float64) {
	fmt.Printf("ip=%d", ip)
	return 0.0, 0.0
}
