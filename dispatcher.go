package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"time"
)

func main() {
	fmt.Printf(("hello, world\n"))
	//read robot files
	//instantiate robots
	//pass files to robots
	//receive messages from each robot
	// Open an input file, exit on error.
	//	var tubeStations []TubeLocation = readTubeFile();

	ids := []int{5937, 6043}
	for _, id := range ids {
		trafficReport := make (chan string)
		go dispatch(id, trafficReport)
		for report := range trafficReport {
			fmt.Println(report)
		}
	}
}


type RobotLocation struct {
	RobotId string
	Point   Point
	Time    time.Time
}

type TubeLocation struct {
	Description string
	Point       Point
}

func dispatch(id int, trafficReport chan string) {
	var roboLocations []RobotLocation = readRobotFile(id);
	var robot Robot = Robot{id, roboLocations[0], readTubeFile()}

	locationsToProcess := make([]RobotLocation, 0, 10)
	counter := 0
	for _, location := range roboLocations {
		//		log.Printf("long : %s\n", location.Point.Lat)
		hour, min, _ := location.Time.Clock()
		if (hour < 8 || hour == 8 && min < 10 ) {
			if (counter < 10) {
				locationsToProcess = append(locationsToProcess, location)
				counter++
			}else {
				robot.Consume(locationsToProcess, trafficReport)
				locationsToProcess = nil
				locationsToProcess = append(locationsToProcess, location)
				counter = 1
			}
		} else {
			break
		}
	}
	//push any remaining locations
	robot.Consume(locationsToProcess, trafficReport)
	close(trafficReport)
}

func readTubeFile() []TubeLocation {
	locations := make([]TubeLocation, 0)
	for _, each := range readCsv("tube.csv",3) {
		var location TubeLocation
		location.Description = each[0]
		lat, _ := strconv.ParseFloat(each[1], 64)
		long, _ := strconv.ParseFloat(each[2], 64)
		location.Point = Point{lat, long}
		locations = append(locations, location)
	}
	return locations
}


func readRobotFile(id int) []RobotLocation {

	locations := make([]RobotLocation, 0)
	for _, each := range readCsv(fmt.Sprintf("%d.csv", id),4) {
		var location RobotLocation
		location.RobotId = each[0]
		lat, _ := strconv.ParseFloat(each[1], 64)
		long, _ := strconv.ParseFloat(each[2], 64)
		location.Point = Point{lat, long}
		location.Time, _ = time.Parse("2006-01-02 15:04:05", each[3])
		locations = append(locations, location)
	}
	return locations
}

func readCsv(file string, noFields int) [][]string {
	pwd, _ := os.Getwd()
	csvfile, err := os.Open(pwd + "/lib/godoc/" + file)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer csvfile.Close()

	reader := csv.NewReader(csvfile);

	reader.FieldsPerRecord = noFields

	csvdata, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return csvdata

}