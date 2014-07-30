package controllers

import (
  // "revel_redis/app/models"
	"fmt"
	"github.com/astaxie/goredis"
	r "github.com/revel/revel"
)

var (
	client goredis.Client
)

func init() {
	client.Addr = "127.0.0.1:6379"
}

type ModelController struct {
	*r.Controller
}

func (c *ModelController) Begin() r.Result {
  fmt.Println("###########################################################")
	fmt.Println("DB HERE---->%+v", client)
	return nil
}
