package skin

import "image"

func RenderFace(skin *image.NRGBA, opts Options) *image.NRGBA {
	uvScale := skin.Bounds().Max.Y / 64
	output := removeTransparency(extract(skin, 8, 8, 8, 8, uvScale))

	if opts.Overlay && !IsOldSkin(skin) {
		output = composite(output, extract(skin, 40, 8, 8, 8, uvScale), 0, 0)
	}

	return scale(output, opts.Scale/uvScale)
}
