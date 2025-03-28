# 📘 GitHub運用ルール

## 1. ブランチ運用ルール

| ブランチ名 | 用途 |
|------------|------|
| `main`     | 常にデプロイ可能な状態。本番運用を想定。 |
| `dev`      | 開発用統合ブランチ。動作確認済みコードを統合。 |
| `feature/*` | 機能追加用の個別ブランチ（例：`feature/login`）。 |
| `hotfix/*` | 緊急修正時のブランチ。 |

マージの流れ： `feature/*` → `dev` → `main`

---

## 2. Pull Requestルール

- PRは`dev`または`main`へのマージ前に必ず作成する
- PRのタイトルは、Conventional Commits規約に準拠する
- PR本文には以下を含める：
  - 概要（何をしたか）
  - 目的（なぜやるのか）
  - 動作確認内容やスクリーンショット
  - チェックリスト（Lint済・テスト済など）
- `.github/PULL_REQUEST_TEMPLATE.md` を使用

---

## 3. コミットメッセージルール（Conventional Commits）

```text
feat: 新機能の追加
fix: バグ修正
docs: ドキュメントの変更
style: フォーマットのみの変更
refactor: 挙動変更なしのリファクタリング
test: テストコードの追加・修正
chore: ビルドツールやCI構成の変更
```

---

## 4. CI/CDルール

- GitHub Actionsを使用
- ブランチ`main`またはPR作成時に自動実行
- 以下を実行：
  - Lintチェック（`go vet` / `golangci-lint`）
  - 単体テスト（`go test ./...`）

READMEにCIバッジを掲載する：
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


