FROM golang:1.23-alpine AS builder


WORKDIR /usr/src/wendingcup
ENV GOPROXY=https://goproxy.io,direct

COPY rpc_gen rpc_gen
COPY common common

RUN cd common && go mod download && go mod verify
RUN cd rpc_gen && go mod download && go mod verify

ARG SVC
COPY app/${SVC}/go.mod app/${SVC}/go.sum ./app/${SVC}/

RUN cd app/${SVC} && go mod download 

COPY app/${SVC}/ app/${SVC}/

RUN cd app/${SVC}/ && CGO_ENABLED=0 GOOS=linux go build -o main .


# 第二阶段：运行应用程序
FROM alpine:3.19
ARG SVC
ENV GO_ENV=online
WORKDIR /app

COPY app/${SVC}/conf ./conf
COPY --from=builder /usr/src/wendingcup/app/${SVC}/main .
CMD ["./main"]