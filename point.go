package main
import (
	"math"
//	"log"
)

var (
	earthRadius = 6371000.0
)


type Point struct {
	Lat float64
	Long float64
}

func CloseBy(p1, p2 Point, withinRange float64) bool {
	distance := Distance(p1, p2)
	if (distance < float64(withinRange)) {
		return true
	}else {
		return false
	}
}

func Distance(p1, p2 Point) float64 {
//	log.Printf("p1.lat %f, p1.long %f, p2.lat %f, p2.long%f", p1.Lat, p1.Long, p2.Lat, p2.Long)

	toRadians := func(x float64) float64 {
		return x * math.Pi / 180.0
	}

	dLat := toRadians(p2.Lat - p1.Lat)

//	log.Printf("Distance Lat %f", dLat)
	dLng := toRadians(p2.Long - p1.Long)
//	log.Printf("Distance Long %f", dLng)
	sindLat := math.Sin(dLat / 2)
	sindLng := math.Sin(dLng / 2)
	a := math.Pow(sindLat, 2) + math.Pow(sindLng, 2) * math.Cos(toRadians(p1.Lat)) * math.Cos(toRadians(p2.Lat))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
	dist := float64(earthRadius) * c

	return dist
}