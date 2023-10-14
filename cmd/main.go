package main

import (
	"github.com/hitoshi-w/go-lang/config"
	"github.com/hitoshi-w/go-lang/infrastructure"
)

func main() {
	// 環境変数の設定
	cf, err := config.Initialize()
	if err != nil {
		panic("failed to initialize configuration:" + err.Error())
	}

	// DBへの接続設定
	dbConfig := infrastructure.DBConfig(cf.DB)
	infrastructure.InitializeDB(&dbConfig)

	// ルーティングの設定
	r := infrastructure.InitializeRouter()

	// サーバーの起動
	if err := r.Run(); err != nil {
		panic("failed to start server")
	}
}
