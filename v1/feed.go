package arxiv

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	XMLName xml.Name  `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	ID      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated TimeStr   `xml:"updated"`
	Author  []*Person `xml:"author"`
	Entry   []*Entry  `xml:"entry"`
}

type Entry struct {
	XMLName         xml.Name
	Doi             string    `xml:"doi"`
	Title           string    `xml:"title"`
	ID              string    `xml:"id"`
	Link            []Link    `xml:"link"`
	Published       TimeStr   `xml:"published"`
	Updated         TimeStr   `xml:"updated"`
	Comment         string    `xml:"comment"`
	Author          []*Person `xml:"author"`
	Summary         *Text     `xml:"summary"`
	Content         *Text     `xml:"content"`
	PrimaryCategory *Class    `xml:"primary_category,omitempty"`
	Category        []*Class  `xml:"category,omitempty"`
}

type Class struct {
	XMLName xml.Name
	Term    Category `xml:"term,attr"`
}

type Link struct {
	XMLName  xml.Name
	Rel      string `xml:"rel,attr,omitempty"`
	Href     string `xml:"href,attr"`
	Type     string `xml:"type,attr,omitempty"`
	HrefLang string `xml:"hreflang,attr,omitempty"`
	Title    string `xml:"title,attr,omitempty"`
	Length   uint   `xml:"length,attr,omitempty"`
}

type Person struct {
	XMLName xml.Name
	Name    string `xml:"name"`
	URI     string `xml:"uri,omitempty"`
	Email   string `xml:"email,omitempty"`
}

type Text struct {
	XMLName xml.Name
	Type    string `xml:"type,attr"`
	Body    string `xml:",chardata"`
}

type TimeStr string

func Time(t time.Time) TimeStr {
	return TimeStr(t.Format("2006-01-02T15:04:05-07:00"))
}
