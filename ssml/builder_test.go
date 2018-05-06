package ssml

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	fmt.Println(
		Speak().
			Text("helo, ").
			AmazonEffect("whispered").
			Text("this is whispered").
			End().
			Text(", but this is not.").String(),
	)
}
