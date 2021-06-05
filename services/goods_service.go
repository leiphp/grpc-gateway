package services

import (
	"context"
	"fmt"
)

type GoodsService struct {

}

//最终实现GetGoodsStock
func (this *GoodsService) GetGoodsStock(ctx context.Context, in *GoodsRequest) (*GoodsResponse, error)  {
	fmt.Println("GoodsRequest:",in)
	fmt.Println("GoodsID:",in.GoodsId)
	return &GoodsResponse{GoodsStock: in.GoodsId*10},nil
}

