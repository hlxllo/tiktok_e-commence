// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cart/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物车服务"
                ],
                "summary": "新增购物车api",
                "parameters": [
                    {
                        "description": "新增的购物车信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddItemReqCopy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "新增成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/cart/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物车服务"
                ],
                "summary": "删除购物车api",
                "parameters": [
                    {
                        "description": "查询的购物车信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EmptyCartReqCopy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/cart/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "购物车服务"
                ],
                "summary": "查询购物车api",
                "parameters": [
                    {
                        "description": "查询的购物车信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetCartReqCopy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/order/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单服务"
                ],
                "summary": "创建订单api",
                "parameters": [
                    {
                        "description": "创建的订单信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PlaceOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/product/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品服务"
                ],
                "summary": "根据 id 查询商品api",
                "parameters": [
                    {
                        "description": "查询的商品 id",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/product/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品服务"
                ],
                "summary": "分页查询商品api",
                "parameters": [
                    {
                        "description": "查询的商品和分页信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListProductsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户服务"
                ],
                "summary": "登录用户api",
                "parameters": [
                    {
                        "description": "登录的用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户服务"
                ],
                "summary": "注册用户api",
                "parameters": [
                    {
                        "description": "新增的用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app_cart_biz_model.CartItem": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "app_order_biz_model.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street_address": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "integer"
                }
            }
        },
        "app_order_biz_model.CartItem": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.AddItemReqCopy": {
            "type": "object",
            "properties": {
                "item": {
                    "$ref": "#/definitions/app_cart_biz_model.CartItem"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.EmptyCartReqCopy": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.GetCartReqCopy": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.GetProductReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.ListProductsReq": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "model.LoginReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.OrderItem": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "item": {
                    "$ref": "#/definitions/app_order_biz_model.CartItem"
                }
            }
        },
        "model.PlaceOrderReq": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/app_order_biz_model.Address"
                },
                "email": {
                    "type": "string"
                },
                "order_items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderItem"
                    }
                },
                "user_currency": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.RegisterReq": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
