package journey

import (
	"fmt"
	"time"
)

// StationOrderItem untuk single station in sequence
type StationOrderItem struct {
	ID    int
	Slug  string
	Name  string
	Order int
}

// StationOrder constant untuk MRT North-South Line ordering
var StationOrder = []StationOrderItem{
	{ID: 4, Slug: "stasiun-lebak-bulus", Name: "Stasiun MRT Lebak Bulus Bank Syariah Indonesia", Order: 1},
	{ID: 33, Slug: "stasiun-fatmawati-indomaret", Name: "Stasiun MRT Fatmawati Indomaret", Order: 2},
	{ID: 32, Slug: "stasiun-cipete-raya", Name: "Stasiun MRT Cipete Raya TUKU", Order: 3},
	{ID: 34, Slug: "stasiun-haji-nawi", Name: "Stasiun MRT Haji Nawi", Order: 4},
	{ID: 35, Slug: "stasiun-blok-a", Name: "Stasiun MRT Blok A", Order: 5},
	{ID: 36, Slug: "stasiun-blok-m-bca", Name: "Stasiun MRT Blok M BCA", Order: 6},
	{ID: 37, Slug: "stasiun-senayan-mastercard", Name: "Stasiun MRT Senayan Mastercard", Order: 7},
	{ID: 38, Slug: "stasiun-istora-mandiri", Name: "Stasiun MRT Istora Mandiri", Order: 8},
	{ID: 39, Slug: "stasiun-bendungan-hilir", Name: "Stasiun MRT Bendungan Hilir", Order: 9},
	{ID: 40, Slug: "stasiun-setiabudi-astra", Name: "Stasiun MRT Setiabudi Astra", Order: 10},
	{ID: 41, Slug: "stasiun-dukuh-atas-bni", Name: "Stasiun MRT Dukuh Atas BNI", Order: 11},
	{ID: 5, Slug: "stasiun-asean", Name: "ASEAN Headquarter", Order: 12},
	{ID: 6, Slug: "bundaran-hi-bank-jakarta", Name: "Bundaran HI Bank Jakarta", Order: 13},
}

// FindStationBySlug return station info dari slug
func FindStationBySlug(slug string) (*StationOrderItem, error) {
	for _, station := range StationOrder {
		if station.Slug == slug || slugContains(station.Slug, slug) {
			return &station, nil
		}
	}
	return nil, fmt.Errorf("station not found: %s", slug)
}

// Helper function untuk partial slug match
func slugContains(stationSlug, reqSlug string) bool {
	// Remove "stasiun-" prefix untuk matching
	stationClean := stationSlug
	if len(stationSlug) > 10 && stationSlug[:10] == "stasiun-" {
		stationClean = stationSlug[10:]
	}
	reqClean := reqSlug
	if len(reqSlug) > 10 && reqSlug[:10] == "stasiun-" {
		reqClean = reqSlug[10:]
	}
	return contains(stationClean, reqClean) || contains(stationSlug, reqSlug)
}

func contains(haystack, needle string) bool {
	return len(haystack) >= len(needle) && (haystack == needle || len(haystack) > len(needle) && findSubstring(haystack, needle))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// CalculateFare berdasarkan jumlah stations antara from dan to
func CalculateFare(stationsCount int) int {
	// Fare structure berdasarkan data:
	// 5 stations (Fatmawati→Senayan) = 9,000 IDR
	// 12 stations (Lebak Bulus→Bundaran HI) = 14,000 IDR
	
	if stationsCount <= 0 {
		return 0
	}
	if stationsCount <= 5 {
		return 9000
	}
	if stationsCount <= 12 {
		return 14000
	}
	return 14000 // Max fare
}

// CalculateDuration estimation: 2 minutes per station
func CalculateDuration(stationsCount int) int {
	if stationsCount <= 0 {
		return 0
	}
	duration := stationsCount * 2
	if duration > 30 {
		return 30 // Max 30 menit untuk full line
	}
	return duration
}

// GetDirection return direction berdasarkan position
func GetDirection(fromOrder, toOrder int) string {
	if fromOrder < toOrder {
		return "northbound"
	}
	return "southbound"
}

// GetNextTrain waktu kereta selanjutnya berdasarkan current time
func GetNextTrain(scheduleTimes []string, currentTime time.Time) (string, error) {
	currentMinutes := currentTime.Hour()*60 + currentTime.Minute()
	
	for _, timeStr := range scheduleTimes {
		// Parse "HH:MM" format
		var hour, minute int
		_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
		if err != nil {
			continue
		}
		
		minute += currentTime.Second() / 60
		
		scheduleMinutes := hour*60 + minute
		if scheduleMinutes > currentMinutes {
			return timeStr, nil
		}
	}
	
	return "", fmt.Errorf("no more trains today")
}

// GetNextTrainFromNow dengan parameter timestamp untuk filter schedule
func GetNextTrainFromNow(scheduleTimes []string, currentTime time.Time, windowMinutes int) ([]string, error) {
	currentMinutes := currentTime.Hour()*60 + currentTime.Minute()
	
	var futureTimes []string
	for _, timeStr := range scheduleTimes {
		var hour, minute int
		_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
		if err != nil {
			continue
		}
		
		minute += currentTime.Second() / 60
		scheduleMinutes := hour*60 + minute
		
		if scheduleMinutes > currentMinutes && scheduleMinutes <= currentMinutes+windowMinutes {
			futureTimes = append(futureTimes, timeStr)
		}
	}
	
	if len(futureTimes) == 0 {
		return nil, fmt.Errorf("no trains in next %d minutes", windowMinutes)
	}
	
	return futureTimes, nil
}