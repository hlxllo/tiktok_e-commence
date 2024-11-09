package service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	model2 "tiktok_e-commence/app/order/biz/model"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

type OrderServer struct {
	model.UnimplementedOrderServiceServer
}

// 实现 PlaceOrder
func (s *OrderServer) PlaceOrder(c context.Context, req *model.PlaceOrderReq) (*model.PlaceOrderResp, error) {
	po := &model2.OrderPo{}
	po.UserId = req.UserId
	po.UserCurrency = req.UserCurrency
	po.Email = req.Email
	// 结构体转 byte 切片
	addr, _ := json.Marshal(req.Address)
	items, _ := json.Marshal(req.OrderItems)
	po.Address = addr
	po.OrderItems = items
	// 创建订单
	id, err := model2.CreateOrder(po)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, common.ErrOrderExists)
	}
	return &model.PlaceOrderResp{Order: &model.OrderResult{OrderId: strconv.Itoa(int(id))}}, nil
}

// 实现 ListOrder
func (s *OrderServer) ListOrder(c context.Context, req *model.ListOrderReq) (*model.ListOrderResp, error) {
	po := &model2.OrderPo{}
	po.UserId = req.UserId
	orderPos := model2.SelectOrders(po)
	// 反序列化
	var orders []*model.Order
	for _, orderPo := range orderPos {
		order := &model.Order{}
		addr := &model.Address{}
		var items []*model.OrderItem
		json.Unmarshal(orderPo.Address, addr)
		json.Unmarshal(orderPo.OrderItems, &items)
		order.OrderId = strconv.Itoa(int(orderPo.ID))
		order.UserId = orderPo.UserId
		order.UserCurrency = orderPo.UserCurrency
		order.Email = orderPo.Email
		order.Address = addr
		order.OrderItems = items
		orders = append(orders, order)
	}
	return &model.ListOrderResp{Orders: orders}, nil
}

// 实现 MarkOrderPaid
func (s *OrderServer) MarkOrderPaid(c context.Context, req *model.MarkOrderPaidReq) (*model.MarkOrderPaidResp, error) {
	po := &model2.OrderPo{}
	orderId, _ := strconv.Atoi(req.OrderId)
	po.ID = uint(orderId)
	po.UserId = req.UserId
	model2.DeleteOrder(po)
	return &model.MarkOrderPaidResp{}, nil
}
