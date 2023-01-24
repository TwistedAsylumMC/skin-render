package skin

import (
	"image"
)

func RenderFrontBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	slimOffset := getSlimOffset(opts.Slim)

	var (
		frontHead  *image.NRGBA = removeTransparency(extract(skin, 8, 8, 8, 8, uvScale))
		frontTorso *image.NRGBA = removeTransparency(extract(skin, 20, 20, 8, 12, uvScale))
		leftArm    *image.NRGBA = nil
		rightArm   *image.NRGBA = removeTransparency(extract(skin, 44, 20, 4-slimOffset, 12, uvScale))
		leftLeg    *image.NRGBA = nil
		rightLeg   *image.NRGBA = removeTransparency(extract(skin, 4, 20, 4, 12, uvScale))
	)

	if IsOldSkin(skin) {
		leftArm = flipHorizontal(rightArm)
		leftLeg = flipHorizontal(rightLeg)
	} else {
		leftArm = removeTransparency(extract(skin, 36, 52, 4-slimOffset, 12, uvScale))
		leftLeg = removeTransparency(extract(skin, 20, 52, 4, 12, uvScale))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			frontHead = composite(frontHead, extract(overlaySkin, 40, 8, 8, 8, uvScale), 0, 0)
			frontTorso = composite(frontTorso, extract(overlaySkin, 20, 36, 8, 12, uvScale), 0, 0)
			leftArm = composite(leftArm, extract(overlaySkin, 52, 52, 4-slimOffset, 64, uvScale), 0, 0)
			rightArm = composite(rightArm, extract(overlaySkin, 44, 36, 4-slimOffset, 48, uvScale), 0, 0)
			leftLeg = composite(leftLeg, extract(overlaySkin, 4, 52, 4, 12, uvScale), 0, 0)
			rightLeg = composite(rightLeg, extract(overlaySkin, 4, 36, 4, 12, uvScale), 0, 0)
		}
	}

	output := image.NewNRGBA(image.Rect(0, 0, 16*uvScale, 32*uvScale))

	// Face
	output = composite(output, frontHead, 4*uvScale, 0)

	// Torso
	output = composite(output, frontTorso, 4*uvScale, 8*uvScale)

	// Left Arm
	output = composite(output, leftArm, 12*uvScale, 8*uvScale)

	// Right Arm
	output = composite(output, rightArm, slimOffset*uvScale, 8*uvScale)

	// Left Leg
	output = composite(output, leftLeg, 8*uvScale, 20*uvScale)

	// Right Leg
	output = composite(output, rightLeg, 4*uvScale, 20*uvScale)

	return scale(output, opts.Scale/uvScale)
}
