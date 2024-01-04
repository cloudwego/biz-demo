// Code generated by Kitex v0.8.0. DO NOT EDIT.

package cartservice

import (
	"context"
	cart "github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/cart"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return cartServiceServiceInfo
}

var cartServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CartService"
	handlerType := (*cart.CartService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AddItem":   kitex.NewMethodInfo(addItemHandler, newAddItemArgs, newAddItemResult, false),
		"GetCart":   kitex.NewMethodInfo(getCartHandler, newGetCartArgs, newGetCartResult, false),
		"EmptyCart": kitex.NewMethodInfo(emptyCartHandler, newEmptyCartArgs, newEmptyCartResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "cart",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func addItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.AddItemRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).AddItem(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AddItemArgs:
		success, err := handler.(cart.CartService).AddItem(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddItemResult)
		realResult.Success = success
	}
	return nil
}
func newAddItemArgs() interface{} {
	return &AddItemArgs{}
}

func newAddItemResult() interface{} {
	return &AddItemResult{}
}

type AddItemArgs struct {
	Req *cart.AddItemRequest
}

func (p *AddItemArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.AddItemRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddItemArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddItemArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddItemArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddItemArgs) Unmarshal(in []byte) error {
	msg := new(cart.AddItemRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddItemArgs_Req_DEFAULT *cart.AddItemRequest

func (p *AddItemArgs) GetReq() *cart.AddItemRequest {
	if !p.IsSetReq() {
		return AddItemArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddItemArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddItemArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddItemResult struct {
	Success *cart.Empty
}

var AddItemResult_Success_DEFAULT *cart.Empty

func (p *AddItemResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddItemResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddItemResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddItemResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddItemResult) Unmarshal(in []byte) error {
	msg := new(cart.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddItemResult) GetSuccess() *cart.Empty {
	if !p.IsSetSuccess() {
		return AddItemResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddItemResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.Empty)
}

func (p *AddItemResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddItemResult) GetResult() interface{} {
	return p.Success
}

func getCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.GetCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).GetCart(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCartArgs:
		success, err := handler.(cart.CartService).GetCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCartResult)
		realResult.Success = success
	}
	return nil
}
func newGetCartArgs() interface{} {
	return &GetCartArgs{}
}

func newGetCartResult() interface{} {
	return &GetCartResult{}
}

type GetCartArgs struct {
	Req *cart.GetCartRequest
}

func (p *GetCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.GetCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.GetCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCartArgs_Req_DEFAULT *cart.GetCartRequest

func (p *GetCartArgs) GetReq() *cart.GetCartRequest {
	if !p.IsSetReq() {
		return GetCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCartResult struct {
	Success *cart.Cart
}

var GetCartResult_Success_DEFAULT *cart.Cart

func (p *GetCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.Cart)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCartResult) Unmarshal(in []byte) error {
	msg := new(cart.Cart)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCartResult) GetSuccess() *cart.Cart {
	if !p.IsSetSuccess() {
		return GetCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.Cart)
}

func (p *GetCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCartResult) GetResult() interface{} {
	return p.Success
}

func emptyCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.EmptyCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).EmptyCart(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *EmptyCartArgs:
		success, err := handler.(cart.CartService).EmptyCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EmptyCartResult)
		realResult.Success = success
	}
	return nil
}
func newEmptyCartArgs() interface{} {
	return &EmptyCartArgs{}
}

func newEmptyCartResult() interface{} {
	return &EmptyCartResult{}
}

type EmptyCartArgs struct {
	Req *cart.EmptyCartRequest
}

func (p *EmptyCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.EmptyCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EmptyCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EmptyCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EmptyCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *EmptyCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.EmptyCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EmptyCartArgs_Req_DEFAULT *cart.EmptyCartRequest

func (p *EmptyCartArgs) GetReq() *cart.EmptyCartRequest {
	if !p.IsSetReq() {
		return EmptyCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EmptyCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EmptyCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EmptyCartResult struct {
	Success *cart.Empty
}

var EmptyCartResult_Success_DEFAULT *cart.Empty

func (p *EmptyCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EmptyCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EmptyCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EmptyCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *EmptyCartResult) Unmarshal(in []byte) error {
	msg := new(cart.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EmptyCartResult) GetSuccess() *cart.Empty {
	if !p.IsSetSuccess() {
		return EmptyCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EmptyCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.Empty)
}

func (p *EmptyCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EmptyCartResult) GetResult() interface{} {
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

func (p *kClient) AddItem(ctx context.Context, Req *cart.AddItemRequest) (r *cart.Empty, err error) {
	var _args AddItemArgs
	_args.Req = Req
	var _result AddItemResult
	if err = p.c.Call(ctx, "AddItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCart(ctx context.Context, Req *cart.GetCartRequest) (r *cart.Cart, err error) {
	var _args GetCartArgs
	_args.Req = Req
	var _result GetCartResult
	if err = p.c.Call(ctx, "GetCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EmptyCart(ctx context.Context, Req *cart.EmptyCartRequest) (r *cart.Empty, err error) {
	var _args EmptyCartArgs
	_args.Req = Req
	var _result EmptyCartResult
	if err = p.c.Call(ctx, "EmptyCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}