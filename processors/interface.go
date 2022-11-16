package processors

import (
	"sort"
	"strconv"
	"sync"
)

type DriverRanking struct {
	driverAverage map[string]float64
}

func NewDrivingRancking() *DriverRanking {
	var dr DriverRanking
	dr.driverAverage = make(map[string]float64)
	return &dr
}
func (*DriverRanking) String() string {
	// Implement this function
	keys := make([]string, 0, len(driverRanking.driverAverage))
	for k := range driverRanking.driverAverage {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return driverRanking.driverAverage[keys[i]] < driverRanking.driverAverage[keys[j]]
	})
	var id string
	var average float64
	for key, value := range hotelRanking.hotelAverage {
		id = key
		average = value
		break
	}

	return id + " " + strconv.FormatFloat(average, 'E', -1, 64)
}

type HotelRanking struct {
	// Fill in your properties here
	hotelAverage map[string]float64
}

func NewHotelRancking() *HotelRanking {
	var hr HotelRanking
	hr.hotelAverage = make(map[string]float64)
	return &hr
}
func (*HotelRanking) String() string {
	// Implement this function
	keys := make([]string, 0, len(hotelRanking.hotelAverage))
	for k := range hotelRanking.hotelAverage {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return hotelRanking.hotelAverage[keys[i]] < hotelRanking.hotelAverage[keys[j]]
	})
	var id string
	var average float64
	for key, value := range hotelRanking.hotelAverage {
		id = key
		average = value
		break
	}

	return id + " " + strconv.FormatFloat(average, 'E', -1, 64)
}

type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) ProcessorInterface {
	// @todo Initialize your processor here
	var processor ProcessorInterface = NewProcessor(data, wg)
	return processor
}
