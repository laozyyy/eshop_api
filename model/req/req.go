package req

type LoginReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type MGetSkuReqDTO struct {
	TagID    string `json:"tag_id" binding:"required"`
	PageNum  int    `json:"page_num" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
}
