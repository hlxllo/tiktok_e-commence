package service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"tiktok_e-commence/app/order/biz/model"
)

type OrderServer struct {
	model.UnimplementedOrderServiceServer
}

// 实现 PlaceOrder
func (s *OrderServer) PlaceOrder(c context.Context, req *model.PlaceOrderReq) (*model.PlaceOrderResp, error) {
	po := &model.OrderPo{}
	po.UserId = req.UserId
	po.UserCurrency = req.UserCurrency
	po.Email = req.Email
	// 结构体转 byte 切片
	addr, _ := json.Marshal(req.Address)
	items, _ := json.Marshal(req.OrderItems)
	po.Address = addr
	po.OrderItems = items
	// 创建订单
	id := model.CreateOrder(po)
	return &model.PlaceOrderResp{Order: &model.OrderResult{OrderId: strconv.Itoa(int(id))}}, nil
}

// 实现 ListOrder
func (s *OrderServer) ListOrder(c context.Context, req *model.ListOrderReq) (*model.ListOrderResp, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}

// 实现 MarkOrderPaid
func (s *OrderServer) MarkOrderPaid(c context.Context, req *model.MarkOrderPaidReq) (*model.MarkOrderPaidResp, error) {

	return nil, status.Errorf(codes.Unimplemented, "method MarkOrderPaid not implemented")
}
