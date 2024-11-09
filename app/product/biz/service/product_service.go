package service

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model2 "tiktok_e-commence/app/product/biz/model"
	"tiktok_e-commence/common/model/model"
)

// 实现服务端接口
type ProductServer struct {
	model.UnimplementedProductCatalogServiceServer
}

// 实现 ListProducts
func (s *ProductServer) ListProducts(c context.Context, req *model.ListProductsReq) (*model.ListProductsResp, error) {
	pos := model2.SelectProductByCat(req.CategoryName, int(req.Page), int(req.PageSize))
	// 映射为返回类型
	var products []*model.Product
	for _, po := range pos {
		product := &model.Product{}
		// 拷贝
		copier.Copy(product, po)
		// 反序列化 Categories 字段
		var categories []string
		json.Unmarshal(po.Categories, &categories)
		product.Categories = categories
		products = append(products, product)
	}
	return &model.ListProductsResp{Products: products}, nil
}

// 实现 GetProduct
func (s *ProductServer) GetProduct(c context.Context, req *model.GetProductReq) (*model.GetProductResp, error) {
	// 构造查询参数
	po := &model2.ProductPo{}
	po.ID = uint(req.Id)
	// 查询数据，取第一个
	pos := model2.SelectProducts(po)
	if len(pos) > 0 {
		po = pos[0]
		product := &model.Product{}
		copier.Copy(product, po)
		// 反序列化 Categories 字段
		var categories []string
		json.Unmarshal(po.Categories, &categories)
		product.Categories = categories
		return &model.GetProductResp{Product: product}, nil
	}
	return &model.GetProductResp{}, nil
}

// 实现 SearchProducts TODO 不知道啥意思，先不写
func (s *ProductServer) SearchProducts(c context.Context, req *model.SearchProductsReq) (*model.SearchProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
