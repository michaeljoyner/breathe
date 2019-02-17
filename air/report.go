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
	Message   string
}

func (r *Report) setMessage() {
	if r.Index < 0 {
		r.Message = "Bad input, don't believe anything"
		return
	}

	if r.Index < 50 {
		r.Message = "Super good, breathe it in."
		return
	}

	if r.Index < 100 {
		r.Message = "Acceptable, nothing to worry about."
		return
	}

	if r.Index < 150 {
		r.Message = "Very sensitive people should take caution, fine for most people."
		return
	}

	if r.Index < 200 {
		r.Message = "The air is not good. Limit outdoor exercise."
		return
	}

	if r.Index < 300 {
		r.Message = "Pollution is very bad. Completely avoid outdoor exercise."
		return
	}

	r.Message = "Get the hell out of there, we are all doomed"

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

	report := Report{
		Station:   "Taichung City (Zhongming rd station)",
		Index:     r.Data.Index,
		CheckedAt: r.Data.Time.S,
	}

	report.setMessage()

	return report, nil

}
