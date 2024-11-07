package service

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"tiktok_e-commence/app/product/biz/model"
)

// 实现服务端接口
type ProductServer struct {
	model.UnimplementedProductCatalogServiceServer
}

// 实现 ListProducts
func (s *ProductServer) ListProducts(c context.Context, req *model.ListProductsReq) (*model.ListProductsResp, error) {
	productPos := model.SelectProductByCat(req.CategoryName, int(req.Page), int(req.PageSize))
	// 映射为返回类型
	var products []*model.Product
	//var categories pq.StringArray
	for _, productPo := range productPos {
		product := &model.Product{}
		copier.Copy(product, productPo)
		// 反序列化 Categories 字段
		var categories []string
		json.Unmarshal(productPo.Categories, &categories)
		product.Categories = categories
		products = append(products, product)
	}
	return &model.ListProductsResp{Products: products}, nil
}

//// 实现 GetProduct
//func (s *ProductServer) GetProduct(c context.Context, req *model.GetProductReq) (*model.GetProductResp, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
//}
//
//// 实现 SearchProducts
//func (s *ProductServer) SearchProducts(c context.Context, req *model.SearchProductsReq) (*model.SearchProductsResp, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
//}
