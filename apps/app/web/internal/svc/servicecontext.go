package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero/apps/app/web/internal/config"
	"go_zero/apps/app/web/internal/middleware"
	"go_zero/apps/order/rpc/orderclient"
	"go_zero/apps/product/rcp/productclient"
)

type ServiceContext struct {
	Config config.Config
	// 在这里注册BFF中的 RPC 服务
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product

	// 中间件
	login rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		login:      middleware.NewLoginMiddleware().Handle,
	}
}
