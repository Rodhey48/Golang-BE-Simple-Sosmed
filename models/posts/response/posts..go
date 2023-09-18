package response

type PostResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	PicUrl  string `json:"picUrl"`
}
