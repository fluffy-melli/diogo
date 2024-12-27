# krapo

## RDR_CMP (기상청_레이더영상 조회서비스) 
 - api token 발급 : https://www.data.go.kr/tcs/dss/selectApiDataDetailView.do?publicDataPk=15056924


```go
package main

import (
	"log"

	"github.com/fluffy-melli/krapo"
	"github.com/fluffy-melli/krapo/RDR_CMP"
	"github.com/fluffy-melli/krapo/render"
)

func main() {
	urls, err := RDR_CMP.GetImagesURL("API-TOKEN", krapo.Time())
	if err != nil {
		log.Fatalln(err)
	}
	// 옵션 (urls | GIF 속도 | url 일부 사진만 가져오기 - true 시 6개중 하나 | false 시 모든 사진 가져오기)
	gif, err := render.GIF(urls, 10, true)
	if err != nil {
		log.Fatalln(err)
	}
	err = render.Write("./test.gif", gif)
	if err != nil {
		log.Fatalln(err)
	}
}

```
 - GIF 렌더링 옵션이 (10 , true) 인 경우
<p align="left">
    <img src="./asset/example_RDR_CMP_low.gif" width="635" alt="RDR_CMP 사용예시 - 옵션 : (10, true)">
</p>

 - GIF 렌더링 옵션이 (2 , false) 인 경우
<p align="left">
    <img src="./asset/example_RDR_CMP_high.gif" width="635" alt="RDR_CMP 사용예시 - 옵션 : (2, false)">
</p>