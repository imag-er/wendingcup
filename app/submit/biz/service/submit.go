package service

import (
	"context"
	"os"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/file"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/model"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

type SubmitService struct {
	ctx context.Context
} // NewSubmitService new SubmitService
func NewSubmitService(ctx context.Context) *SubmitService {
	return &SubmitService{ctx: ctx}
}

// Run create note info
func (s *SubmitService) Run(req *submit.SubmitRequest) (resp *submit.SubmitResponse, err error) {

	klog.Infof("submit by: %v", req.TeamId)

	// 删除当前队伍的旧文件
	file.FileManager.RemoveOldFiles(req.TeamId)

	// 读取form内的文件
	fileContent := req.File

	// 创建提交记录
	var val model.Submit
	val = model.Submit{
		TeamId:  req.TeamId,
		Status:  model.StatusUploaded,
		Message: "已上传",
	}
	err = mysql.DB.Create(&val).Error
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	ID_str := strconv.FormatUint(uint64(val.ID), 10)
	// 存入submit_file文件夹
	filePath := file.FileManager.GetFilePathById(ID_str)
	// 写入文件
	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return nil, err
	}

	resp = &submit.SubmitResponse{
		Code:     0,
		Msg:      "上传成功",
		SubmitId: ID_str,
	}

	// 异步判题
	// 排队
	go func() {
		file.TaskQueue <- file.JudgeTask{
			SubmitId: ID_str,
			FilePath: filePath,
			UploadTime: val.CreatedAt,
			TeamId: req.TeamId,
		}	
	}()
	return
}
