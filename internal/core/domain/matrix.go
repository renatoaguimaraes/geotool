package domain

import (
	"math"
)

const earthRaidusKm = 6371 // radius of the earth in kilometers.

// Segment from a coordinate to other coordinate with Haversine distance
type Segment struct {
	From     *Coordinate `json:"from"`
	To       *Coordinate `json:"to"`
	Distance float64     `json:"distance"`
}

// Calculate the distance between two locations
func (segment *Segment) Calculate() {
	segment.Distance = distance(segment.From, segment.To)
}

// Coordinate represents a geographic coordinate.
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lgt float64 `json:"lgt"`
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth.
func distance(p, q *Coordinate) (km float64) {
	lat1 := degreesToRadians(p.Lat)
	lgt1 := degreesToRadians(p.Lgt)
	lat2 := degreesToRadians(q.Lat)
	lgt2 := degreesToRadians(q.Lgt)

	diffLat := lat2 - lat1
	difflgt := lgt2 - lgt1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(difflgt/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRaidusKm
}

// NewSegment create a Segment with the calculated distance between them.
func NewSegment(fromLat, fromLgt, toLat, toLgt float64) *Segment {
	segment := &Segment{
		From: &Coordinate{Lat: fromLat, Lgt: fromLgt},
		To:   &Coordinate{Lat: toLat, Lgt: toLgt},
	}
	segment.Calculate()
	return segment
}
