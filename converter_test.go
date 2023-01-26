package xmltojson

import "testing"


type Notes struct {
	To        string    `xml:"to"      json:"to"`; 
	From      string    `xml:"from"    json:"from"`
	Heading   string    `xml:"heading" json:"heading"`
	Body      string    `xml:"body"    json:"note_body"`
}


func TestJsonFromXml(t *testing.T) {

	notes := &Notes{}

	json := JsonFromXml("remiders.xml", notes)

	want := "{\"to\":\"Tove\",\"from\":\"Jani\",\"heading\":\"Reminder\",\"note_body\":\"Don't forget me this weekend!\"}"

	if (json != want) {
		t.Errorf("got %q, wanted %q", json, want)
	}
}