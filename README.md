# GoSelfLearn
My First golang mvc sample with simple blog post concept

# Requirements
- Docker (v20 or later)
- Docker Compose (v2 or later)
- Go (v1.19.x or later)

## How to run
### 1.Database
The project requires a postgres database. So, first run docker-compose file to make postgres database up and running:
```bash
make run_compose
```

### 2.Environment variables
The project use `.env` files. Copy `app.env.sample` to `app.env` and fill the field with proper values

### 3.Run in dev mode
Run this command in project root directory
```bash
make dev
```
This command checked that `air` is installed or not and if not exists, install it for you and then run the project by the property you defined in `app.env`

## How to build
In order to build project for your target host, run below command:
```bash
make build GOOS=linux GOARCH=amd64
```
