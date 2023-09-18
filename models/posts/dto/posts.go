package dto

type PostUserDto struct {
	Message string `json:"message" form:"message"`
	PicUrl  string `json:"picUrl" form:"picUrl"`
}
