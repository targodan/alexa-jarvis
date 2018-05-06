package ssml

type WordRole string

const (
	WordRoleAmazonVerb             WordRole = "amazon:VB"
	WordRoleAmazonPastParticiple            = "amazon:VBD"
	WordRoleAmazonNoun                      = "amazon:NN"
	WordRoleAmazonAlternativeSense          = "amazon:SENSE_1"
)
