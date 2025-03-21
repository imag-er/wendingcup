FROM python:3.12-slim
WORKDIR /judger


RUN pip install numpy uvicorn tqdm

COPY main.py judge.py /judger/

EXPOSE 8381
# ENV NAME World

CMD ["uvicorn","main:app","--reload","--host","0.0.0.0","--port","8381"]
