package MID_FCST

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/fluffy-melli/krapo"
)

type Response struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			DataType   string `json:"dataType"`
			PageNo     int    `json:"pageNo"`
			NumOfRows  int    `json:"numOfRows"`
			TotalCount int    `json:"totalCount"`
			Items      struct {
				Item []struct {
					WfSv string `json:"wfSv"`
				} `json:"item"`
			} `json:"items"`
		} `json:"body"`
	} `json:"response"`
}

func GetFcst(apikey string) ([]string, error) {
	var tmfc string
	currentTime := time.Now()
	if currentTime.Hour()*60+currentTime.Minute() >= 18*60 {
		tmfc = krapo.Time() + "1800"
	} else if currentTime.Hour()*60+currentTime.Minute() >= 6*60 {
		tmfc = krapo.Time() + "0600"
	} else {
		tmfc = krapo.LTime(1) + "1800"
	}
	URL := "http://apis.data.go.kr/1360000/MidFcstInfoService/getMidFcst"
	params := url.Values{}
	params.Add("serviceKey", apikey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "1")
	params.Add("dataType", "JSON")
	params.Add("stnId", "108")
	params.Add("tmFc", tmfc)
	resp, err := http.Get(URL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	if response.Response.Header.ResultCode != "00" {
		return nil, fmt.Errorf(response.Response.Header.ResultMsg)
	}
	respond := make([]string, 0)
	for _, item := range response.Response.Body.Items.Item {
		respond = append(respond, item.WfSv)
	}
	return respond, nil
}
