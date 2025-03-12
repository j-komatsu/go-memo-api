# メモアプリ API 設計書

## 1. プロジェクト概要
**プロジェクト名:** Go Memo API  
**目的:** シンプルなメモ管理APIを作成し、CRUD操作を実装することでGoのWeb開発を理解する

## 2. 技術スタック
| 項目 | 技術 |
|---|---|
| 言語 | Go (Golang) |
| フレームワーク | Gin (高速Webフレームワーク) |
| データベース | SQLite (軽量なデータベース) |
| ORM | GORM (GoのORMライブラリ) |
| JSON処理 | `encoding/json` (Go標準ライブラリ) |
| ログ管理 | `log` パッケージ |
| APIテスト | Postman / curl |

## 3. API エンドポイント一覧
| メソッド | エンドポイント | 説明 |
|---|---|---|
| `POST` | `/notes` | メモを新規作成 |
| `GET` | `/notes` | すべてのメモを取得 |
| `GET` | `/notes/{id}` | 指定したメモを取得 |
| `PUT` | `/notes/{id}` | メモを更新 |
| `DELETE` | `/notes/{id}` | メモを削除 |

## 4. データモデル
メモを管理するためのデータモデル。

```go
package models

type Note struct {
    ID        uint   `gorm:"primaryKey"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
```

## 5. API詳細
### `POST /notes`
**リクエストボディ:**
```json
{
  "title": "買い物リスト",
  "content": "牛乳、卵、パン"
}
```

**レスポンス:**
```json
{
  "id": 1,
  "title": "買い物リスト",
  "content": "牛乳、卵、パン",
  "created_at": "2025-03-12T10:00:00Z"
}
```

### `GET /notes`
**レスポンス:**
```json
[
  {
    "id": 1,
    "title": "買い物リスト",
    "content": "牛乳、卵、パン",
    "created_at": "2025-03-12T10:00:00Z"
  },
  {
    "id": 2,
    "title": "アイデアメモ",
    "content": "新しいアプリの設計メモ",
    "created_at": "2025-03-12T11:30:00Z"
  }
]
```

### `PUT /notes/{id}`
**リクエストボディ:**
```json
{
  "title": "仕事のメモ",
  "content": "会議の内容を記録"
}
```

**レスポンス:**
```json
{
  "id": 1,
  "title": "仕事のメモ",
  "content": "会議の内容を記録",
  "created_at": "2025-03-12T10:00:00Z"
}
```

### `DELETE /notes/{id}`
**レスポンス:**
```json
{
  "message": "Note deleted successfully"
}
```

## 6. データベース設定
SQLiteを使用し、データベースの初期化を行う。

```go
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "go-memo-api/models"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
    db.AutoMigrate(&models.Note{})
    return db
}
```

## 7. まとめ
このAPIは、Goの基本的なWeb開発スキルを身につけるために、シンプルながらも実用的な機能を備えています。  
今後、認証機能や検索機能を追加することでさらに拡張できます。
