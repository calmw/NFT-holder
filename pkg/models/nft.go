package models

type NftData struct {
	Id      int32  `gorm:"column:id;primaryKey"`
	Nft     string `json:"nft" gorm:"column:nft"`
	NftFrom string `json:"nft_from" gorm:"column:nft_from"`
	NftTo   string `json:"nft_to" gorm:"column:nft_to"`
	TokenId string `json:"token_id" gorm:"column:token_id"`
	Height  int64  `json:"height" gorm:"column:height"`
	Holder  int    `json:"holder" gorm:"column:holder"`
	LogId   string `json:"log_id" gorm:"column:log_id"`
	TxHash  string `json:"tx_hash" gorm:"column:tx_hash"`
}
