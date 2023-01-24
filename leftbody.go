package skin

import (
	"image"
)

func RenderLeftBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	slimOffset := getSlimOffset(opts.Slim)

	var (
		leftHead    *image.NRGBA = removeTransparency(extract(skin, 16, 8, 8, 8, uvScale))
		leftLeftArm *image.NRGBA = nil
		leftLeftLeg *image.NRGBA = nil
	)

	if IsOldSkin(skin) {
		leftLeftArm = flipHorizontal(removeTransparency(extract(skin, 40, 20, 4, 12, uvScale)))
		leftLeftLeg = flipHorizontal(removeTransparency(extract(skin, 0, 20, 4, 12, uvScale)))
	} else {
		leftLeftArm = removeTransparency(extract(skin, 40-slimOffset, 52, 4, 12, uvScale))
		leftLeftLeg = removeTransparency(extract(skin, 24, 52, 4, 12, uvScale))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			leftHead = composite(leftHead, extract(overlaySkin, 48, 8, 8, 8, uvScale), 0, 0)
			leftLeftArm = composite(leftLeftArm, extract(overlaySkin, 56-slimOffset, 52, 4, 12, uvScale), 0, 0)
			leftLeftLeg = composite(leftLeftLeg, extract(overlaySkin, 8, 52, 4, 12, uvScale), 0, 0)
		}
	}

	output := image.NewNRGBA(image.Rect(0, 0, 8*uvScale, 32*uvScale))

	// Left Head
	output = composite(output, leftHead, 0, 0)

	// Left Arm
	output = composite(output, leftLeftArm, 2*uvScale, 8*uvScale)

	// Left Leg
	output = composite(output, leftLeftLeg, 2*uvScale, 20*uvScale)

	return scale(output, opts.Scale/uvScale)
}
