package handler

import (
	"context"

	usermgr "github.com/Ankr-network/dccn-hub/app_dccn_usermgr/proto/usermgr"
	pb "github.com/Ankr-network/dccn-hub/gateway/proto/usermgr"
)

type UserApi struct {
	client usermgr.UserMgrService
}

func userSrvToApi(user *usermgr.User) *pb.User {
	var u pb.User
	u.Id = user.Id
	u.Name = user.Name
	u.Password = user.Password
	u.IsDeleted = user.IsDeleted
	u.Balance = user.Balance
	u.Email = user.Email
	u.Nickname = user.Nickname
	u.Token = user.Token
	return &u
}

func userApiToSrv(u *pb.User) *usermgr.User {
	var user usermgr.User
	user.Id = u.Id
	user.Name = u.Name
	user.Password = u.Password
	user.IsDeleted = u.IsDeleted
	user.Balance = u.Balance
	user.Email = u.Email
	user.Nickname = u.Nickname
	user.Token = u.Token
	return &user
}

func (p *UserApi) Create(ctx context.Context, user *pb.User, rsp *pb.Response) error {
	_, err := p.client.Create(ctx, userApiToSrv(user))
	return err
}

func (p *UserApi) Login(ctx context.Context, user *pb.LoginRequest, rsp *pb.User) error {
	u, err := p.client.Login(ctx, &usermgr.LoginRequest{Email: user.Email, Password: user.Password})
	rsp = userSrvToApi(u)
	return err
}

func (p *UserApi) Get(ctx context.Context, id *pb.Email, user *pb.User) error {
	response, err := p.client.Get(ctx, &usermgr.Email{Email: id.Email})
	if err != nil {
		return err
	}
	user = userSrvToApi(response)
	return nil
}

func (p *UserApi) GetByEmail(ctx context.Context, email *pb.Email, user *pb.User) error {
	response, err := p.client.Get(ctx, &usermgr.Email{Email: email.Email})
	if err != nil {
		return err
	}
	user = userSrvToApi(response)
	return nil
}

func (p *UserApi) NewToken(ctx context.Context, user *pb.User, tokenString *pb.Token) error {
	if _, err := p.client.NewToken(ctx, userApiToSrv(user)); err != nil {
		return err
	}
	return nil
}

func (p *UserApi) VerifyToken(ctx context.Context, tokenString *pb.Token, rsp *pb.Response) error {
	if _, err := p.client.VerifyToken(ctx, &usermgr.Token{Token: tokenString.Token}); err != nil {
		return err
	}
	return nil
}

func NewUserApi(client usermgr.UserMgrService) *UserApi {
	return &UserApi{client: client}
}
