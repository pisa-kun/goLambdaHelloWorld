#### Go言語でのAWS Lambdaの実装サンプル

---

- **go: modules disabled inside GOPATH/src by GO111MODULE=auto** が出た場合
> システム環境変数が優先されるので、GO111MODULE=onをセットしておく

- ビルドまでの流れ
1. go mod init
2. スクリプト記述
3. クロスビルド用の環境変数設定
>cmd # 普段powershellなので、cmdで環境変数設定
>C:\Users\pisa0\go\src\TestLambda>set GOARCH=amd64
>C:\Users\pisa0\go\src\TestLambda>set GOOS=linux

4.  go build -o handler handler.go
```
PS C:\Users\pisa0\go\src\TestLambda>  go build -o handler handler.go
go: finding github.com/aws/aws-lambda-go/lambda latest
go: finding github.com/aws/aws-lambda-go v1.25.0
go: downloading github.com/aws/aws-lambda-go v1.25.0
go: extracting github.com/aws/aws-lambda-go v1.25.0
go: finding github.com/stretchr/testify v1.6.1
go: finding github.com/cpuguy83/go-md2man/v2 v2.0.0
go: finding github.com/urfave/cli/v2 v2.2.0
go: finding gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
go: finding github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d
```
※ linuxビルドするため、生成物にexeが明記されておらず、windowsOSで実行すると対象OSが異なりエラーになるので注意

- awsコンソールでlambda生成

> Go言語におけるハンドラとは、実行可能ファイル名のことです

- テストを実行
詳細 **null**、 ログ出力中に fmt.Println()の出力内容が記述されていればOK

```
START RequestId: b1146149-76f2-4762-ae6f-ecda1431072c Version: $LATEST
Hello GOLANG!!
END RequestId: b1146149-76f2-4762-ae6f-ecda1431072c
REPORT RequestId: b1146149-76f2-4762-ae6f-ecda1431072c	Duration: 0.85 ms	Billed Duration: 1 ms	Memory Size: 512 MB	Max Memory Used: 31 MB	Init Duration: 73.15 ms	
```

[参考](https://www.cview.co.jp/cvcblog/2021.01.15.Pu1pTwLHNUXVe6dnNHrHc)

---

#### おまけ
git push急にできなくなったのでメモ
https://qiita.com/seal_qiita/items/1c5e8621cd22bcd84186