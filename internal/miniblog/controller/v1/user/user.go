package user

import (
	"github.com/DanteSu/miniblog/internal/miniblog/biz"
	"github.com/DanteSu/miniblog/internal/miniblog/store"
	"github.com/DanteSu/miniblog/pkg/auth"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
