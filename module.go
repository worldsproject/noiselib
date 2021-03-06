package noiselib

type Module interface {

	//Returns a source module connected to thise noise module.
	GetSourceModule(index int) Module

	//Generates an output value given the coordinates of the specified input
	// value. Before an applicaion can call this method, it must first connect
	// all required source modules via SetSourceModule(). If these source modules
	// are not connected to this module, this method raises an error.
	GetValue(x, y, z float64) float64

	//Sets a source module in the given index.
	SetSourceModule(index int, sourceModule Module)
}

func ClampValue(value, lowerBound, upperBound int) int {
	if value < lowerBound {
		return lowerBound
	} else if value > upperBound {
		return upperBound
	} else {
		return value
	}
}

func ClampValueFloat(value, lowerBound, upperBound float64) float64 {
	if value < lowerBound {
		return lowerBound
	} else if value > upperBound {
		return upperBound
	} else {
		return value
	}
}
