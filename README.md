# Заготовка fake/stub-хранилища и тестов

fakeStorage, errorStub и CacheService иллюстрируют тестовые двойники. NewCacheService принимает конкретный fakeStorage, но один вызов передаёт *errorStub. Функции Test... и bench... находятся в tests.go: стандартный запуск тестов их не обнаруживает.

Учебная папка `OOP/tests`. Код сохранён в исходном состоянии.

## Структура

- `.gitignore`
- `go.mod`
- `tests.go`

## Требования и команды

Проверено на Windows с Go 1.27.1 (windows/amd64); `go.mod` требует Go 1.27.1. Команды выполняются из корня репозитория. Для самостоятельной публикации добавлен go.mod; исходные .go не изменялись.

Исполняемой точки входа нет. Это пакет или незавершённая заготовка; `go run .` для него не применяется.

```sh
go test ./...
go vet ./...
```

## Результаты проверки

- `go test ./...`: ошибка (код 1).
- `go vet ./...`: ошибка (код 1).

Файлов `*_test.go` нет. Успех `go test` здесь означает проверку сборки, а не прохождение набора тестов.

Диагностика `go test ./...` (фактический результат):

```text
FAIL	github.com/K0uzer/oop-tests [build failed]
FAIL
# github.com/K0uzer/oop-tests
.\tests.go:47:25: cannot use store (variable of type *errorStub) as fakeStorage value in argument to NewCacheService
```

Проверки выполнены 22 ноября 2026 года. Сборка и тесты не исправляют существующие учебные ошибки.
