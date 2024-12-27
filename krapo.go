package krapo

import (
	"time"

	_ "github.com/fluffy-melli/krapo/RDR_CMP" // https://www.data.go.kr/tcs/dss/selectApiDataDetailView.do?publicDataPk=15056924
)

func Time() string {
	currentTime := time.Now()
	return currentTime.Format("20060102")
}
