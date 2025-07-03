# 📘 GitHub 運用ルール

## 1. ブランチ運用ルール

| ブランチ名  | 用途                                              |
| ----------- | ------------------------------------------------- |
| `main`      | 常にデプロイ可能な状態。本番運用を想定。          |
| `dev`       | 開発用統合ブランチ。動作確認済みコードを統合。    |
| `feature/*` | 機能追加用の個別ブランチ（例：`feature/login`）。 |
| `hotfix/*`  | 緊急修正時のブランチ。                            |

マージの流れ： `feature/*` → `dev` → `main`

---

## 2. Pull Request ルール

- PR は`dev`または`main`へのマージ前に必ず作成する
- PR のタイトルは、Conventional Commits 規約に準拠する
- PR 本文には以下を含める：
  - 概要（何をしたか）
  - 目的（なぜやるのか）
  - 動作確認内容やスクリーンショット
  - チェックリスト（Lint 済・テスト済など）
- `.github/PULL_REQUEST_TEMPLATE.md` を使用

---

## 3. コミットメッセージルール（Conventional Commits）

### 🏷️ タイプ一覧

| タイプ   | 説明                         |
| -------- | ---------------------------- |
| `fix`    | バグ修正                     |
| `add`    | 新規（ファイル・機能）の追加 |
| `update` | 機能の修正（バグ以外）       |
| `remove` | 不要なファイルや機能の削除   |

### 📝 記述ルール

- **1 行目**：変更内容の概要（簡潔に）
- **2 行目**：必ず空行
- **3 行目以降**：変更理由・詳細説明（必ず英語）
- 1 コミット 1 目的を意識
- 文末のピリオドは省略可能

---

## 4. CI/CD ルール

- GitHub Actions を使用
- ブランチ`main`または PR 作成時に自動実行
- 以下を実行：
  - Lint チェック（`go vet` / `golangci-lint`）
  - 単体テスト（`go test ./...`）

README に CI バッジを掲載する：

```md
[![CI](https://github.com/hirakan045/feelog/actions/workflows/ci.yml/badge.svg)](https://github.com/hirakan045/feelog/actions/workflows/ci.yml)
```

---

## 5. ドキュメント管理

- 設計資料や構成図は `docs/` ディレクトリに保存
- ファイル構成例：

```text
/docs
├── architecture.md       # システム構成図
├── api-spec.md           # API仕様
├── db-schema.md          # データベース設計（ER図など）
├── reliability-design.md # 可用性・監視設計
├── github-guidlines.md   # GitHub運用ルール
├── roadmap.md            # 今後の開発予定
```

---
