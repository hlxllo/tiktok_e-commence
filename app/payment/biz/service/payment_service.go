package service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	model2 "tiktok_e-commence/app/order/biz/model"
	"tiktok_e-commence/app/payment/biz/model"
	"tiktok_e-commence/common"
)

type PaymentServer struct {
	model.UnimplementedPaymentServiceServer
}

// 实现 Charge
func (s *PaymentServer) Charge(c context.Context, req *model.ChargeReq) (*model.ChargeResp, error) {
	// 先调用订单服务，查询订单是否存在
	userId := req.UserId
	instances, err := common.SelectHealthyInstance("order-server")
	if err != nil {
		return nil, status.Errorf(codes.Internal, common.ErrDiscoverServiceFailed)
	}
	conn, err := common.CreateGRPCConn(instances.Ip, int(instances.Port))
	defer conn.Close()
	orderClient := model2.NewOrderServiceClient(conn)
	orderResp, _ := orderClient.ListOrder(c, &model2.ListOrderReq{UserId: userId})
	orders := orderResp.Orders
	if len(orders) == 0 {
		return nil, status.Errorf(codes.NotFound, common.ErrOrderNotFound)
	}
	// 存在则继续
	po := &model.PaymentPo{}
	po.Amount = req.Amount
	po.OrderId = req.OrderId
	po.UserId = req.UserId
	card, _ := json.Marshal(req.CreditCard)
	po.CreditCard = card
	id := model.CreatePayment(po)
	return &model.ChargeResp{TransactionId: strconv.Itoa(int(id))}, nil
}
