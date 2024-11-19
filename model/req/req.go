package req

type LoginReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
