
# Geo Core 


## Library for common Geometry and location functions in Go.


### Point
 - Create points from XY
 - Create point from Lat Lng
 - Add Point
 - Subtract point

### Line
 - Create line from 2 points
 - Get Length
 - Get Centroid
 - Get Bounding Box

### Polyline

- Create Polyline from points
- Create Polyline from lines
- Get Length
- Get Centroid
- Get Vertices (Array of Points)
- Get Number of vertices
- Get Bounding Box
- Get Number of edges
- Check if closed chain of lines (i.e. a polygon

### Polygon

- Create Polygon from points
- Get Vertices (Array of Points)
- Get Number of Edges
- Get Perimeter (length)
- Get Area
- Get Bounding Box
- Check if closed chain. 


### Utilities

 - Get point to point distance (Cartesian)
 - Get point to point distance (Geographic)
 - Get point to line distance (Cartesian)
 - Get point to line distance (Geographic)
 - Get point to bounding box distance (Cartesian)
 - Get point to bounding box distance (Geographic)
 - Get Bearing
 - Get Cross Track Distance
