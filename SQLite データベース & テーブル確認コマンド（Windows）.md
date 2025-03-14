# SQLite データベース & テーブル確認コマンド（Windows）

## **1. SQLite の対話モードを開く**
```sh
sqlite3 notes.db
```
**成功すると** `sqlite>` というプロンプトが表示される。

## **2. データベースの確認**
### **2.1 データベース一覧の確認**
```sh
.databases
```
**期待する出力例**
```
main: C:/path/to/notes.db
```

### **2.2 データベースファイルが作成されているか確認**
```sh
dir notes.db
```
**期待する出力例（ファイルがある場合）**
```
2025/03/12  10:00 AM            3,072 notes.db
```

## **3. テーブルの確認**
### **3.1 作成されたテーブルの一覧を表示**
```sh
.tables
```
**期待する出力例**
```
notes
```

### **3.2 テーブルの構造を確認**
```sh
PRAGMA table_info(notes);
```
**期待する出力例**
```
0|id|INTEGER|1||1
1|title|TEXT|1||0
2|content|TEXT|1||0
3|created_at|DATETIME|0||0
```

### **3.3 テーブルのスキーマ（定義）を確認**
```sh
.schema notes
```
**期待する出力例**
```sql
CREATE TABLE notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### **3.4 テーブルにデータが入っているか確認**
```sh
SELECT * FROM notes;
```
**期待する出力例（まだデータがない場合）**
```
id|title|content|created_at
```

## **4. テーブルが存在しない場合の確認**
```sh
SELECT name FROM sqlite_master WHERE type='table';
```
**期待する出力例（`notes` テーブルがある場合）**
```
notes
```

## **5. SQLite の対話モードを終了する**
### **方法 1: `.exit` コマンド**
```sh
.exit
```

### **方法 2: `.quit` コマンド**
```sh
.quit
```

### **方法 3: `Ctrl + Z`（Windows のみ）**
キーボードで **`Ctrl` + `Z`** を押してから **`Enter`** を押すと終了。

---

## **まとめ**
1. `sqlite3 notes.db` でDBに接続
2. `.databases` でデータベース一覧を確認
3. `.tables` でテーブル一覧を確認
4. `PRAGMA table_info(notes);` でテーブル構造を確認
5. `SELECT * FROM notes;` でデータを確認
6. `.schema notes` でスキーマを確認
7. `.exit` または `.quit` で終了

これで **Windows 環境で SQLite のデータベースとテーブルの作成が正しくできているか** をチェックできます！ ✅
