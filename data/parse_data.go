package data

import (
	"errors"
	"gorm.io/gorm"
	"nft-data/pkg/db"
	"nft-data/pkg/models"
	"strings"
)

func Parse(nftAddress string) {
	var nfts []models.NftData
	db.Mysql.Model(&models.NftData{}).Where("nft = ?", strings.ToLower(nftAddress)).Order("height asc").Find(&nfts)

	for _, nft := range nfts {
		ParseNft(nft)
	}

}

func ParseNft(nft models.NftData) {
	var n models.NftData
	err := db.Mysql.Model(&models.NftData{}).Where("nft = ? and nft_from = ? and token_id = ? and height > ?", nft.Nft, nft.NftTo, nft.TokenId, nft.Height).First(&n).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Mysql.Model(&models.NftData{}).Where("id=?", nft.Id).Update("holder", 1)
	} else if err == nil {
		ParseNft(n)
	}

}
