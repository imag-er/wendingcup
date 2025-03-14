// Code generated by Kitex v0.9.1. DO NOT EDIT.

package board

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetBoard": kitex.NewMethodInfo(
		getBoardHandler,
		newGetBoardArgs,
		newGetBoardResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AppendJudgeResult": kitex.NewMethodInfo(
		appendJudgeResultHandler,
		newAppendJudgeResultArgs,
		newAppendJudgeResultResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	boardServiceInfo                = NewServiceInfo()
	boardServiceInfoForClient       = NewServiceInfoForClient()
	boardServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return boardServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return boardServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return boardServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Board"
	handlerType := (*board.Board)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "board",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func getBoardHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(board.GetBoardRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(board.Board).GetBoard(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetBoardArgs:
		success, err := handler.(board.Board).GetBoard(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetBoardResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetBoardArgs() interface{} {
	return &GetBoardArgs{}
}

func newGetBoardResult() interface{} {
	return &GetBoardResult{}
}

type GetBoardArgs struct {
	Req *board.GetBoardRequest
}

func (p *GetBoardArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(board.GetBoardRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetBoardArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetBoardArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetBoardArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetBoardArgs) Unmarshal(in []byte) error {
	msg := new(board.GetBoardRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetBoardArgs_Req_DEFAULT *board.GetBoardRequest

func (p *GetBoardArgs) GetReq() *board.GetBoardRequest {
	if !p.IsSetReq() {
		return GetBoardArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBoardArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetBoardArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetBoardResult struct {
	Success *board.GetBoardResponse
}

var GetBoardResult_Success_DEFAULT *board.GetBoardResponse

func (p *GetBoardResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(board.GetBoardResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetBoardResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetBoardResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetBoardResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetBoardResult) Unmarshal(in []byte) error {
	msg := new(board.GetBoardResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBoardResult) GetSuccess() *board.GetBoardResponse {
	if !p.IsSetSuccess() {
		return GetBoardResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBoardResult) SetSuccess(x interface{}) {
	p.Success = x.(*board.GetBoardResponse)
}

func (p *GetBoardResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetBoardResult) GetResult() interface{} {
	return p.Success
}

func appendJudgeResultHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(board.AppendJudgeResultRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(board.Board).AppendJudgeResult(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AppendJudgeResultArgs:
		success, err := handler.(board.Board).AppendJudgeResult(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AppendJudgeResultResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAppendJudgeResultArgs() interface{} {
	return &AppendJudgeResultArgs{}
}

func newAppendJudgeResultResult() interface{} {
	return &AppendJudgeResultResult{}
}

type AppendJudgeResultArgs struct {
	Req *board.AppendJudgeResultRequest
}

func (p *AppendJudgeResultArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(board.AppendJudgeResultRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AppendJudgeResultArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AppendJudgeResultArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AppendJudgeResultArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AppendJudgeResultArgs) Unmarshal(in []byte) error {
	msg := new(board.AppendJudgeResultRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AppendJudgeResultArgs_Req_DEFAULT *board.AppendJudgeResultRequest

func (p *AppendJudgeResultArgs) GetReq() *board.AppendJudgeResultRequest {
	if !p.IsSetReq() {
		return AppendJudgeResultArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AppendJudgeResultArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AppendJudgeResultArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AppendJudgeResultResult struct {
	Success *board.AppendJudgeResultResponse
}

var AppendJudgeResultResult_Success_DEFAULT *board.AppendJudgeResultResponse

func (p *AppendJudgeResultResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(board.AppendJudgeResultResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AppendJudgeResultResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AppendJudgeResultResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AppendJudgeResultResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AppendJudgeResultResult) Unmarshal(in []byte) error {
	msg := new(board.AppendJudgeResultResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AppendJudgeResultResult) GetSuccess() *board.AppendJudgeResultResponse {
	if !p.IsSetSuccess() {
		return AppendJudgeResultResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AppendJudgeResultResult) SetSuccess(x interface{}) {
	p.Success = x.(*board.AppendJudgeResultResponse)
}

func (p *AppendJudgeResultResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AppendJudgeResultResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetBoard(ctx context.Context, Req *board.GetBoardRequest) (r *board.GetBoardResponse, err error) {
	var _args GetBoardArgs
	_args.Req = Req
	var _result GetBoardResult
	if err = p.c.Call(ctx, "GetBoard", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AppendJudgeResult(ctx context.Context, Req *board.AppendJudgeResultRequest) (r *board.AppendJudgeResultResponse, err error) {
	var _args AppendJudgeResultArgs
	_args.Req = Req
	var _result AppendJudgeResultResult
	if err = p.c.Call(ctx, "AppendJudgeResult", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
