package noiselib

type Displace struct {
	SourceModule []Module
}

func (d Displace) GetSourceModule(index int) Module {
	return d.SourceModule[index]
}

func (d Displace) SetSourceModule(index int, sourceModule Module) {
	d.SourceModule[index] = sourceModule
}

func (d Displace) GetValue(x, y, z float64) float64 {
	if d.SourceModule[0] == nil || d.SourceModule == nil ||
		d.SourceModule[2] == nil || d.SourceModule == nil {
		panic("Displace must have 4 source modules.")
	}

	// Get the output values from the three displacement modules. Add each
	// value to the coorespnding coordinate in the input value.
	xDisplace := x + d.SourceModule[1].GetValue(x, y, z)
	yDisplace := y + d.SourceModule[2].GetValue(x, y, z)
	zDisplace := z + d.SourceModule[3].GetValue(x, y, z)

	//Retrieve the output value using the offsetted input value instead of
	// the orignal input value.
	return d.SourceModule[0].GetValue(xDisplace, yDisplace, zDisplace)
}

func (d Displace) SetXDisplaceModule(source Module) {
	d.SourceModule[1] = source
}

func (d Displace) SetYDisplaceModule(source Module) {
	d.SourceModule[2] = source
}

func (d Displace) SetZDisplaceModule(source Module) {
	d.SourceModule[3] = source
}

func (d Displace) SetXYZDisplaceModule(sourceX, sourceY, sourceZ Module) {
	d.SourceModule[1] = sourceX
	d.SourceModule[2] = sourceY
	d.SourceModule[3] = sourceZ
}
