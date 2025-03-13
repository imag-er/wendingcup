package service

import (
	"context"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
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
	// Finish your business logic.
	klog.Infof("submit: %v", req)

	var val model.Submit
	val = model.Submit{
		TeamId:  req.TeamId,
		Status:  model.StatusPending,
		Message: "判题中",
	}
	err = mysql.DB.Create(&val).Error
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	fileContent := req.File
	klog.Info(fileContent)
	// 获取模块根目录
	moduleRoot, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// 构建 submit_files 文件夹路径
	submitFilesDir := filepath.Join(moduleRoot, "submit_files")
	fileName := strconv.FormatUint(uint64(val.ID), 10)
	filePath := filepath.Join(submitFilesDir, fileName)
	// 写入文件
	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return nil, err
	}
	resp = &submit.SubmitResponse{
		SubmitId: fileName,
	}
	return
}
