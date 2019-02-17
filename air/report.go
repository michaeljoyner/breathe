package air

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Report struct {
	Station   string
	Index     int
	CheckedAt string
}

type aqiData struct {
	Index int        `json:"aqi"`
	Time  timeString `json:"time"`
}

type timeString struct {
	S string `json:"s"`
}

type apiResp struct {
	Status string  `json:"status"`
	Data   aqiData `json:"data"`
}

//GetReport returns the current air quality index report
func GetReport() (Report, error) {
	resp, err := http.Get("http://api.waqi.info/feed/taichung/?token=" + os.Getenv("AQI_KEY"))
	if err != nil {
		return Report{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Report{}, err
	}

	var r apiResp

	if err = json.Unmarshal(b, &r); err != nil {
		return Report{}, err
	}

	return Report{
		Station:   "Taichung City (Zhongming rd station)",
		Index:     r.Data.Index,
		CheckedAt: r.Data.Time.S,
	}, nil

}
