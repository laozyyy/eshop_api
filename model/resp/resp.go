package resp

import "eshop_api/model"

type LoginRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Uid  string `json:"uid,omitempty"`
}

type RegisterRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Uid  string `json:"uid,omitempty"`
}

type UserRespDTO struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	User *model.User `json:"user"`
}
