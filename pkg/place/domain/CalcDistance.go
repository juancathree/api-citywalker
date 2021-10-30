package domain

import "math"

func (p *Place) CalcDistance(lat, long float64) float64 {
	PI := 3.1415926535897932384
	thisLat := p.Latitude() * (PI / 180)
	lat *= PI / 180
	thisLong := p.Longitude() * (PI / 180)
	long *= PI / 180
	return 6371.0 * math.Acos(math.Cos(thisLat)*math.Cos(lat)*math.Cos(long-thisLong)+math.Sin(thisLat)*math.Sin(lat))
}
