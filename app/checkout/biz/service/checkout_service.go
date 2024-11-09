package service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	model2 "tiktok_e-commence/app/checkout/biz/model"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

type CheckoutServer struct {
	model.UnimplementedCheckoutServiceServer
}

// 实现 Checkout
func (s *CheckoutServer) Checkout(c context.Context, req *model.CheckoutReq) (*model.CheckoutResp, error) {
	// 先调用订单服务，查询订单是否存在
	userId := req.UserId
	instances, err := common.SelectHealthyInstance("order-server")
	if err != nil {
		return nil, status.Errorf(codes.Internal, common.ErrDiscoverServiceFailed)
	}
	conn, err := common.CreateGRPCConn(instances.Ip, int(instances.Port))
	defer conn.Close()
	orderClient := model.NewOrderServiceClient(conn)
	orderResp, _ := orderClient.ListOrder(c, &model.ListOrderReq{UserId: userId})
	orders := orderResp.Orders
	if len(orders) == 0 {
		return nil, status.Errorf(codes.NotFound, common.ErrOrderNotFound)
	}
	// 存在，获取订单 id
	resp := &model.CheckoutResp{}
	resp.OrderId = orders[0].OrderId
	// 再保存数据，返回结果
	po := &model2.CheckoutPo{}
	po.UserId = req.UserId
	po.Email = req.Email
	po.Firstname = req.Firstname
	po.Lastname = req.Lastname
	addr, _ := json.Marshal(req.Address)
	card, _ := json.Marshal(req.CreditCard)
	po.Address = addr
	po.CreditCard = card
	// 添加数据
	resp.TransactionId = strconv.Itoa(int(model2.CreateCheckout(po)))
	return resp, nil
}
