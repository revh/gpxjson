package gpxjson

import "testing"

var sample = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<gpx version="1.1" creator="Endomondo.com" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www.garmin.com/xmlschemas/GpxExtensionsv3.xsd http://www.garmin.com/xmlschemas/TrackPointExtension/v1 http://www.garmin.com/xmlschemas/TrackPointExtensionv1.xsd" xmlns="http://www.topografix.com/GPX/1/1" xmlns:gpxtpx="http://www.garmin.com/xmlschemas/TrackPointExtension/v1" xmlns:gpxx="http://www.garmin.com/xmlschemas/GpxExtensions/v3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <metadata>
    <author>
      <name>Lorenzo Greco</name>
      <email id="lorenzo.rev" domain="gmail.com"/>
    </author>
    <link href="http://www.endomondo.com">
      <text>Endomondo</text>
    </link>
    <time>2015-07-28T19:02:36Z</time>
  </metadata>
  <trk>
    <src>http://www.endomondo.com/</src>
    <link href="https://secreturl.endomondo.com/workouts/XXXX/YYYY">
      <text>endomondo</text>
    </link>
    <type>CYCLING_SPORT</type>
    <trkseg>
      <trkpt lat="43.76319" lon="11.149139">
        <time>2015-07-25T07:17:59Z</time>
      </trkpt>
    </trkseg>
    <trkseg>
      <trkpt lat="43.76319" lon="11.149139">
        <ele>95.1</ele>
        <time>2015-07-25T07:18:00Z</time>
      </trkpt>
      <trkpt lat="43.76319" lon="11.149139">
        <ele>95.2</ele>
        <time>2015-07-25T07:18:00Z</time>
      </trkpt>
    </trkseg>
  </trk>
</gpx>`)

func TestUnmarshal(t *testing.T) {
	if _, err := JSON(sample); err != nil {
		t.Errorf(`JSON() error %s`, err)
	}
}
