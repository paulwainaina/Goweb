package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golangchallenge/processors"
	"io"
	"os"
	"sync"
	"testing"
)

// Implement the following tests:
// Ensure top driver is found the processor
// Ensure top hotel is found the processor
// Ensure memory consumption is less than 128 mb
type JsonData struct {
	Drivers []processors.Driver
	Hotels  []processors.Hotel
}

var data = &processors.TripsData{}

func TestRetrieveData(t *testing.T) {
	jsonFile, err := os.Open("./data.json")
	if err != nil {
		t.Error(err.Error())
	}
	defer jsonFile.Close()
	fileByte, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Error(err.Error())
	}
	var jsonData JsonData
	err = json.Unmarshal(fileByte, &jsonData)
	if err != nil {
		t.Error(err.Error())
	}
	for _, driver := range jsonData.Drivers {
		data.Drivers = append(data.Drivers, &processors.Driver{
			Id:   driver.Id,
			Name: driver.Name,
		})
	}
	for _, hotel := range jsonData.Hotels {
		data.Hotels = append(data.Hotels, &processors.Hotel{
			Id:   hotel.Id,
			Name: hotel.Name,
		})
	}
	if data == nil {
		t.Error(errors.New("data is invalid"))
	}
	data.Trips = make(chan *processors.Trip,2)
	wg := &sync.WaitGroup{}

	//add trip
	wg.Add(1)
	go func() {
		defer wg.Done()
		data.Trips <- &processors.Trip{
			DriverId:     "57d9072e-8708-4d87-b5fd-4ea573e2e20b",
			HotelId:      "ecf9ba5f-4770-44b4-a40d-fbdc0a3e987c",
			DriverRating: 2.0,
			HotelRating:  3.0,
		}
		data.Trips <- &processors.Trip{
			DriverId:     "57d9072e-8708-4d87-b5fd-4ea573e2e20b",
			HotelId:      "03f38100-f900-4f63-bb58-3e52718a87ee",
			DriverRating: 3.0,
			HotelRating:  2.0,
		}
		data.Trips <- &processors.Trip{
			DriverId:     "79e8784e-afe9-4860-af9a-bcc74f007b96",
			HotelId:      "7399397f-919d-4a54-a3b1-02e07e754165",
			DriverRating: 8.0,
			HotelRating:  0.0,
		}
		data.Trips <- &processors.Trip{
			DriverId:     "79e8784e-afe9-4860-af9a-bcc74f007b96",
			HotelId:      "f50e651d-d1a5-44c7-958c-30b90c9461ce",
			DriverRating: 10.0,
			HotelRating:  1.5,
		}
		data.Trips <- &processors.Trip{
			DriverId:     "79e8784e-afe9-4860-af9a-bcc74f007b96",
			HotelId:      "99bac8df-0d7a-4f26-abdc-92b8122e92ff",
			DriverRating: 9.0,
			HotelRating:  3.5,
		}
		//close(data.Trips)		
	}()
	//test top
	processor := processors.CreateProcessorFromData(data, wg)
	err = processor.StartProcessing()
	if err != nil {
		t.Error(err)
	}
	wg.Wait()
	topDriver := processor.GetTopRankedDriver()
	fmt.Printf("Top driver found: %s\n", topDriver)
	topHotel := processor.GetTopRankedHotel()
	fmt.Printf("Top hotel found: %s\n", topHotel)
}
