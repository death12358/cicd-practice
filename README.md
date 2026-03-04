# Go + Docker Compose CI/CD 練習範例

這是一個用 **Go** 撰寫的簡單 HTTP 服務，搭配 **Dockerfile**、**docker-compose** 與 **GitHub Actions**，讓你可以完整練習從開發、容器化到 CI/CD 的基本流程。

## 專案結構

```text
.
├── cmd
│   └── api
│       └── main.go        # Go HTTP 服務 (含 /health 與 /hello)
├── Dockerfile             # 建置 Go 執行檔的 Docker 映像
├── docker-compose.yml     # 本機用 docker compose 啟動服務
├── go.mod                 # Go module 設定
└── .github
    └── workflows
        └── ci.yml         # GitHub Actions CI/CD workflow
```

## 1. 本機直接跑 Go 服務

```bash
cd cicd
go run ./cmd/api
```

在瀏覽器或 curl 測試：

- `http://localhost:8080/health` 應該會回傳 `OK`
- `http://localhost:8080/hello` 或 `http://localhost:8080/hello?name=AI`

## 2. 使用 Docker 單獨建置與執行

```bash
# 建置映像
docker build -t cicd-go-api .

# 執行容器
docker run --rm -p 8080:8080 --name cicd-go-api cicd-go-api
```

一樣可用瀏覽器或 curl 測試 `localhost:8080/health`。

## 3. 使用 docker compose 啟動

```bash
docker compose up --build
```

預設會：

- 依照 `Dockerfile` 建置映像
- 啟動 `api` 服務並對外開放 `8080` 連接埠

停止：

```bash
docker compose down
```

## 4. CI/CD（GitHub Actions）說明

在 `.github/workflows/ci.yml` 中，定義了以下流程：

- **事件觸發**：`push` / `pull_request` 到 `main` 分支
- **build-and-test job**：
  - Checkout 原始碼
  - 安裝 Go 1.22
  - 執行 `go test ./...`
  - 執行 `go build ./cmd/api`
  - 使用 `docker/build-push-action` 嘗試建置 Docker 映像（不推送）
- **deploy-simulated job**：
  - 依賴 `build-and-test`
  - 只在 `main` 分支上觸發
  - 示範「部署階段」概念（實務上會在這裡 SSH 到伺服器，執行 `docker compose pull && docker compose up -d` 等指令）

### 如何使用這個 CI/CD 範例

1. 在 GitHub 上建立一個新的 repository（例如 `cicd-practice`）。
2. 把這個專案內容整個複製進去（包含 `.github/workflows/ci.yml`）。
3. 推上 GitHub：

```bash
git init
git add .
git commit -m "Add Go + Docker CI/CD example"
git branch -M main
git remote add origin <你的 repo URL>
git push -u origin main
```

4. 到 GitHub repository 的 **Actions** 頁面，就可以看到 CI/CD workflow 的執行情況。

## 5. 可以延伸練習的方向

- 在 `main.go` 新增更多 API (例如 `/time`, `/echo`)
- 加入 `*_test.go` 單元測試，觀察 CI 測試失敗 / 通過的狀況
- 在 `ci.yml` 加上：
  - `golangci-lint` 做靜態程式碼檢查
  - 把 Docker 映像推送到 Docker Hub 或 GitHub Container Registry
- 實作真正的部署流程（例如：
  - 使用 SSH Action 登入你自己的 VPS
  - 在 VPS 上執行 `docker compose pull && docker compose up -d`
)

你可以先照 README 步驟跑一遍，如果想要改成 GitLab CI、Jenkins 或別的平台，我也可以幫你把這個範例轉成對應的設定檔。

