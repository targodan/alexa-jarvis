package ssml

import "fmt"

type SpeechRate string

const (
	SpeechRateXSlow  SpeechRate = "x-slow"
	SpeechRateSlow              = "slow"
	SpeechRateMedium            = "medium"
	SpeechRateFast              = "fast"
	SpeechRateXFast             = "x-fast"
)

// RelativeSpeechRate returns a relative speech rate.
// 100 is the regular rate, 50 is half the speed and 200
// is twice the speed.
func RelativeSpeechRate(percent int) SpeechRate {
	return SpeechRate(fmt.Sprintf("%d%%", percent))
}

type Pitch string

const (
	PitchXLow   Pitch = "x-low"
	PitchLow          = "low"
	PitchMedium       = "medium"
	PitchHigh         = "high"
	PitchXHigh        = "x-high"
)

// RelativePitch returns a relative pitch.
// Positive 10 is ten percent higher, negative 10 is ten percent lower.
func RelativePitch(percentPlusMinus int) Pitch {
	return Pitch(fmt.Sprintf("%+d%%", percentPlusMinus))
}

type Volume string

const (
	VolumeSilent Volume = "silent"
	VolumeXSoft         = "x-soft"
	VolumeSoft          = "soft"
	VolumeMedium        = "medium"
	VolumeLoud          = "loud"
	VolumeXLoud         = "x-loud"
)

// RelativeVolume returns a relative volume.
// Positive 10 is ten db louder, negative 10 is db less loud.
func RelativeVolume(dbPlusMinus int) Volume {
	return Volume(fmt.Sprintf("%+ddb", dbPlusMinus))
}
