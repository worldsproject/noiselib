package noiselib

import (
	"image"
	"image/color"
	"image/png"
	// "image/rectangle"
	"math"
	"os"
	// "image/png"
)

type RenderImage struct {
	CosAzimuth, CosElev, SinAzimuth, SinElev     float64
	LightAzimuth, LightBrightness, LightContrast float64
	LightElev, LightIntensity                    float64
	Gradient                                     GradientColor
	LightEnabled, WrapEnabled, RecalculateLight  bool
	LightColor                                   color.RGBA
	BackgroundImage                              image.RGBA
	DestinationImage                             image.RGBA
	NoiseMap                                     [][]float64
}

func (r *RenderImage) BuildGrayscaleGradient() {
	r.Gradient.ClearGradient()
	r.Gradient.AddGradientPoint(-1.0, color.RGBA{0, 0, 0, 255})
	r.Gradient.AddGradientPoint(1.0, color.RGBA{255, 255, 255, 255})
}

func (r *RenderImage) BuildTerrainGradient() {
	r.Gradient.ClearGradient()
	r.Gradient.AddGradientPoint(-1.00, color.RGBA{0, 0, 128, 255})
	r.Gradient.AddGradientPoint(-0.20, color.RGBA{32, 64, 128, 255})
	r.Gradient.AddGradientPoint(-0.04, color.RGBA{64, 96, 192, 255})
	r.Gradient.AddGradientPoint(-0.02, color.RGBA{192, 192, 128, 255})
	r.Gradient.AddGradientPoint(0.00, color.RGBA{0, 192, 0, 255})
	r.Gradient.AddGradientPoint(0.25, color.RGBA{192, 192, 0, 255})
	r.Gradient.AddGradientPoint(0.50, color.RGBA{160, 96, 64, 255})
	r.Gradient.AddGradientPoint(0.75, color.RGBA{128, 255, 255, 255})
	r.Gradient.AddGradientPoint(1.00, color.RGBA{255, 255, 255, 255})
}

func (r *RenderImage) Render() {
	width := len(r.NoiseMap)
	height := len(r.NoiseMap[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			destColor := r.Gradient.GetColor(r.NoiseMap[x][y])

			lightIntensity := 0.0

			if r.LightEnabled {
				var xLeftOffset, xRightOffset, yUpOffset, yDownOffset int

				if r.WrapEnabled {
					if x == 0 {
						xLeftOffset = width - 1
						xRightOffset = 1
					} else if x == width-1 {
						xLeftOffset = -1
						xRightOffset = -(width - 1)
					} else {
						xLeftOffset = -1
						xRightOffset = 1
					}

					if y == 0 {
						yDownOffset = height - 1
						yUpOffset = 1
					} else if y == height-1 {
						yDownOffset = -1
						yUpOffset = -(height - 1)
					} else {
						yDownOffset = -1
						yUpOffset = 1
					}
				} else {
					if x == 0 {
						xLeftOffset = 0
						xRightOffset = 1
					} else if x == width-1 {
						xLeftOffset = -1
						xRightOffset = 0
					} else {
						xLeftOffset = -1
						xRightOffset = 1
					}
					if y == 0 {
						yDownOffset = 0
						yUpOffset = 1
					} else if y == height-1 {
						yDownOffset = -1
						yUpOffset = 0
					} else {
						yDownOffset = -1
						yUpOffset = 1
					}
				}

				nc := r.NoiseMap[x][y]
				nl := r.NoiseMap[x+xLeftOffset][y]
				nr := r.NoiseMap[x+xRightOffset][y]
				nd := r.NoiseMap[x][y+yDownOffset]
				nu := r.NoiseMap[x][y+yUpOffset]

				lightIntensity = r.calcLightIntensity(nc, nl, nr, nd, nu)
				lightIntensity *= r.LightBrightness
			} else {
				lightIntensity = 1.0
			}

			backgroundColor := color.RGBA{255, 255, 255, 255}

			if &r.BackgroundImage != nil {
				backgroundColor = r.BackgroundImage.At(x, y).(color.RGBA)
			}

			newColor := r.calcDestColor(destColor, backgroundColor, lightIntensity)
			r.DestinationImage.Set(x, y, newColor)
		}
	}

	file, err := os.Create("test.png")

	if err != nil {
		panic("Something went wrong with opening a file.")
	}
	defer file.Close()

	png.Encode(file, r.DestinationImage.SubImage(image.Rect(0, 0, width, height)))
}

func (r *RenderImage) SetLightAzimuth(azimuth float64) {
	r.LightAzimuth = azimuth
	r.RecalculateLight = true
}

func (r *RenderImage) SetLightBrightness(brightness float64) {
	r.LightBrightness = brightness
	r.RecalculateLight = true
}

func (r *RenderImage) SetLightContrast(contrast float64) {
	if contrast < 0.0 {
		panic("Contrast must be greater than 0.")
	}

	r.LightContrast = contrast
	r.RecalculateLight = true
}

func (r *RenderImage) SetLightElev(elev float64) {
	r.LightElev = elev
	r.RecalculateLight = true
}

func (r *RenderImage) SetLightIntensity(intensity float64) {
	if intensity < 0.0 {
		panic("Light Intensity must be greater than 0")
	}

	r.LightIntensity = intensity
	r.RecalculateLight = true
}

func (r *RenderImage) calcDestColor(source, background color.RGBA, lightValue float64) color.RGBA {
	sourceRed := float64(source.R) / 255.0
	sourceGreen := float64(source.G) / 255.0
	sourceBlue := float64(source.B) / 255.0
	sourceAlpha := float64(source.A) / 255.0

	backgroundRed := float64(background.R) / 255.0
	backgroundGreen := float64(background.G) / 255.0
	backgroundBlue := float64(background.B) / 255.0

	red := LinearInterp(backgroundRed, sourceRed, sourceAlpha)
	green := LinearInterp(backgroundGreen, sourceGreen, sourceAlpha)
	blue := LinearInterp(backgroundBlue, sourceBlue, sourceAlpha)

	if r.LightEnabled {
		lightRed := lightValue * float64(r.LightColor.R) / 255.0
		lightGreen := lightValue * float64(r.LightColor.G) / 255.0
		lightBlue := lightValue * float64(r.LightColor.B) / 255.0

		red *= lightRed
		green *= lightGreen
		blue *= lightBlue
	}

	red = ClampValueFloat(red, 0.0, 1.0)
	green = ClampValueFloat(green, 0.0, 1.0)
	blue = ClampValueFloat(blue, 0.0, 1.0)

	cRed := uint8(red*255.0) & 0xff
	cGreen := uint8(green*255.0) & 0xff
	cBlue := uint8(blue*255.0) & 0xff
	return color.RGBA{cRed, cGreen, cBlue, uint8(math.Max(float64(source.A), float64(background.A)))}
}

func (r *RenderImage) calcLightIntensity(center, left, right, up, down float64) float64 {
	if r.RecalculateLight {
		r.CosAzimuth = math.Cos(r.LightAzimuth * DegToRad)
		r.SinAzimuth = math.Sin(r.LightAzimuth * DegToRad)
		r.CosElev = math.Cos(r.LightElev * DegToRad)
		r.SinElev = math.Sin(r.LightElev * DegToRad)
		r.RecalculateLight = false
	}

	IMax := 1.0
	io := IMax * SQRT_2 * r.SinElev / 2.0
	ix := (IMax - io) * r.LightContrast * SQRT_2 * r.CosElev * r.CosAzimuth
	iy := (IMax - io) * r.LightContrast * SQRT_2 * r.CosElev * r.SinAzimuth
	intensity := (ix*(left-right) + iy*(down-up) + io)

	if intensity < 0.0 {
		return 0.0
	} else {
		return intensity
	}

}
