// 发送http请求给python-judger
package infra

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/imag-er/wendingcup/app/submit/conf"
)

type JudgerResponse struct {
	Miou           float32   `json:"miou"`
	Miou_per_class []float32 `json:"iou_per_class"`
	Err            error     `json:"error"`
}

func JudgeRequest() (result *JudgerResponse) {
	c, _ := client.NewClient()
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	// req.SetRequestURI("http://localhost:8381/judge")
	req.SetRequestURI(conf.GetConf().Judge.JudgerAddress + "/judge")
	req.SetMethod("GET")

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		result.Err = err
		return result
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		result.Err = err
		return result
	}

	return result

}
