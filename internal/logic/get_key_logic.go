package logic

import (
	"context"

	"st-key-manager/internal/svc"
	"st-key-manager/keymanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKeyLogic {
	return &GetKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetKeyLogic) GetKey(in *keymanager.GetKeyRequest) (*keymanager.GetKeyReply, error) {
	data := l.svcCtx.Redis.Get(in.Key)

	return &keymanager.GetKeyReply{
		Data: data,
	}, nil
}
