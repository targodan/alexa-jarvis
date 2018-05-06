package ssml

type BreakStrength string

const (
	BreakStrengthNone    BreakStrength = "none"
	BreakStrengthXWeak                 = "x-weak"
	BreakStrengthWeak                  = "weak"
	BreakStrengthMedium                = "medium"
	BreakStrengthStrong                = "strong"
	BreakStrengthXStrong               = "x-strong"
)
