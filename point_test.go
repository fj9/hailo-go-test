package main
import (
	"testing"
	"log"
)

func TestDistance(t *testing.T) {
	log.Printf("Test Distance")
	cases := []struct {
		p1lat, p1long, p2lat, p2long, expected float64
	}{
		{51.483136, -0.308410, 51.483296, -0.307380, 73},
	}
	for _, c := range cases {
		result := Distance(Point{c.p1lat, c.p1long}, Point{c.p2lat, c.p2long})
		log.Printf("Distance %f", result)
		if result < c.expected - 1 || result > c.expected + 1 {
			t.Errorf("Distance between Point 1(%f, %f) Point 2(%f, %f) == %f, " +
			" wanted between %f and %f", c.p1lat, c.p1long, c.p2lat, c.p2long, result, c.expected - 1, c.expected + 1)
		}
	}
}

func TestCloseBy(t *testing.T) {
	log.Printf("Test Close By")
	cases := []struct {
		p1lat, p1long, p2lat, p2long, withinRange float64
	}{
		{51.483136, -0.308410, 51.483296, -0.307380, 300},
	}
	for _, c := range cases {
		result := CloseBy(Point{c.p1lat, c.p1long}, Point{c.p2lat, c.p2long}, c.withinRange)
		if !result {
			t.Errorf("Expected Point 1(%f, %f) to be near Point 2(%f, %f)",
				c.p1lat, c.p1long, c.p2lat, c.p2long)
		}
	}
}

func TestNotCloseBy(t *testing.T) {
	log.Printf("Test Not Close By")
	cases := []struct {
		p1lat, p1long, p2lat, p2long, withinRange float64
	}{
		{51.483136, -0.308410, 51.483296, -5.307380, 300},
	}
	for _, c := range cases {
		result := CloseBy(Point{c.p1lat, c.p1long}, Point{c.p2lat, c.p2long}, c.withinRange)
		if result {
			t.Errorf("Expected Point 1(%f, %f) to not be near Point 2(%f, %f)",
				c.p1lat, c.p1long, c.p2lat, c.p2long)
		}
	}
}