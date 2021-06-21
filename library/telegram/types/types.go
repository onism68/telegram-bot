package types

type TgMsg struct {
	ChatId  int64
	Message string
	Type    string
	ImgList []string
}
