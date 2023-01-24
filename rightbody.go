package skin

import (
	"image"
)

func RenderRightBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	var (
		rightHead     *image.NRGBA = removeTransparency(extract(skin, 0, 8, 8, 8, uvScale))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, 40, 20, 4, 12, uvScale))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, 0, 20, 4, 12, uvScale))
	)

	if opts.Overlay && !IsOldSkin(skin) {
		overlaySkin := fixTransparency(skin)

		rightHead = composite(rightHead, extract(overlaySkin, 32, 8, 8, 8, uvScale), 0, 0)
		rightRightArm = composite(rightRightArm, extract(overlaySkin, 40, 36, 4, 12, uvScale), 0, 0)
		rightRightLeg = composite(rightRightLeg, extract(overlaySkin, 0, 36, 4, 12, uvScale), 0, 0)
	}

	output := image.NewNRGBA(image.Rect(0, 0, 8*uvScale, 32*uvScale))

	// Right Head
	output = composite(output, rightHead, 0, 0)

	// Right Arm
	output = composite(output, rightRightArm, 2*uvScale, 8*uvScale)

	// Right Leg
	output = composite(output, rightRightLeg, 2*uvScale, 20*uvScale)

	return scale(output, opts.Scale/uvScale)
}
