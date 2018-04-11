## SemiRevel
ゼミ資料管理システム

## revelのinstall
```
$ go get github.com/revel/revel
$ go get github.com/revel/cmd/revel
```

## パッケージ管理ツール`dep`の`install`
```
$ go get -u github.com/golang/dep/cmd/dep
```

## 実行コマンド
```
$ dep ensure                       // プロジェクトの依存関係をインストール
$ revel run                        // 起動
```

## Docker-Composeで起動
```
$ docker-compose up
```
