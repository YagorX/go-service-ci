### Структура приложения

Проект организован в соответствии с рекомендациями Go и принципами Clean Architecture.

- `cmd/` — точка входа приложения (main.go).
- `internal/` — внутренняя логика приложения:
  - `application/` — инициализация приложения и зависимостей.
  - `config/` — конфигурация приложения.
  - `model/` — доменные модели.
  - `repository/` — слой доступа к данным.
  - `service/` — бизнес-логика.
  - `transport/http/v1/` — HTTP API (контроллеры).
- `test/` — интеграционные тесты и docker-compose окружение.
- `.github/workflows/` — CI/CD пайплайны GitHub Actions.
- `k8s/app/` — Kubernetes манифесты (Deployment и Service).

Такая структура обеспечивает модульность, тестируемость и изоляцию слоёв приложения.

![alt text](image/image.png)

### Локальная сборка и проверка кода

Для локальной разработки используется Makefile, который унифицирует команды:

- `make lint` — запуск golangci-lint.
- `make unit-test` — запуск unit-тестов с покрытием.
- `make build` — сборка бинарного файла приложения.
- `make test-integration` — запуск интеграционных тестов в docker-compose окружении.

Для статического анализа используется golangci-lint с включёнными линтерами:
govet, errcheck, staticcheck.

![alt text](image/image-1.png)


### CI/CD пайплайн

CI/CD реализован с использованием GitHub Actions и описан в `.github/workflows/ci.yml`.

Пайплайн состоит из следующих этапов:

1. **Lint**
   - Проверка кода с помощью golangci-lint.
2. **Unit tests**
   - Запуск unit-тестов с покрытием.
3. **Integration tests**
   - Запуск docker-compose окружения (Kafka, PostgreSQL, Redis).
   - Выполнение интеграционных тестов.
4. **Security**
   - Проверка уязвимостей с помощью gosec и govulncheck.
5. **Build**
   - Сборка Go-приложения.
6. **Build & Push Docker image**
   - Сборка Docker-образа.
   - Публикация образа в GitHub Container Registry (GHCR).

Каждый этап выполняется автоматически при push и pull request в ветки `main` и `master`.
![alt text](image/image-2.png)

В CI/CD используется `gosec` и `govulncheck` для проверки безопасности.

`govulncheck` обнаруживает потенциально достижимую уязвимость в стандартной
библиотеке Go (crypto/tls), которая используется транзитивно через сторонние
библиотеки (Echo, database/sql, Redis).

Приложение не настраивает TLS вручную, а уязвимость относится к реализации
стандартной библиотеки. В связи с этим результат проверки не блокирует
CI/CD пайплайн.

### Docker-образ приложения

Приложение упаковано в Docker-образ с использованием multi-stage сборки:

- Stage `builder`:
  - Сборка Go-бинарника в Alpine-образе.
  - Отключён CGO, используется статическая сборка.
- Runtime stage:
  - Используется distroless образ `gcr.io/distroless/base-debian12`.
  - Минимальный размер и повышенная безопасность.
  - Приложение запускается от non-root пользователя.

Образ публикуется в GitHub Container Registry.
![alt text](image/image-3.png)

### Registry

Docker-образ публикуется в GitHub Container Registry (ghcr.io).

Аутентификация выполняется автоматически в CI/CD с использованием `GITHUB_TOKEN`.
Тег образа соответствует SHA коммита.
![alt text](image/image-4.png)

### Деплой в Kubernetes

Для запуска приложения используется Kubernetes (minikube).

Созданы манифесты:
- Deployment — развёртывание приложения.
- Service — публикация приложения внутри кластера.

Контейнер загружается из GHCR с использованием imagePullSecrets.

![alt text](image/image2.png)

![alt text](image/image2-1.png)

![alt text](image/image2-2.png)

### Проверка работы приложения

После деплоя приложение доступно внутри кластера Kubernetes.
Работа сервиса подтверждена HTTP-запросами.

![alt text](image/image2-3.png)

---

## load/e2e Тестирование и оценка

### Что добавлено

- В CI/CD подключены тестовые этапы:
  - `unit-test` (блокирующий перед упаковкой/публикацией образа)
  - `integration-test`
  - `e2e-test` (ручной запуск через `workflow_dispatch`)
- Добавлен отдельный e2e-контур:
  - `e2e_tests/docker-compose.common.yml`
  - `e2e_tests/docker-compose.e2e.yml`
  - `e2e_tests/Dockerfile_compose`
  - `e2e_tests/tests/*`
- Добавлен health-check endpoint: `GET /health/check`
- Добавлен кастомный HTTP error handler для 404 (HTML/JSON поведение)
- Добавлено нагрузочное тестирование в Apache JMeter:
  - `load_tests/Summary Report.jmx`
  - `load_tests/README.md`

![alt text](image3.png)

### Как запускать тесты

- Unit:
  - `make unit-test`
- Integration:
  - `make test-integration`
- E2E локально:
  - `make test-e2e-docker`
- E2E в CI вручную:
  - `Actions -> Go CI -> Run workflow -> run_e2e=true`
- Нагрузочный тест:
  - см. `load_tests/README.md`

  ![alt text](image3-1.png)

### Оценка результатов нагрузочного теста

По результатам JMeter:

- ошибок: `0.00%`
- throughput: около `503 req/sec`
- задержки: `p95 ~ 4 ms`, `p99 ~ 8 ms`

![alt text](image3-2.png)

Вывод:

- при выбранном профиле нагрузки сервис стабильно обрабатывает запросы без ошибок;
- заметных деградаций по latency на текущем стенде не наблюдается;
- для финальной оценки рекомендуется дополнительно прогнать stress-профиль с большей конкуренцией.

### Протокол испытаний

Текстовый протокол испытаний:

- `load_tests/TEST_PROTOCOL.md`

### Какие скриншоты приложить в PR

1. Успешный прогон CI с этапами `lint`, `unit-test`, `integration-test`, `security`, `build`, `login_push`.
2. Ручной запуск workflow (`Run workflow`) и успешный `E2E tests (manual)`.
3. Результаты JMeter (`Aggregate Report`) с ключевыми метриками.
4. Проверка API:
   - `GET /health/check` -> 200
   - `GET /api/v1/satellite/moon` -> 200
5. (Опционально) `docker compose ps` с поднятыми `app`, `postgres`, `redis`.
