package file

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/imag-er/wendingcup/app/submit/biz/dal/model"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/submit/conf"
	"gorm.io/gorm"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	folderInCompressedFile = "infer_results_npy"
)

type fileManager struct {
	SubmitFilePath string
	expressedDir   string
}

var FileManager *fileManager

func InitFileManager() {
	moduleRoot, err := os.Getwd()
	if err != nil {
		return
	}
	SubmitFilePath := filepath.Join(moduleRoot, "submit_files")
	FileManager = &fileManager{
		SubmitFilePath: SubmitFilePath,                                        // 存放提交文件的文件夹
		expressedDir:   filepath.Join(SubmitFilePath, folderInCompressedFile), // 解压后的文件夹
	}

	for range conf.GetConf().Judge.NumThreads {
		go judgeWorker()
	}
	klog.Info("fileManager init with value: ", FileManager)

}

func (f *fileManager) GetFilePathById(fileId string) string {
	filePath := filepath.Join(f.SubmitFilePath, fileId+".zip")
	return filePath
}

// 通过teamId查询到Id字段, 删除submit_files文件夹下的{Id}.zip文件
func (f *fileManager) RemoveOldFiles(teamId string) {
	var val model.Submit

	// err := mysql.DB.Where("team_id = ?", teamId).First(&val).Error
	// 改为查询最新一次记录
	err := mysql.DB.Where("team_id = ?", teamId).Order("id desc").First(&val).Error
	if err != nil  {
		if err == gorm.ErrRecordNotFound {
			klog.Info("第一次提交: ", teamId)
			return
		}
		klog.Error(err)
		return
	}
	klog.Info("查询到submit记录 ", val.ID)
	fileName := filepath.Join(f.SubmitFilePath, strconv.FormatUint(uint64(val.ID), 10)+".zip")
	cmd := exec.Command("rm", "-f", fileName)
	_, err = cmd.CombinedOutput()
	if err != nil {
		klog.Error("Error removing file:", err)
		return
	}
	klog.Info("删除文件" + fileName + "成功")

}

// 排到队了再解压
func (f *fileManager) Extract(fileId string) {

	cmd := exec.Command("rm", "-rf", f.expressedDir)
	_, err := cmd.CombinedOutput()
	if err != nil {
		klog.Error("Error removing file:", err)
		return
	}
	klog.Info("删除文件夹成功")

	fileName := filepath.Join(f.SubmitFilePath, fileId+".zip")

	klog.Infof("expdir: %s   file: %s", f.expressedDir, fileName)

	// 压缩包内是一个infer_results_npy文件夹

	cmd = exec.Command("unzip", fileName, "-d", f.SubmitFilePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		// 如果解压过程中发生错误，输出错误信息
		klog.Error("Error unzipping file:", err)
		return
	}

	klog.Info("解压成功")
}
