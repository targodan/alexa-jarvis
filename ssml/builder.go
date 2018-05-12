package ssml

import (
	"strconv"
	"time"
)

type Builder struct {
	previousBuilder *Builder
	lastElement     Element
}

func Speak() Builder {
	return Builder{
		previousBuilder: nil,
		lastElement:     newElement("speak"),
	}
}

func (b Builder) String() string {
	return b.lastElement.String()
}

func (b Builder) Element() Element {
	return b.lastElement
}

func (b Builder) Text(text string) Builder {
	b.lastElement.addSubElement(&textElement{
		content: text,
	})
	return b
}

func (b Builder) AmazonEffect(effect AmazonEffect) Builder {
	elem := newElement("amazon:effect")
	elem.addAttribute(&attribute{
		name:  "name",
		value: string(effect),
	})
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Audio(url string) Builder {
	elem := newElement("audio")
	elem.addAttribute(&attribute{
		name:  "src",
		value: url,
	})
	b.lastElement.addSubElement(elem)
	return b
}

func (b Builder) Break(strength BreakStrength, time time.Duration) Builder {
	elem := newElement("break")
	elem.addAttribute(&attribute{
		name:  "strength",
		value: string(strength),
	})
	elem.addAttribute(&attribute{
		name:  "time",
		value: strconv.Itoa(int(time.Seconds()*1000)) + "ms",
	})
	b.lastElement.addSubElement(elem)
	return b
}

func (b Builder) Emphasis(level EmphasisLevel) Builder {
	elem := newElement("emphasis")
	elem.addAttribute(&attribute{
		name:  "level",
		value: string(level),
	})
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Paragraph() Builder {
	elem := newElement("p")
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Phoneme(alphabet PhonemeAlphabet, pronounciation string) Builder {
	elem := newElement("phoneme")
	elem.addAttribute(&attribute{
		name:  "alphabet",
		value: string(alphabet),
	})
	elem.addAttribute(&attribute{
		name:  "ph",
		value: pronounciation,
	})
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Prosody(rate SpeechRate, pitch Pitch, volume Volume) Builder {
	elem := newElement("prosody")
	elem.addAttribute(&attribute{
		name:  "rate",
		value: string(rate),
	})
	elem.addAttribute(&attribute{
		name:  "pitch",
		value: string(pitch),
	})
	elem.addAttribute(&attribute{
		name:  "volume",
		value: string(volume),
	})
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Sentence() Builder {
	elem := newElement("s")
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) SayAs(interpretation Interpretation) Builder {
	elem := newElement("say-as")
	elem.addAttribute(&attribute{
		name:  "interpret-as",
		value: string(interpretation),
	})
	b.lastElement.addSubElement(elem)
	return Builder{
		previousBuilder: &b,
		lastElement:     elem,
	}
}

func (b Builder) Duration(dur time.Duration) Builder {
	return b.SayAs(InterpretAsTime).Text(dur.String()).End()
}

func (b Builder) Date(date time.Time, sayYear, sayMonth, sayDay bool) Builder {
	text := ""
	if sayYear {
		text += date.Format("2006")
	}
	if sayMonth {
		text += date.Format("01")
	}
	if sayDay {
		text += date.Format("02")
	}
	return b.SayAs(InterpretAsDate).Text(text).End()
}

func (b Builder) FullDate(date time.Time) Builder {
	return b.Date(date, true, true, true)
}

func (b Builder) Word(role WordRole, word string) Builder {
	elem := newElement("w")
	elem.addAttribute(&attribute{
		name:  "role",
		value: string(role),
	})
	elem.addSubElement(&textElement{
		content: word,
	})
	b.lastElement.addSubElement(elem)
	return b
}

func (b Builder) Append(elem Element) Builder {
	if elem.Name() == "speak" {
		for _, subElem := range elem.SubElements() {
			b.lastElement.addSubElement(subElem)
		}
	} else {
		b.lastElement.addSubElement(elem)
	}
	return b
}

func (b Builder) End() Builder {
	return *b.previousBuilder
}
