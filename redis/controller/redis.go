package controller

import (
	"context"
	"net/http"
	redisss "redis-apis/redis"

	"github.com/labstack/echo/v4"
)

type RedisController struct {
	usecase redisss.RedisUsecaseInterface
}

func (r *RedisController) GetData(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authresp, _ := r.usecase.GetData(ctx)
	if authresp == nil {
		return c.JSON(http.StatusBadRequest, authresp)
	}
	return c.JSON(http.StatusOK, authresp)

}

func NewRedisController(e *echo.Echo, redisUsecase redisss.RedisUsecaseInterface) {
	handler := &RedisController{
		usecase: redisUsecase,
	}
	e.POST("/get_data", handler.GetData)
}
