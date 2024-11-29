package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"nft-data/pkg/db"
	"nft-data/pkg/models"
	"strconv"
	"strings"
	"time"
)

type Nft struct {
	ID       string `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	TokenID  string `json:"tokenId"`
	Ctime    string `json:"ctime"`
	Height   string `json:"height"`
	UtcTtime string `json:"utcTtime"`
	TxHash   string `json:"txHash"`
}

type NftData struct {
	Data struct {
		TransferLogs []Nft `json:"transferLogs"`
	} `json:"data"`
}

func SaveDataFromGraph() bool {
	nftAddress := strings.ToLower("0x29d0C4a595A05632864F6aA02C80A37cC9b4623A")
	url := "https://subgraph.intoverse.co/subgraphs/name/nft3a"
	method := "POST"

	index := 0
	for {
		query := fmt.Sprintf(`{"query": "query myQuery { transferLogs(first: 500, skip:%d orderBy: ctime,  orderDirection: asc) { id  from to tokenId ctime utcTtime height txHash utcTtime  txHash  } }"}`,
			index*500)
		log.Println(query)
		payload := strings.NewReader(query)
		client := &http.Client{Timeout: time.Second * 30}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		req.Header.Add("Content-Type", "application/json")
		var res *http.Response
		for {
			res, err = client.Do(req)
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Println(err.Error())
			return false
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		var rechargeRecord NftData
		err = json.Unmarshal(body, &rechargeRecord)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		fmt.Println(len(rechargeRecord.Data.TransferLogs), 12312)
		if len(rechargeRecord.Data.TransferLogs) <= 0 {
			log.Println("没数据了", index)
			return false
		}
		for _, transferLog := range rechargeRecord.Data.TransferLogs {
			SaveData(transferLog, nftAddress)
		}

		index++
	}

}
func SaveDataFromGraph2() bool {
	nftAddress := strings.ToLower("0x1483F8DC6CFAcbf68C3A3F6D64Ca3Bd33e666491")
	url := "https://subgraph.intoverse.co/subgraphs/name/nft91"
	method := "POST"

	index := 0
	for {
		query := fmt.Sprintf(`{"query": "query myQuery { transferLogs(first: 500, skip:%d orderBy: ctime,  orderDirection: asc) { id  from to tokenId ctime height utcTtime txHash utcTtime  txHash  } }"}`,
			index*500)
		log.Println(query)
		payload := strings.NewReader(query)
		client := &http.Client{Timeout: time.Second * 30}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		req.Header.Add("Content-Type", "application/json")
		var res *http.Response
		for {
			res, err = client.Do(req)
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Println(err.Error())
			return false
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		var rechargeRecord NftData
		err = json.Unmarshal(body, &rechargeRecord)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		fmt.Println(len(rechargeRecord.Data.TransferLogs), 12312)
		if len(rechargeRecord.Data.TransferLogs) <= 0 {
			log.Println("没数据了", index)
			return false
		}
		for _, transferLog := range rechargeRecord.Data.TransferLogs {
			SaveData(transferLog, nftAddress)
		}

		index++
	}

}

func SaveData(nft Nft, nftAddress string) {
	var nftData models.NftData
	lodId := fmt.Sprintf("%s-%s", nft.ID, strings.ToLower(nftAddress))
	err := db.Mysql.Model(&models.NftData{}).Where("log_id = ?", lodId).First(&nftData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h, _ := strconv.ParseInt(nft.Height, 10, 64)
		err = db.Mysql.Model(&models.NftData{}).Create(&models.NftData{
			LogId:   lodId,
			NftFrom: strings.ToLower(nft.From),
			NftTo:   strings.ToLower(nft.To),
			TokenId: nft.TokenID,
			Height:  h,
			Holder:  0,
			Nft:     strings.ToLower(nftAddress),
			TxHash:  strings.ToLower(nft.TxHash),
		}).Error
		if err != nil {
			log.Println("save data error", err)
			return
		}
	}
}
