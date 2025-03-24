package file

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/model"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/submit/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

type JudgeTask struct {
	SubmitId   string
	FilePath   string
	UploadTime time.Time
	TeamId     string
}

// 任务队列通道 value: Id
var TaskQueue = make(chan JudgeTask, 10)

// 互斥锁，用于保护共享资源
var mutex sync.Mutex

// 处理文件任务的工作函数
func judgeWorker() {
	for fileId := range TaskQueue {
		mutex.Lock()
		judge(fileId)
		mutex.Unlock()
	}
}

func onStartJudging(task JudgeTask) {
	klog.Info("处理任务: ", task.SubmitId)
	mysql.DB.Model(&model.Submit{}).Where("id = ?", task.SubmitId).
		Update("status", model.StatusPending).
		Update("message", "判题中")
}
func onProcJudging(task JudgeTask) (score float32, err error) {

	klog.Info("处理中: ", task.SubmitId)

	resp := infra.JudgeRequest()
	if resp.Err != nil {
		return 0, resp.Err
	}

	klog.Infof("处理结果: %v", resp)
	return resp.Miou, nil

}

func onCompleted(task JudgeTask, score float32) {
	klog.Info("处理完成: ", task)

	resp, err := infra.BoardClient.AppendJudgeResult(context.Background(), &board.AppendJudgeResultRequest{
		JudgeResult: &board.JudgeResult{
			TeamId:          task.TeamId,
			FileUploadTime:  task.UploadTime.Format("2006-01-02 15:04:05"),
			JudgeResultTime: time.Now().Format("2006-01-02 15:04:05"),
			Score:           score,
		},
	})
	if err != nil {
		klog.Error("处理错误: ", err)
	}
	klog.Info("处理完成: ", resp)

	mysql.DB.Model(&model.Submit{}).Where("id = ?", task.SubmitId).
		Update("status", model.StatusCompleted).
		Update("message", "处理完成").
		Update("score", score)
}

func onFailed(task JudgeTask) {
	klog.Info("处理失败: ", task.SubmitId)
	mysql.DB.Model(&model.Submit{}).Where("id = ?", task.SubmitId).
		Update("status", model.StatusFailed).
		Update("message", "处理失败")
}

func judge(task JudgeTask) {
	onStartJudging(task)
	score, err := onProcJudging(task)
	if err != nil {
		onFailed(task)
		return
	}

	onCompleted(task, score)
}
