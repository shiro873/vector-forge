**Contributing to vector-forge**

Thanks for your interest in contributing! This document explains how to report issues, propose changes, and run the project locally.

- **Report an issue**: Open an issue describing the bug, expected behaviour, and steps to reproduce. Include logs or minimal repro when possible.
- **Feature requests**: Open an issue labeled `enhancement` with a short motivation and proposed API or UX.

Development workflow

- Fork the repository and create a topic branch named `feat/<short-desc>` or `fix/<short-desc>`.
- Keep changes small and focused. Rebase/squash commits when appropriate.

Local setup (examples)

- The project contains several services. For a quick dev stack using Docker Compose:

```bash
cd infra
docker compose up --build
```

- To run the NestJS gateway locally:

```bash
cd gateway
npm install
npm run start:dev
```

- To run the AI embedding service locally (FastAPI/Uvicorn):

```bash
cd ai-embedding
python -m uvicorn main:app --host 0.0.0.0 --port 8001
```

Testing & linting

- Follow each subproject's instructions. Example for the gateway:

```bash
cd gateway
npm run test
npm run lint
```

Guidelines

- **Small PRs**: Aim for small, reviewable pull requests with a clear description and linked issue.
- **Commit messages**: Use concise imperative-style messages (e.g., "Add health check for embedding service").
- **Code style**: Follow existing project style. Run linters and formatters before submitting.
- **Documentation**: Update `README.md`, component READMEs, or inline docs for any behavior or API changes.

Review process

- PRs will be reviewed by maintainers. Address review comments and add tests where possible.

Security

- For security issues, open a private issue or contact the maintainers directly instead of public issues.

Thanks for helping improve `vector-forge`!
