package dto

type PostDto struct {
	Message string `json:"message" form:"message"`
	PicUrl  string `json:"picUrl" form:"picUrl"`
}
