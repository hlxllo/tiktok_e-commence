package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/product/biz/model"
)

// 实现服务端接口
type ProductServer struct {
	model.UnimplementedProductCatalogServiceServer
}

// 实现 ListProducts
func (s *ProductServer) ListProducts(c context.Context, req *model.ListProductsReq) (*model.ListProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}

// 实现 GetProduct
func (s *ProductServer) GetProduct(c context.Context, req *model.GetProductReq) (*model.GetProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}

// 实现 SearchProducts
func (s *ProductServer) SearchProducts(c context.Context, req *model.SearchProductsReq) (*model.SearchProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
