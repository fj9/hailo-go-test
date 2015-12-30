package main
import (
	"testing"
	"time"
	"log"
	"fmt"
)
func TestCalculateSpeed(t *testing.T) {
	log.Printf("Test Speed")
	var location RobotLocation
	location.RobotId = "1"
	location.Point = Point{51.483136, -0.308410}
	location.Time = time.Date(2011, time.March, 22, 7, 55, 52, 0, time.UTC)
	var location2 RobotLocation
	location2.RobotId = "1"
	location2.Point = Point{51.483296, -0.307380}
	location2.Time = time.Date(2011, time.March, 22, 7, 59, 52, 0, time.UTC)

	result := CalculateSpeed(location, location2)
	expected := 0.304
	if (result > expected + 0.1 || result < expected - 0.1) {
		t.Errorf("Expected speed between point 1 and 2 should be between %f and %f instead it was %f",
			expected - 0.1, expected + 0.1, result)
	}
}

func TestConsume(t *testing.T) {
	log.Printf("Test Consume")
	var location RobotLocation
	location.RobotId = "1234"
	location.Point = Point{51.483136, -0.308410}
	location.Time = time.Date(2011, time.March, 22, 7, 55, 52, 0, time.UTC)
	var robot Robot = Robot{1234, location}
	//	1234,"51.476105","-0.100224","2011-03-22 07:55:26"
	//	1234,"51.475967","-0.100368","2011-03-22 07:55:40"
	trafficConditions := make (chan string, 10)
	robot.Consume(readRobotFile(1234), trafficConditions)
	close(trafficConditions)
	go Print(trafficConditions)
	//	expected := "Robot ID = 1234, Time = 2011-03-22 07:55:26, Speed = 1, Traffic Condition: Heavy" +
	//	"Robot ID = 1234, Time = 2011-03-22 07:55:40, Speed = 1, Traffic Condition: Heavy"
}

func Print(report chan string) {
	for {
		fmt.Println(<-report)
	}
}