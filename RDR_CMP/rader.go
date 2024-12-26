package RDR_CMP

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
					Images string `json:"rdr-img-file"`
				} `json:"item"`
			} `json:"items"`
		} `json:"body"`
	} `json:"response"`
}

func Time() string {
	currentTime := time.Now()
	return currentTime.Format("20060102")
}

// time : YYYYMMDD | 20241226
func GetImages(apikey, time string) ([]string, error) {
	URL := "http://apis.data.go.kr/1360000/RadarImgInfoService/getCmpImg"
	params := url.Values{}
	params.Add("serviceKey", apikey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "1")
	params.Add("dataType", "JSON")
	params.Add("data", "CMP_WRC")
	params.Add("time", time)
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
	images := response.Response.Body.Items.Item[0].Images
	images = strings.ReplaceAll(images, "[", "")
	images = strings.ReplaceAll(images, "]", "")
	return strings.Split(images, ","), nil
}
