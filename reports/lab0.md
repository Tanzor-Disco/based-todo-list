# Лабораторная работа 0

## Тема работы

Разработка full-stack веб-приложения для управления задачами с использованием Go, React и PostgreSQL.

## Цель работы

Разработать веб-приложение, позволяющее пользователям регистрироваться, выполнять вход в систему и работать со своим списком задач.

В рамках работы необходимо реализовать клиентскую и серверную части приложения, организовать взаимодействие между ними посредством HTTP API, обеспечить хранение данных в PostgreSQL и реализовать механизм пользовательских сессий.

## Используемые технологии

### Backend

- Go
- `net/http`
- PostgreSQL
- `pgx`
- bcrypt
- SQL migrations
- HTTP cookies
- `embed`

### Frontend

- React
- TypeScript
- Vite
- CSS

## Структура проекта

```text
based-todo-list/
├── backend/
│   ├── auth/
│   │   ├── auth.go
│   │   ├── cookie.go
│   │   └── session.go
│   │
│   ├── db/
│   │   ├── db.go
│   │   ├── migrations/
│   │   │   ├── 001_init.down.sql
│   │   │   └── 001_init.up.sql
│   │   ├── tasks.go
│   │   ├── user_sessions.go
│   │   └── users.go
│   │
│   ├── internal/
│   │   └── apperrors/
│   │       └── server.go
│   │
│   ├── models/
│   │   ├── task.go
│   │   ├── user.go
│   │   └── user_session.go
│   │
│   ├── server/
│   │   ├── login.go
│   │   ├── main.go
│   │   ├── main_task_create.go
│   │   ├── register.go
│   │   ├── server.go
│   │   ├── static.go
│   │   └── static/
│   │       ├── assets/
│   │       ├── favicon.png
│   │       └── index.html
│   │
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
└── frontend/
    ├── src/
    ├── public/
    ├── dist/
    ├── eslint.config.js
    ├── index.html
    ├── package.json
    ├── tsconfig.json
    ├── tsconfig.app.json
    ├── tsconfig.node.json
    └── vite.config.ts
```

## Архитектура приложения

Приложение состоит из трёх основных компонентов:

1. React frontend;
2. Go backend;
3. PostgreSQL database.

Frontend отвечает за пользовательский интерфейс и отправку HTTP-запросов.

Backend реализует HTTP API, авторизацию пользователей и взаимодействие с базой данных.

Общая схема взаимодействия:

```text
┌──────────────┐
│    React     │
│   Frontend   │
└──────┬───────┘
       │
       │ HTTP / JSON
       ▼
┌──────────────┐
│ Go Backend   │
│  net/http    │
└──────┬───────┘
       │
       │ SQL
       ▼
┌──────────────┐
│  PostgreSQL  │
└──────────────┘
```

Во время разработки запросы frontend к `/api` проксируются Vite на Go backend.

```text
React
  │
  │ /api/*
  ▼
Vite
  │
  ▼
Go backend
  │
  ▼
PostgreSQL
```

## Проектирование базы данных

Для хранения данных используется PostgreSQL.

В базе предусмотрены следующие сущности:

- пользователи;
- задачи;
- пользовательские сессии.

Основными таблицами являются:

```text
users
tasks
user_sessions
```

Связи между таблицами:

```text
              ┌──────────────┐
              │    users     │
              └──────┬───────┘
                     │
              ┌──────┴───────┐
              │              │
              ▼              ▼
       ┌─────────────┐ ┌───────────────┐
       │    tasks    │ │ user_sessions │
       └─────────────┘ └───────────────┘
```

Каждая задача содержит `user_id`, который связывает её с конкретным пользователем.

Это позволяет получать только задачи текущего пользователя.

Структура базы данных создаётся с помощью SQL-миграций.

Миграции расположены в:

```text
backend/db/migrations/
```

## Подключение к PostgreSQL

Для взаимодействия с PostgreSQL используется библиотека `pgx`.

Подключение создаётся через connection pool. Использование пула позволяет повторно использовать существующие соединения с базой данных.

URI подключения передаётся через переменную окружения.

Пример:

```env
DB_URI=postgres://postgres@localhost/todo
```

Таким образом, параметры подключения не хранятся непосредственно в исходном коде.

## Регистрация пользователя

Для регистрации frontend отправляет HTTP POST-запрос:

```http
POST /api/register
```

Данные пользователя передаются в формате JSON.

Пример:

```json
{
    "username": "user",
    "password": "password"
}
```

На backend тело запроса декодируется с помощью `json.Decoder`:

```go
err := json.NewDecoder(r.Body).Decode(&user)
```

Перед сохранением пользователя пароль хэшируется с помощью bcrypt:

```go
hash, err := bcrypt.GenerateFromPassword(
    []byte(password),
    bcrypt.DefaultCost,
)
```

В базу данных сохраняется хэш пароля, а исходный пароль не сохраняется.

## Авторизация пользователя

Для входа пользователь отправляет свои учётные данные на backend.

Backend получает сохранённый хэш пароля и сравнивает его с введённым паролем:

```go
err := bcrypt.CompareHashAndPassword(
    []byte(passwordHash),
    []byte(password),
)
```

Метод возвращает `nil`, если пароль соответствует сохранённому хэшу.

После успешной проверки создаётся пользовательская сессия.

## Пользовательские сессии

Для идентификации авторизованного пользователя используется session ID.

Сессия сохраняется в базе данных в таблице:

```text
user_sessions
```

Сессия связана с пользователем через `user_id`.

После успешной авторизации session ID передаётся клиенту через HTTP cookie.

Пример:

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session_id",
    Value:    sessionID,
    Path:     "/",
    HttpOnly: true,
    Secure:   false,
    SameSite: http.SameSiteLaxMode,
})
```

Использование `HttpOnly` запрещает JavaScript напрямую получать значение cookie.

При последующих HTTP-запросах браузер отправляет cookie серверу, благодаря чему backend может определить текущего пользователя.

## HTTP API

Взаимодействие frontend и backend осуществляется через HTTP API.

Для передачи данных используется JSON.

Основные операции приложения:

```text
POST /api/register
POST /api/login
GET  /api/main/tasks
POST /api/main/task/create
```

Регистрация и авторизация используют POST-запросы, поскольку пользовательские данные передаются в теле запроса.

Получение задач выполняется через GET-запрос.

Создание новой задачи выполняется через POST-запрос.

## Работа с задачами

При создании задачи frontend отправляет запрос:

```http
POST /api/main/task/create
```

Данные новой задачи передаются в формате JSON:

```json
{
    "description": "Learn React"
}
```

На backend выполняется SQL-запрос:

```sql
INSERT INTO tasks(user_id, description)
VALUES ($1, $2)
RETURNING id
```

Параметризованный запрос позволяет передавать пользовательские данные отдельно от SQL-кода.

После создания задачи backend возвращает её идентификатор.

## Получение задач

Для получения задач используется:

```http
GET /api/main/tasks
```

Backend получает задачи текущего пользователя из PostgreSQL.

Полученные строки преобразуются в структуры Go, после чего данные сериализуются в JSON и отправляются frontend.

Общая схема:

```text
GET /api/main/tasks
      │
      ▼
Go handler
      │
      ▼
PostgreSQL
      │
      ▼
[]Task
      │
      ▼
JSON response
      │
      ▼
React
```

## Frontend

Frontend реализован с использованием React и TypeScript.

Список задач хранится в состоянии компонента:

```tsx
const [tasks, setTasks] = useState<TaskData[]>([])
```

Для описания структуры задачи используется TypeScript-интерфейс:

```ts
interface TaskData {
    TaskID: number;
    Description: string;
}
```

Типизация позволяет контролировать структуру данных во время разработки.

## Получение данных через `fetch`

После загрузки страницы frontend запрашивает существующие задачи.

Для выполнения запроса используется `useEffect`:

```tsx
useEffect(() => {
    async function fetchTasks() {
        const response = await fetch("/api/tasks")
        const tasks = await response.json()

        setTasks(tasks)
    }

    fetchTasks()
}, [])
```

Асинхронная функция объявляется внутри `useEffect`, поскольку callback самого `useEffect` не должен возвращать Promise.

## Создание задачи на frontend

Для создания задачи используется HTML-форма.

После отправки формы данные извлекаются из `FormData`:

```tsx
const formDataObj = Object.fromEntries(formData)
```

После этого данные преобразуются в JSON и отправляются на backend:

```tsx
const response = await fetch("/api/main/task/create", {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify(formDataObj),
})
```

После успешного создания новая задача добавляется в существующий массив:

```tsx
setTasks(prevTasks => [...prevTasks, newTask])
```

Таким образом, интерфейс обновляется без полной перезагрузки страницы.

## Маршрутизация

Для клиентской маршрутизации используется React Router.

В приложении предусмотрены страницы:

- регистрация;
- авторизация;
- главная страница с задачами;

После успешной регистрации пользователь может перейти на страницу авторизации.

После успешной авторизации пользователь переходит к списку задач.

Клиентская маршрутизация позволяет переключать страницы без полной перезагрузки документа.

## Vite

Для разработки и сборки frontend используется Vite.

Установка зависимостей:

```bash
npm install
```

Запуск development-сервера:

```bash
npm run dev
```

Production-сборка:

```bash
npm run build
```

Во время разработки Vite используется для проксирования API-запросов к Go backend.

```text
/api/*
    ↓
http://localhost:8080
```

Благодаря этому frontend может использовать относительные пути:

```ts
fetch("/api/tasks")
```

## Раздача frontend через Go

Production-версия frontend используется в директории:

```text
backend/server/static/
```

Статические файлы встраиваются в Go-приложение с помощью `embed.FS`:

```go
//go:embed static
var staticFS embed.FS
```

Это позволяет Go-серверу самостоятельно отдавать HTML, JavaScript, CSS и другие статические ресурсы.

Для существующих файлов из `static` сервер возвращает соответствующий файл.

Для маршрутов React Router, которым не соответствует отдельный статический файл, backend возвращает `index.html`.

Это позволяет React Router самостоятельно обработать URL на стороне клиента.

## Возникшие проблемы

### MIME type при раздаче frontend

При первоначальной реализации раздачи статических файлов backend при запросе JavaScript-файла мог возвращать `index.html`.

Например, браузер запрашивал:

```text
/assets/index-BMP86Wxa.js
```

но вместо JavaScript получал HTML.

В результате браузер сообщал об ошибке:

```text
blocked because of a disallowed MIME type ("text/html")
```

Причина заключалась в том, что backend не находил запрошенный статический файл и переходил к fallback на `index.html`.

Проблема была решена исправлением обработки путей внутри `embed.FS`, после чего существующие JavaScript и CSS-файлы начали отдаваться непосредственно, а `index.html` используется для frontend-маршрутов.

### Несовпадение типов TypeScript

При добавлении новой задачи возникла ошибка TypeScript из-за несовпадения названий полей объекта.

Интерфейс ожидал поле:

```ts
TaskID
```

а создаваемый объект содержал:

```ts
ID
```

В результате объект не соответствовал типу `TaskData`.

Проблема была решена приведением структуры объекта к единому интерфейсу.

## Конфигурация

Для подключения к PostgreSQL используется файл `.env`, расположенный в директории `backend`.

Пример:

```env
DB_URI=postgres://postgres@localhost/todo
```

## Запуск приложения

### Установка frontend-зависимостей

```bash
cd frontend
npm install
```

### Сборка frontend

Данный пункт можно пропустить, так как уже есть закоммиченная папка static в которой лежит готовая сборка frontend
Вообще, адекватнее всего было бы прописать в vite config, чтобы при сборке файлы попадали сразу в backend/server/staic, но в процессеработы компилировать прошлось лишь единожды, поэтому не было смысла возиться с конфигом vite

```bash
npm run build
```

### Запуск backend

Backend запускается из директории `backend`:

```bash
cd ../backend
go run .
```

После запуска сервер доступен по адресу:

```text
http://localhost:8080
```

## Результат работы

В результате было разработано full-stack веб-приложение для управления задачами.

В приложении реализованы:

- регистрация пользователей;
- авторизация;
- хэширование паролей с помощью bcrypt;
- пользовательские сессии;
- хранение session ID в базе данных;
- HTTP cookies;
- создание задач;
- получение задач пользователя;
- PostgreSQL;
- connection pool через `pgx`;
- HTTP API на Go;
- React frontend;
- TypeScript;
- React Router;
- Vite;
- Vite proxy;
- SQL-миграции;
- раздача production frontend через Go;
- embedding статических файлов с помощью `embed.FS`.

## Вывод

В ходе выполнения лабораторной работы было разработано full-stack веб-приложение для управления задачами.

Была реализована серверная часть на Go с использованием стандартного пакета `net/http`, библиотеки `pgx` для работы с PostgreSQL и bcrypt для хэширования паролей.

Для авторизации был реализован механизм пользовательских сессий с использованием HTTP cookies. Данные пользователей, задач и сессий сохраняются в PostgreSQL.

Клиентская часть была разработана на React и TypeScript. Для управления состоянием использовался React state, получение данных выполнялось с помощью `fetch` и `useEffect`, а навигация между страницами реализована с помощью React Router.

Также была настроена интеграция Vite и Go backend, production-сборка React-приложения и его раздача непосредственно Go-сервером с использованием `embed.FS`.

В результате были получены практические навыки разработки и интеграции frontend, backend и базы данных в рамках одного приложения
