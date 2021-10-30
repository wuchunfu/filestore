// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	srvv1 "filestore/internal/apiserver/service/v1"
	"filestore/internal/apiserver/store"
)

// UserController create a user handler used to handler request for user resource.
type UserController struct {
	srv srvv1.Service
}

// NewUserController creates a user handler.
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
