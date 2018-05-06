package ssml

type Interpretation string

const (
	InterpretAsCharacters   Interpretation = "characters"
	InterpretAsSpellOut                    = "spell-out"
	InterpretAsCardinal                    = "cardinal"
	InterpretAsNumber                      = "number"
	InterpretAsOrdinal                     = "ordinal"
	InterpretAsDigits                      = "digits"
	InterpretAsFraction                    = "fraction"
	InterpretAsUnit                        = "unit"
	InterpretAsDate                        = "date"
	InterpretAsTime                        = "time"
	InterpretAsTelephone                   = "telephone"
	InterpretAsAddress                     = "address"
	InterpretAsInterjection                = "interjection"
	InterpretAsExpletive                   = "expletive"
)
