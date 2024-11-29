package main

import (
	"nft-data/data"
	"nft-data/pkg/db"
)

func main() {
	db.InitMysql()

	data.SaveDataFromGraph()
	data.Parse("0x29d0C4a595A05632864F6aA02C80A37cC9b4623A")
	data.SaveDataFromGraph2()
	data.Parse("0x1483f8dc6cfacbf68c3a3f6d64ca3bd33e666491")
}
