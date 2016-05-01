package weather

import (
	"log"
)

// TemperatureFeel is how the temperature feels
type TemperatureFeel string

// MoistureFeel is a subjective gauge of humidity
type MoistureFeel string

// WindFeel is how windy it is
type WindFeel string

// for TemperatureFeel
const (
	UNKNOWNTEMPERATURE TemperatureFeel = "UNKNOWNTEMPERATURE"
	COLD                               = "COLD"
	NORMAL                             = "NORMAL"
	WARM                               = "WARM"
	HOT                                = "HOT"
)

// for MoistureFeel
const (
	UNKNOWNHUMIDITY MoistureFeel = "UNKNOWNHUMIDITY"
	DRY                          = "DRY"
	MOIST                        = "MOIST"
	DRIZZLE                      = "DRIZZLE"
	WET                          = "WET"
)

// for WindFeel
const (
	UNKNOWNWINDY WindFeel = "UNKNOWNWINDY"
	CALM                  = "CALM"
	WINDY                 = "WINDY"
	STRONG                = "STRONG"
)

// Feel to communicate out
type Feel struct {
	Wind        WindFeel        `json:"wind"`
	Temperature TemperatureFeel `json:"temperature"`
	Moisture    MoistureFeel    `json:"moisture"`
	Longitude   float64         `json:"longitude"`
	Latitude    float64         `json:"latitude"`
}

// GetConditions finds current weather conditions
func GetConditions(ip uint32) Feel {
	log.Printf("passed in ip: %d", ip)
	lon, lat := GetLongLat(100)
	return Feel{WINDY, COLD, WET, lon, lat}
}
