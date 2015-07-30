package gpxjson

import (
	"encoding/json"
	"encoding/xml"
)

type GPX struct {
	Trkseg []Trkseg `xml:"trk>trkseg" json:"segments"`
}

type Trkseg struct {
	Trkpt []struct {
		Lat  float32 `xml:"lat,attr" json:"lat"`
		Lon  float32 `xml:"lon,attr" json:"lon"`
		Ele  float32 `xml:"ele" json:"ele,omitempty"`
		Time string  `xml:"time" json:"time,omitempty"`
	} `xml:"trkpt"`
}

func (t *Trkseg) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Trkpt)
}

// JSON return a minimal version of a gpx
func Convert(in []byte) ([]byte, error) {
	var g GPX
	err := xml.Unmarshal(in, &g)

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(g)
	return out, err
}
