package main

import (
	"encoding/xml"
	"html"
)

type URLEncodedString struct {
	s string
}

func (u *URLEncodedString) String() string {
	return u.s
}

func (u *URLEncodedString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	*u = URLEncodedString{s: html.UnescapeString(s)}

	return nil
}
