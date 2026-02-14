#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
COMPOSE_FILE="${ROOT_DIR}/e2e_tests/docker-compose.e2e.yml"
ENV_FILE="${ROOT_DIR}/.env"

if [[ ! -f "${ENV_FILE}" ]]; then
  echo ".env not found at ${ENV_FILE}" >&2
  exit 1
fi

docker compose --env-file "${ENV_FILE}" -f "${COMPOSE_FILE}" --profile local down -v || true
docker compose --env-file "${ENV_FILE}" -f "${COMPOSE_FILE}" --profile local pull app-local
docker compose --env-file "${ENV_FILE}" -f "${COMPOSE_FILE}" --profile local up -d --wait app-local
