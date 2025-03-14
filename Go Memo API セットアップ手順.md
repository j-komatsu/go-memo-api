# Go Memo API セットアップ手順

## 1. 環境構築

### 必要なツール
- Go 1.16 以上
- Git
- SQLite
- GCC（Windows の場合のみ必要）

## 2. リポジトリのクローン
```sh
git clone https://github.com/your-repo/go-memo-api.git
cd go-memo-api
```

## 3. 依存関係のインストール
```sh
go mod tidy
```

## 4. データベースのセットアップ
SQLite を使用するため、データベースファイルを作成します。
```sh
sqlite3 notes.db "CREATE TABLE notes (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, content TEXT NOT NULL, created_at DATETIME DEFAULT CURRENT_TIMESTAMP);"
```

## 5. アプリケーションの起動
```sh
go run main.go
```

## 6. API の動作確認
ターミナルで `curl` コマンドを使用して API をテストできます。

### メモ一覧取得
```sh
curl -X GET http://localhost:8080/notes
```

### メモの追加
```sh
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d '{"title":"テストメモ", "content":"これはテストです"}'
```

## 7. フロントエンドの確認
ブラウザで以下の URL にアクセスし、画面の表示を確認してください。
```
http://localhost:8080/index.html
```

---

## トラブルシューティング

### 1. `cgo: C compiler \"gcc\" not found` エラー
Go-SQLite3 を使用する場合、GCC が必要です。
#### Windows での GCC のインストール手順
1. [MSYS2](https://www.msys2.org/) をインストール
2. MSYS2 を起動し、以下のコマンドを実行
   ```sh
   pacman -S mingw-w64-x86_64-gcc
   ```
3. 環境変数を設定（永続化）
   ```sh
   setx PATH "%PATH%;C:\msys64\mingw64\bin"
   ```
4. ターミナルを再起動し、`gcc --version` でインストール確認
   ```sh
   gcc --version
   ```

### 2. `CORS policy: No 'Access-Control-Allow-Origin' header is present` エラー
API へのアクセス制限が原因です。`main.go` に CORS 設定を追加してください。
```go
r.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    c.Next()
})
```

この修正後、サーバーを再起動してください。

### 3. `ERR_CONNECTION_REFUSED`
サーバーが正しく起動していない可能性があります。
1. `go run main.go` が正常に実行されているか確認
2. `localhost:8080` がリッスン状態になっているか確認
3. ターミナルで `curl -X GET http://localhost:8080/notes` を試す
4. ファイアウォールやセキュリティソフトがブロックしていないか確認

