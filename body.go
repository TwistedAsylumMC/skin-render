package skin

import "image"

func RenderBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	scaleDouble := float64(opts.Scale)
	slimOffset := getSlimOffset(opts.Slim)

	output := image.NewNRGBA(image.Rect(0, 0, 20*opts.Scale, 45*opts.Scale+int(scaleDouble*(1.0/16.0))))

	var (
		frontHead     *image.NRGBA = removeTransparency(extract(skin, 8, 8, 8, 8, uvScale))
		topHead       *image.NRGBA = removeTransparency(extract(skin, 8, 0, 8, 8, uvScale))
		rightHead     *image.NRGBA = removeTransparency(extract(skin, 0, 8, 8, 8, uvScale))
		frontTorso    *image.NRGBA = removeTransparency(extract(skin, 20, 20, 8, 12, uvScale))
		frontLeftArm  *image.NRGBA = nil
		topLeftArm    *image.NRGBA = nil
		frontRightArm *image.NRGBA = removeTransparency(extract(skin, 44, 20, 4-slimOffset, 12, uvScale))
		topRightArm   *image.NRGBA = removeTransparency(extract(skin, 44, 16, 4-slimOffset, 4, uvScale))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, 40, 20, 4, 12, uvScale))
		frontLeftLeg  *image.NRGBA = nil
		frontRightLeg *image.NRGBA = removeTransparency(extract(skin, 4, 20, 4, 12, uvScale))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, 0, 20, 4, 12, uvScale))
	)

	if IsOldSkin(skin) {
		frontLeftArm = flipHorizontal(frontRightArm)
		topLeftArm = flipHorizontal(topRightArm)
		frontLeftLeg = flipHorizontal(frontRightLeg)
	} else {
		frontLeftArm = removeTransparency(extract(skin, 36, 52, 4-slimOffset, 12, uvScale))
		topLeftArm = removeTransparency(extract(skin, 36, 48, 4-slimOffset, 4, uvScale))
		frontLeftLeg = removeTransparency(extract(skin, 20, 52, 4, 12, uvScale))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			frontHead = composite(frontHead, extract(overlaySkin, 40, 8, 8, 8, uvScale), 0, 0)
			topHead = composite(topHead, extract(overlaySkin, 40, 0, 8, 8, uvScale), 0, 0)
			rightHead = composite(rightHead, extract(overlaySkin, 32, 8, 8, 8, uvScale), 0, 0)
			frontTorso = composite(frontTorso, extract(overlaySkin, 20, 36, 8, 12, uvScale), 0, 0)
			frontLeftArm = composite(frontLeftArm, extract(overlaySkin, 52, 52, 4-slimOffset, 64, uvScale), 0, 0)
			topLeftArm = composite(topLeftArm, extract(overlaySkin, 52, 48, 4-slimOffset, 4, uvScale), 0, 0)
			frontRightArm = composite(frontRightArm, extract(overlaySkin, 44, 36, 4-slimOffset, 48, uvScale), 0, 0)
			topRightArm = composite(topRightArm, extract(overlaySkin, 44, 48, 4-slimOffset, 4, uvScale), 0, 0)
			rightRightArm = composite(rightRightArm, extract(overlaySkin, 40, 36, 4, 12, uvScale), 0, 0)
			frontLeftLeg = composite(frontLeftLeg, extract(overlaySkin, 4, 52, 4, 12, uvScale), 0, 0)
			frontRightLeg = composite(frontRightLeg, extract(overlaySkin, 4, 36, 4, 12, uvScale), 0, 0)
			rightRightLeg = composite(rightRightLeg, extract(overlaySkin, 0, 36, 4, 12, uvScale), 0, 0)
		}
	}

	// Right Side of Right Leg
	output = compositeTransform(output, scale(rightRightLeg, opts.Scale/uvScale), transformRight, 4*scaleDouble, 23*scaleDouble)

	// Front of Right Leg
	output = compositeTransform(output, scale(frontRightLeg, opts.Scale/uvScale), transformForward, 8*scaleDouble, 31*scaleDouble)

	// Front of Left Leg
	output = compositeTransform(output, scale(frontLeftLeg, opts.Scale/uvScale), transformForward, 12*scaleDouble, 31*scaleDouble)

	// Front of Torso
	output = compositeTransform(output, scale(frontTorso, opts.Scale/uvScale), transformForward, 8*scaleDouble, 19*scaleDouble)

	// Front of Right Arm
	output = compositeTransform(output, scale(frontRightArm, opts.Scale/uvScale), transformForward, float64(4+slimOffset)*scaleDouble, 19*scaleDouble-1)

	// Front of Left Arm
	output = compositeTransform(output, scale(frontLeftArm, opts.Scale/uvScale), transformForward, 16*scaleDouble, 21*scaleDouble-1)

	// Top of Left Arm
	output = compositeTransform(output, scale(topLeftArm, opts.Scale/uvScale), transformUp, -5*scaleDouble, 17*scaleDouble)

	// Right Side of Right Arm
	output = compositeTransform(output, scale(rightRightArm, opts.Scale/uvScale), transformRight, float64(slimOffset)*scaleDouble, float64(15-slimOffset)*scaleDouble)

	// Top of Right Arm
	output = compositeTransform(output, scale(topRightArm, opts.Scale/uvScale), transformUp, float64(-15+slimOffset)*scaleDouble, 15*scaleDouble)

	// Front of Head
	output = compositeTransform(output, scale(frontHead, opts.Scale/uvScale), transformForward, 10*scaleDouble, 13*scaleDouble-1)

	// Top of Head
	output = compositeTransform(output, scale(topHead, opts.Scale/uvScale), transformUp, -3*scaleDouble, 5*scaleDouble)

	// Right Side of Head
	output = compositeTransform(output, scale(rightHead, opts.Scale/uvScale), transformRight, 2*scaleDouble, 3*scaleDouble)

	return output
}
