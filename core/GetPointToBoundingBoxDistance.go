package core

func geoBoxDist(l LatLng, b BoundingBox) float64 {

	converted := l.ConvertToRadian()
	cAsPoint := converted.ConvertToPoint()

	if checkBetweenLongs(l, b) {
		//If between Lngs then:
		//Either straight down or straight up to bounding box.
		if l.Lat < b.Min.Y {
			return PointToPointDistanceCosine(cAsPoint, Point{X: cAsPoint.X, Y: DegToRad(b.Min.Y)})
		}
		if l.Lat > b.Max.Y {
			return PointToPointDistanceCosine(cAsPoint, Point{X: cAsPoint.X, Y: DegToRad(b.Max.Y)})
		}
		//Actually inside so return 0 distance
		return 0
	}

	//If on West
	if l.Lng < b.Min.X {
		//If North West, Use North West Corner
		if l.Lat > b.Max.Y {
			return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Min.X), Y: DegToRad(b.Max.Y)})
		}
		//If South West, use South West Corner
		if l.Lat < b.Min.Y {
			return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Min.X), Y: DegToRad(b.Min.Y)})
		}
		//Else Use line
		return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Min.X), Y: cAsPoint.Y})

	}

	//must be on east
	//If North East, Use North East Corner
	if l.Lat > b.Max.Y {
		return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Max.X), Y: DegToRad(b.Max.Y)})
	}
	//If South East, use South East Corner
	if l.Lat < b.Min.Y {
		return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Max.X), Y: DegToRad(b.Min.Y)})
	}
	//Else Use line
	return PointToPointDistanceCosine(cAsPoint, Point{X: DegToRad(b.Max.X), Y: cAsPoint.Y})

}

func checkBetweenLongs(l LatLng, b BoundingBox) bool {
	if l.Lng >= b.Min.X && l.Lng <= b.Max.X {
		return true
	}
	return false
}

func checkBetweenLats(l LatLng, b BoundingBox) bool {
	if l.Lat >= b.Min.Y && l.Lat <= b.Max.Y {
		return true
	}
	return false
}
