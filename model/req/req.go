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

type PageReqDTO struct {
	PageNum  int `json:"page_num" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}

type AddItemRequestDTO struct {
	SkuId    string `thrift:"skuId,1,required" frugal:"1,required,string" json:"sku_id"`
	Quantity int32  `thrift:"quantity,2,required" frugal:"2,required,i32" json:"quantity"`
	Uid      string `thrift:"uid,3" frugal:"3,default,string" json:"uid"`
}

type SearchReqDTO struct {
	Keyword  string `json:"keyword" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
	PageNum  int    `json:"page_num" binding:"required"`
}

type UpdateItemRequestDTO struct {
	SkuId    string `json:"sku" binding:"required"`
	Quantity int32  `json:"quantity"`
	Selected bool   `json:"selected"`
	Uid      string `json:"uid" binding:"required"`
}

type DeleteItemRequestDTO struct {
	Skus []string `json:"skus" binding:"required"`
	Uid  string   `json:"uid" binding:"required"`
}

type SeckillReqDTO struct {
	UserID     string `json:"user_id" binding:"required"`
	Sku        string `json:"sku" binding:"required"`
	ActivityID string `json:"activity_id" binding:"required"`
}

type LotteryReqDTO struct {
	UserID     string `json:"user_id" binding:"required"`
	ActivityID string `json:"activity_id" binding:"required"`
}
