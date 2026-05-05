# Jinja Journey

Jinja Journey は、神社巡りをもっと自由に楽しむための神社ポータルサービスです。

地域やご利益から神社を探し、由緒・御朱印・アクセス情報などを閲覧できます。

---

## コンセプト

「神社巡りをもっと自由に」

和の静けさや自然の空気感を大切にしながら、
現代的で使いやすいUI/UXを目指しています。

---

## 主な機能

- 神社検索
    - エリア検索
    - ご利益検索
    - フリーワード検索

- 神社一覧表示
    - 人気順
    - 新着順
    - 評価順
    - 地図表示

- 神社詳細ページ
    - 基本情報
    - ご利益
    - 由緒
    - 御朱印
    - アクセスマップ

- イベント情報表示

- お気に入り機能

---

## ご利益カテゴリ

- 縁結び
- 金運
- 健康
- 学業成就
- 厄除け
- 仕事運

---

## デザインコンセプト

### キーワード

- 和
- 静けさ
- 自然
- 余白
- モダン

### カラーパレット

| 用途 | Color |
|---|---|
| Primary | `#D9331D` |
| Sub | `#3A7D44` |
| Accent | `#6BAFAF` |
| Background | `#F9F7F4` |
| Surface | `#FFFFFF` |
| Text | `#333333` |
| Border | `#E5E5E5` |

---

## 画面構成

### TOPページ

- Hero画像
- 神社検索導線
- イベント一覧
- フッター

### 検索モーダル

- エリア検索
- ご利益検索
- キーワード検索

### 一覧ページ

- 検索条件表示
- 並び替え
- 神社カード一覧

### 詳細ページ

- Hero画像
- 基本情報
- ご利益タグ
- 由緒
- 御朱印
- アクセスマップ

---

## 使用技術

### Frontend

- React

- Next.js

- TypeScript

- TailwindCSS

### Backend

- Go

### API

- REST API

### Infrastructure

- Vercel

- AWS（予定）

- Docker Compose（開発環境）

---

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

### 停止

```bash
docker compose down
```

DB のデータも削除する場合:

```bash
docker compose down -v
```

---

## 開発予定

- レスポンシブ対応（PC/SP）
- Google Maps連携
- お気に入り機能
- レビュー機能
- 神社イベント管理
- 御朱印ギャラリー

---

## デザイン方針

SPファーストで設計し、
PCでは余白と一覧性を強化するレイアウトを採用予定。

---

## 開発目的

神社巡りをもっと身近に、
もっと楽しくするサービスを目指しています。
