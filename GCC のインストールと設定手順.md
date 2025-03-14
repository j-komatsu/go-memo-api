# GCC のインストールと設定手順

## なぜ GCC をインストールする必要があったのか？
Go で `gorm.io/driver/sqlite` を使用する際、`github.com/mattn/go-sqlite3` というパッケージが内部的に使われます。この `go-sqlite3` は **C 言語で書かれているため、Go のコンパイラが C 言語のコードをビルドできる環境 (GCC) を必要とします**。

`CGO_ENABLED=1` にして `go run` を実行すると、C コンパイラ (`gcc`) を要求されるため、GCC をインストールする必要がありました。

---

## **GCC のインストール手順**
### **Windows 環境 (MSYS2 を使用)**

1. **MSYS2 をインストール**
   - Windows ターミナルで以下のコマンドを実行
   ```sh
   winget install -e --id MSYS2.MSYS2
   ```
   - インストール後、MSYS2 を起動する。

2. **パッケージマネージャー (`pacman`) を更新**
   ```sh
   pacman -Syu
   ```
   - 途中で一度 MSYS2 を閉じる必要がある場合がある。
   - 再度 MSYS2 を開き、以下を実行。
   ```sh
   pacman -Syu
   ```

3. **GCC をインストール**
   ```sh
   pacman -S mingw-w64-x86_64-gcc
   ```

4. **GCC が正しくインストールされたか確認**
   ```sh
   gcc --version
   ```
   - `gcc.exe (Rev3, Built by MSYS2 project) X.X.X` という出力があれば成功。

5. **GCC のパスを Windows の環境変数に追加**
   - `C:\msys64\mingw64\bin` を **環境変数 PATH に追加**。
   - これにより、ターミナルを閉じても `gcc` が使えるようになる。

---

## **CGO を有効にして Go アプリを実行する**
GCC をインストールした後、Go の `CGO_ENABLED=1` を設定してアプリを実行する。

```sh
set CGO_ENABLED=1
```

または、一時的に有効化した状態で `go run` する場合:

```sh
CGO_ENABLED=1 go run main.go
```

---

## **補足**
- **GCC は C 言語のコンパイラ** であり、Go の `go-sqlite3` のような **C 言語のコードを含むパッケージを使う際に必要**。
- **MSYS2 は Windows 上で Linux 風の環境を提供するツール** であり、`pacman` でパッケージを管理できる。
- **環境変数を正しく設定しないと、ターミナルを閉じると `gcc` が認識されなくなる**。

---

### **まとめ**
✅ `GCC` をインストールしないと `go-sqlite3` がビルドできない。
✅ `MSYS2` を使って `GCC` をインストールし、`pacman -S mingw-w64-x86_64-gcc` でセットアップ。
✅ `CGO_ENABLED=1` を設定し、Go のビルドを正しく実行。

---

これで `go-sqlite3` を含む Go プロジェクトを正しく実行できるようになります！
