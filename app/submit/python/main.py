# uvicorn main:app --reload --host "0.0.0.0" --port 13271
import judge
import fastapi
from fastapi.middleware.cors import CORSMiddleware

gt_folder = './ground_truth'
pred_folder = './infer_results_npy'  # 这里要通过volume挂到app/submit/submit_files
num_classes = 3
judger = judge.Judger(gt_folder, pred_folder, num_classes)

app = fastapi.FastAPI()
# allow cors
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/judge")
async def judge_endpoint() -> dict:
    try:
        iou_per_class, miou = await judger.calculate_iou_metrics()
        print(iou_per_class, miou)
        return {
            "iou_per_class": iou_per_class.tolist(),
            "miou": miou.tolist(),
            "error": None
        }
    except Exception as e:
        return {
            "error": str(e),
            "iou_per_class": None,
            "miou": None
        }
