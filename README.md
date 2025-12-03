**Project Overview**

- **Name**: `vector-forge` — a small multi-service repo for vector embeddings, inference, and a NestJS gateway.
- **Purpose**: Provide building blocks for an embeddings + inference stack (embedding service, inference service, API gateway, infra definitions and example pipelines).

**Repository Layout**

- **`ai-embedding/`**: Python FastAPI embedding service (stub). Exposes `/embed` and `/health` endpoints (port 8001 in `docker-compose`). See `pyproject.toml` and `main.py`.
- **`gateway/`**: NestJS (TypeScript) API gateway. Typical Nest starter structure with `package.json` scripts for build/test/run (port 3000 in `docker-compose`).
- **`inference-go/`**: Go-based inference service (Go module present). Built to run as a separate service (port 8081 in `docker-compose`).
- **`infra/`**: Development `docker-compose.yml` that defines Postgres, Redis, gateway, ai-embedding, and inference-go services.
- **`ai-reasoning/`, `document-pipeline/`, `examples/`, `ops/`, `vector-db-migrations/`**: placeholder directories for additional reasoning components, pipelines, examples, operational scripts and DB migrations.

**Quick Start (development)**

- **Requirements**: Docker & Docker Compose, Node.js (for gateway dev), Go (for `inference-go`), Python (for `ai-embedding`).
- Launch the full stack with Docker Compose:

```bash
cd /home/ridwan-cs/Desktop/Other Projects/vector-forge/infra
docker compose up --build
```

- Gateway (local dev):

```bash
cd gateway
npm install
npm run start:dev
```

- AI Embedding (local dev):

```bash
# uses FastAPI; run via Uvicorn or your preferred runner
cd ai-embedding
python -m uvicorn main:app --host 0.0.0.0 --port 8001
```

**Ports (defaults in `docker-compose`)**

- `gateway`: `3000`
- `ai-embedding`: `8001`
- `inference-go`: `8081`
- `postgres`: `5432`
- `redis`: `6379`

**Notes & Next Steps**

- The `ai-embedding` service is a simple stub returning a fixed-dimension vector — replace with a real embedding model or connector.
- `ai-reasoning`, `document-pipeline`, and `vector-db-migrations` are scaffolds; add implementations and README files per component.
- Add health checks, CI workflows, and configuration (secrets management, env files) before production use.

**Maintainer / How to Contribute**

- See `CONTRIBUTING.md` for contribution guidelines.
- For local development questions, open an issue or contact the repo owner.

---

Generated summary: scanned top-level manifests and service skeletons to populate this README.
