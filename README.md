## 開発環境の起動

Next.js + Go + PostgreSQL の最小開発環境を Docker Compose で起動できます。

### 必要なもの

- Docker
- Docker Compose

### 起動

```bash
docker compose up --build
```

### アクセス先

| Service | URL |
|---|---|
| Frontend | http://localhost:3000 |
| Backend health | http://localhost:8080/health |
| PostgreSQL | localhost:5432 |

### PostgreSQL 接続情報

| Item | Value |
|---|---|
| Host | localhost |
| Port | 5432 |
| Database | shrine_portal |
| User | postgres |
| Password | postgres |

### テーブル定義

- Google Spreadsheet
  - https://docs.google.com/spreadsheets/d/1zmamhX1EgMS88ZcgPnJRGa5e1uKqOhpu244GNVQApoE/edit?gid=1528500678#gid=1528500678

### 停止

```bash
docker compose down
```

DB のデータも削除する場合:

```bash
docker compose down -v
```
