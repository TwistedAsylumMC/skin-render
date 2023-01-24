package skin

import (
	"image"
)

func RenderBackBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	slimOffset := getSlimOffset(opts.Slim)

	var (
		backHead     *image.NRGBA = removeTransparency(extract(skin, 24, 8, 8, 8, uvScale))
		backTorso    *image.NRGBA = removeTransparency(extract(skin, 32, 20, 8, 12, uvScale))
		backLeftArm  *image.NRGBA = nil
		backRightArm *image.NRGBA = removeTransparency(extract(skin, 52-slimOffset, 20, 4-slimOffset, 12, uvScale))
		backLeftLeg  *image.NRGBA = nil
		backRightLeg *image.NRGBA = removeTransparency(extract(skin, 12, 20, 4, 12, uvScale))
	)

	if IsOldSkin(skin) {
		backLeftArm = flipHorizontal(backRightArm)
		backLeftLeg = flipHorizontal(backRightLeg)
	} else {
		backLeftArm = removeTransparency(extract(skin, 44-slimOffset, 52, 4-slimOffset, 12, uvScale))
		backLeftLeg = removeTransparency(extract(skin, 28, 52, 4, 12, uvScale))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			backHead = composite(backHead, extract(overlaySkin, 56, 8, 8, 8, uvScale), 0, 0)
			backTorso = composite(backTorso, extract(overlaySkin, 32, 36, 8, 12, uvScale), 0, 0)
			backLeftArm = composite(backLeftArm, extract(overlaySkin, 60-slimOffset, 52, 4-slimOffset, 64, uvScale), 0, 0)
			backRightArm = composite(backRightArm, extract(overlaySkin, 52-slimOffset, 36, 4-slimOffset, 48, uvScale), 0, 0)
			backLeftLeg = composite(backLeftLeg, extract(overlaySkin, 12, 52, 8, 64, uvScale), 0, 0)
			backRightLeg = composite(backRightLeg, extract(overlaySkin, 12, 36, 8, 48, uvScale), 0, 0)
		}
	}

	output := image.NewNRGBA(image.Rect(0, 0, 16*uvScale, 32*uvScale))

	// Face
	output = composite(output, backHead, 4*uvScale, 0)

	// Torso
	output = composite(output, backTorso, 4*uvScale, 8*uvScale)

	// Left Arm
	output = composite(output, backLeftArm, slimOffset*uvScale, 8*uvScale)

	// Right Arm
	output = composite(output, backRightArm, 12*uvScale, 8*uvScale)

	// Left Leg
	output = composite(output, backLeftLeg, 4*uvScale, 20*uvScale)

	// Right Leg
	output = composite(output, backRightLeg, 8*uvScale, 20*uvScale)

	return scale(output, opts.Scale/uvScale)
}
