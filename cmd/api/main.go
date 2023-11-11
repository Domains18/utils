package  main

import (
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"

	orderModel "goshop/internal/order/model"
	productModel "goshop/internal/product/model"
	grpcServer "goshop/internal/server/grpc"
	httpServer "goshop/internal/server/http"
	userModel "goshop/internal/user/model"
	"goshop/pkg/config"
	"goshop/pkg/dbs"
	"goshop/pkg/redis"
)

func main() {
	cfg  := LoadComfig()
}
