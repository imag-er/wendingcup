import numpy as np
from pathlib import Path
import os
import json


class Judger:
    def __init__(self, gt_folder, pred_folder, num_classes=3):
        self.gt_folder = gt_folder
        self.pred_folder = pred_folder
        self.num_classes = num_classes
        self.confusion_matrix = np.zeros(
            (num_classes, num_classes), dtype=np.int64)
        self.gt_files = list(Path(gt_folder).glob("*.npy"))
        print(f"judger inited with gt_folder={gt_folder}, pred_folder={pred_folder}, num_classes={num_classes}")

    async def calculate_iou_metrics(self):
        if not self.gt_files:
            raise RuntimeWarning("没有找到ground truth文件,请联系管理员")

        for gt_file in self.gt_files:
            filename = gt_file.stem
            pred_file = Path(self.pred_folder) / f"{filename}.npy"
            if not pred_file.exists():
                raise RuntimeWarning("预测文件不存在,请联系管理员")

            gt = np.load(str(gt_file))
            pred = np.load(str(pred_file))
            if gt.shape != pred.shape:
                raise RuntimeWarning(
                    f"形状不匹配: gt={gt.shape}, pred={pred.shape}, 文件={filename}")

            for i in range(self.num_classes):
                for j in range(self.num_classes):
                    self.confusion_matrix[i,
                                          j] += np.sum((gt == i) & (pred == j))

        iou_per_class = np.zeros(self.num_classes)
        for i in range(self.num_classes):
            intersection = self.confusion_matrix[i, i]
            union = (np.sum(self.confusion_matrix[i, :]) +
                     np.sum(self.confusion_matrix[:, i]) -
                     intersection)
            if union == 0:
                iou_per_class[i] = 0
            else:
                iou_per_class[i] = intersection / union

        # 计算总体mIoU
        mIoU = np.mean(iou_per_class)
        return iou_per_class, mIoU
