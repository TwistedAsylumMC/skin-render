package skin

import "image"

func RenderHead(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	scaleDouble := float64(opts.Scale)
	output := image.NewNRGBA(image.Rect(0, 0, 16*opts.Scale, 19*opts.Scale-int(scaleDouble/2.0)-1))

	var (
		frontHead *image.NRGBA = removeTransparency(extract(skin, 8, 8, 8, 8, uvScale))
		topHead   *image.NRGBA = removeTransparency(extract(skin, 8, 0, 8, 8, uvScale))
		rightHead *image.NRGBA = removeTransparency(extract(skin, 0, 8, 8, 8, uvScale))
	)

	if opts.Overlay && !IsOldSkin(skin) {
		overlaySkin := fixTransparency(skin)

		frontHead = composite(frontHead, extract(overlaySkin, 40, 8, 8, 8, uvScale), 0, 0)
		topHead = composite(topHead, extract(overlaySkin, 40, 0, 8, 8, uvScale), 0, 0)
		rightHead = composite(rightHead, extract(overlaySkin, 32, 8, 8, 8, uvScale), 0, 0)
	}

	// Front Head
	output = compositeTransform(output, scale(frontHead, opts.Scale/uvScale), transformForward, 8*scaleDouble, 12*scaleDouble-1)

	// Top Head
	output = compositeTransform(output, scale(topHead, opts.Scale/uvScale), transformUp, -4*scaleDouble, 4*scaleDouble)

	// Right Head
	output = compositeTransform(output, scale(rightHead, opts.Scale/uvScale), transformRight, 0, 4*scaleDouble)

	return output
}
