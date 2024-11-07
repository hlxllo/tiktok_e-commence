package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/cart/biz/model"
	"tiktok_e-commence/common"
)

// 实现服务端接口
type CartServer struct {
	model.UnimplementedCartServiceServer
}

// 实现 AddItem
func (s *CartServer) AddItem(c context.Context, req *model.AddItemReq) (*model.AddItemResp, error) {
	po := &model.CartPo{}
	po.UserId = req.UserId
	po.ProductId = req.Item.ProductId
	po.Quantity = req.Item.Quantity
	// 添加购物车
	_, err := model.CreateCart(po)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, common.ErrCartExists)
	}
	return &model.AddItemResp{}, nil
}

// 实现 GetCart
func (s *CartServer) GetCart(c context.Context, req *model.GetCartReq) (*model.GetCartResp, error) {
	po := &model.CartPo{UserId: req.UserId}
	// 批量查询
	cartPos := model.SelectCarts(po)
	var items []*model.CartItem
	for _, cartPo := range cartPos {
		item := &model.CartItem{}
		item.ProductId = cartPo.ProductId
		item.Quantity = cartPo.Quantity
		items = append(items, item)
	}
	resp := &model.Cart{UserId: req.UserId, Items: items}
	return &model.GetCartResp{Cart: resp}, nil
}

// 实现 EmptyCart
func (s *CartServer) EmptyCart(c context.Context, req *model.EmptyCartReq) (*model.EmptyCartResp, error) {
	po := &model.CartPo{UserId: req.UserId}
	model.DeleteCarts(po)
	return &model.EmptyCartResp{}, nil
}
