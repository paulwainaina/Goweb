package processors

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Processor should implement ProcessorInterface
// Hint: Try using goroutines to process the data in parallel
// Requirement: Do not store all Trips in memory and cache your results
type Processor struct {
	data *TripsData
	wg   *sync.WaitGroup
	err error
}

var (
	driverRanking DriverRanking =*NewDrivingRancking()
	hotelRanking  HotelRanking  =*NewHotelRancking()
)

func NewProcessor(data *TripsData, wg *sync.WaitGroup) *Processor {
	return &Processor{data: data, wg: wg}
}
func (p *Processor) ProcessFunc() {
	p.err=nil
	defer p.wg.Done()
	x := 0
	for {
		x++
		select {
		case channel, ok := <-p.data.Trips:
			{
				fmt.Print(x)
				fmt.Println(channel)
				if !ok { //channel closed
					p.err = errors.New("channel closed")
					goto exit
				}
				av, isPresent := hotelRanking.hotelAverage[channel.HotelId]

				if isPresent {
					hotelRanking.hotelAverage[channel.HotelId] = (hotelRanking.hotelAverage[channel.HotelId] + av) / float64(2)
				} else {
					hotelRanking.hotelAverage[channel.HotelId] = channel.HotelRating
				}
				dav, disPresent := driverRanking.driverAverage[channel.DriverId]
				if disPresent {
					driverRanking.driverAverage[channel.DriverId] = (driverRanking.driverAverage[channel.DriverId] + dav) / float64(2)
				} else {
					driverRanking.driverAverage[channel.DriverId] = channel.DriverRating
				}
			}
		case <-time.After(100 * time.Millisecond):
			{ //Timeout for channel
				//p.err = errors.New("channel time out")
				goto exit
			}
		default:
			{
				//p.err = errors.New("default called")
				//goto exit
			}
		}
	}
exit:
}

func (p *Processor) StartProcessing() error {
	p.wg.Add(1)
	go p.ProcessFunc()
	p.wg.Wait()	
	return p.err
}
func (p *Processor) GetTopRankedDriver() *DriverRanking {
	return &driverRanking
}
func (p *Processor) GetTopRankedHotel() *HotelRanking {
	return &hotelRanking
}
