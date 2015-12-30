package main
// consume ten points
// traverse the ten points in order
// check for near by station
// if near a station establish traffic conditions
// calculate speed
// post traffic report
// shutdown
import (
	"math/rand"
	"fmt"
//	"log"
)

type Robot struct {
	ID              int
	CurrentLocation RobotLocation
	TubeLocations   []TubeLocation
}

func (r *Robot) Consume(locations []RobotLocation, trafficConditions chan string) {
//	log.Printf("%d", len(locations))
	if len(locations) < 11 {
		for _, location := range locations {
			pastLocation := r.CurrentLocation
			r.CurrentLocation = location
			if (r.nearTubeStation()) {
				speed := CalculateSpeed(pastLocation, location)
				trafficConditions <- r.createTrafficReport(speed)
			}
		}
	} else {
		panic("too many locations");
	}
}

func CalculateSpeed(location1, location2 RobotLocation) float64 {
	distance := Distance(location1.Point, location2.Point)
//		log.Printf("Distance %f", distance)

	durationSec := location2.Time.Sub(location1.Time).Seconds()

	//	log.Printf("Duration %f", durationSec)
	return distance / durationSec

}


func (r *Robot) nearTubeStation() bool {
	for _, tube := range r.TubeLocations {
		if (CloseBy(r.CurrentLocation.Point, tube.Point, 350)) {
			return true
		}
	}
	return false;
}

func randomTrafficCondition() string {
	switch rand.Intn(3) {
	case 0:
		return "HEAVY"
	case 1:
		return "MODERATE"
	case 2:
		return "LIGHT"
	}
	return ""
}

func (r *Robot) createTrafficReport(speed float64) string {
	//	log.Printf("Robot ID = %d, Time = %s, Speed = %f, Traffic conditions = %s", r.ID, r.CurrentLocation.Time.String(), speed, randomTrafficCondition())
	return fmt.Sprintf("Robot ID = %d, Time = %s, Speed = %f, Traffic conditions = %s", r.ID, r.CurrentLocation.Time.Format("2006-01-02 15:04:05"), speed, randomTrafficCondition())

}