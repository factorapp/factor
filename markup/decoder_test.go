package markup

import (
	"bytes"
	"testing"
)

const (
	fooXML = `
<div>
    <h1>Hello</h1>
    <input type="text" onchange="OnChange" />
    <button type="button">Click Me!</button>
	<!-- Write your comments here -->
	<Bar />
	<div />
    end
</div>
    `
	invalidXML = `
</div>
    `
	invalidXMLTag = `
<div></span>
    `
	invalidXMLText = `
Hello World
    `
)

func TestNewDecoder(t *testing.T) {
	r := bytes.NewBufferString(fooXML)
	newDecoder(r)
}

func TestDecoderDecode(t *testing.T) {
	r := bytes.NewBufferString(fooXML)
	d := newDecoder(r)

	n, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(n.Markup())
}

func TestDecoderDecodeEmpty(t *testing.T) {
	r := bytes.NewBufferString("")
	d := newDecoder(r)

	_, err := d.Decode()
	if err == nil {
		t.Error("should error")
	}
}

func TestDecoderDecodeInvalid(t *testing.T) {
	r := bytes.NewBufferString(invalidXML)
	d := newDecoder(r)

	_, err := d.Decode()
	if err == nil {
		t.Error("should error")
	}

	r = bytes.NewBufferString(invalidXMLTag)
	d = newDecoder(r)

	_, err = d.Decode()
	if err == nil {
		t.Error("should error")
	}

	r = bytes.NewBufferString(invalidXMLText)
	d = newDecoder(r)

	_, err = d.Decode()
	if err == nil {
		t.Error("should error")
	}
}
