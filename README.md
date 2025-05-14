<p align="center">
  <img  width="200" src="https://github.com/user-attachments/assets/2992c99d-0201-42c5-ab16-5ff73073b929" alt="Centered Image"/>
  <h1 align="center">Liftoff</h1>
</p>

<p align="center">
Welcome to <b>Liftoff</b> 🚀 Your go‑to CLI for multi‑region Cloud Run canary deployments! This README will guide you through installation, configuration, and usage with plenty of examples and emojis to keep things fun 🎉.
  <br/>
  <!-- <a href="https://flynnfc.dev/work/bagginsdb">
    Learn more on how it's made here
  </a> -->
</p>

<p align="center">
  <img src="https://github.com/framequery/liftoff/actions/workflows/release.yaml/badge.svg" alt="Build badge">
  <a href="https://github.com/framequery/liftoff/blob/main/LICENSE.md">
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" title="MIT License" />
  </a>
</p>
  

### 📥 Pre-requisites

This CLI heavily relies on the GCLOUD cli tool. You must have it installed.
If you don't, follow [this guide](https://cloud.google.com/sdk/docs/install)

### 🛠️ Installation

Make sure you have Go 1.20+ installed and your `GOPATH` configured. Then:

```bash
# Install the CLI
go install github.com/framequery/liftoff/cmd/liftoff@latest

# Verify installation
liftoff --help
```

You can also build from source:

```bash
git clone https://github.com/fromequery/liftoff.git
cd liftoff
go build -o liftoff ./cmd/liftoff
```

---

### ⚙️ Configuration

By default, **liftoff** saves your settings to `$TMP/liftoff_config.json`. You can override options via flags or environment variables.

| Option         | Env Var              | Description                                              | Default                  |
| -------------- | -------------------- | -------------------------------------------------------- | ------------------------ |
| `--project`    | `LIFTOFF_PROJECT`    | GCP project ID                                           | **none** (required)      |
| `--service`    | `LIFTOFF_SERVICE`    | Cloud Run service name                                   | **none** (required)      |
| `--image`      | `LIFTOFF_IMAGE`      | Container image URL for the canary revision              | **none** (required)      |
| `--regions`    | `LIFTOFF_REGIONS`    | Comma‑separated GCP regions (e.g. `europe-west2,europe-west4`) | `europe-west2,europe-west4` |
| `--percentages`| `LIFTOFF_PCTS`       | Traffic percentages (e.g. `10,50,100`)                   | `10,50,100`              |
| `--intervals`  | `LIFTOFF_INTVLS`     | Seconds between rollout steps (e.g. `300,300`)           | `300,300`                |
| `--env-vars`   | `LIFTOFF_ENV_VARS`   | Comma-separated KEY=VALUE pairs                          | **none** (required)      |

---

Once you have **liftoff** installed, here are some example workflows:

### 🎯 Default canary rollout

```bash
liftoff canary \
  --project=my-gcp-project \
  --service=my-service \
  --image=gcr.io/my-gcp-project/my-app:canary
```
- Deploys no-traffic revisions in `europe-west2` & `europe-west4`.
- Routes 10% → wait 5m → 50% → wait 5m → 100%.

### 🌐 Custom regions & speed

```bash
liftoff canary \
  --project=my-gcp-project \
  --service=api-service \
  --image=gcr.io/my-gcp-project/api:v2 \
  --regions=us-central1,asia-northeast1 \
  --percentages=5,25,50,100 \
  --intervals=60,120,180
```
- Targets `us-central1` and `asia-northeast1`.
- Gradually shifts traffic 5% → wait 1m → 25% → wait 2m → 50% → wait 3m → 100%.
### 🌳 Environment variables

```bash
liftoff canary \
  --project=my-gcp-project \
  --service=api-service \
  --image=gcr.io/my-gcp-project/api:v2 \
  --regions=us-central1,asia-northeast1 \
  --percentages=5,25,50,100 \
  --intervals=60,120,180 \
  --env-vars=GOOGLE_PROJECT_ID=canary,DEBUG=true"
```
### ⚡ Instant full rollout

```bash
liftoff canary \
  --project=my-gcp-project \
  --service=static-site \
  --image=gcr.io/my-gcp-project/site:latest \
  --percentages=100 \
  --intervals=0
```
- Skips staging phases and sends 100% traffic immediately.
### ⚙️ Config & Defaults
```bash
# Save defaults for 'internal-api'
liftoff config set internal-api \
  --project=my-gcp-project \
  --image=gcr.io/my-gcp-project/internal-api:canary \
  --regions=europe-west2,europe-west4 \
  --percentages=10,50,100 \
  --intervals=300,300

# View all saved configs
liftoff config view
```
- You can set/view default config settings per service video with the config command
---

### 🔍 Flags & Commands Reference

```
liftoff --help
```

```
liftoff canary --help
```

Key flags for `canary`:
- `--project, -p`   : GCP project ID (required)
- `--service, -s`   : Cloud Run service name (required)
- `--image, -i`     : Container image URL (required)
- `--regions`       : Regions list
- `--percentages`   : Traffic split percentages
- `--intervals`     : Seconds between shifts


### ❤️ Contributing

We ❤️ pull requests!
1. Fork ✅
2. Create a feature branch 🌿
3. Write tests 🧪
4. Send PR 📬

Please follow our [Contributing Guidelines](CONTRIBUTING.md).

---

### 📜 License

[MIT](LICENSE)

Enjoy safe liftoffs! 🚀
Made by [Framequery](https://www.framequery.com/)
