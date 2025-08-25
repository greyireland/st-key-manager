package svc

import (
	"context"
	"github.com/greyireland/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"st-key-manager/internal/svc/pricing"
	"sync"
	"time"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)
var auth = NewAuth()

type Auth struct {
	users map[string]*pricing.UserInfo
	lock  sync.Mutex
}

func NewAuth() *Auth {
	a := &Auth{
		users: make(map[string]*pricing.UserInfo),
	}
	go a.run()
	return a
}
func (a *Auth) run() {
	tick := time.NewTicker(time.Second * 10)
	a.loadUsers()
	for {
		select {
		case <-tick.C:
			a.loadUsers()
		}
	}
}

// load users from rpc client
func (a *Auth) loadUsers() {
	list := pricing.GetUserInfo()

	newUsers := make(map[string]*pricing.UserInfo)
	for i := 0; i < len(list); i++ {
		if list[i].Auth == "" || list[i].Role == 0 {
			continue
		}
		newUsers[list[i].Auth] = &list[i]
	}
	a.lock.Lock()
	a.users = newUsers
	a.lock.Unlock()
	log.Debug("Get user info success", "count", len(newUsers))
}
func (a *Auth) IsValidToken(token string) bool {
	if token == "" {
		return false
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	if user, exist := a.users[token]; exist && user.Role > 0 {
		return true
	}
	return false
}

func GetTokenFromCtx(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if len(md["authorization"]) < 1 {
		return ""
	}
	return md["authorization"][0]

}
func EnsureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	token := GetTokenFromCtx(ctx)
	if !auth.IsValidToken(token) {
		log.Warn("invalid token", "token", token)
		return nil, errInvalidToken
	}
	return handler(ctx, req)
}
func EnsureValidTokenStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	token := GetTokenFromCtx(ss.Context())
	if !auth.IsValidToken(token) {
		log.Warn("invalid token", "token", token)
		return errInvalidToken
	}
	return handler(srv, ss)
}
