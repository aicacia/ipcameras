# IPCameras

## Getting Started

- go >= 1.21.x [brew](https://formulae.brew.sh/formula/go) or golang [official](https://go.dev/doc/install)
- [go-task](https://taskfile.dev/installation/)
- node [nvm](https://github.com/nvm-sh/nvm)
- cargo [rustup](https://rustup.rs/) is recommended
- install api dev deps `task install`

## Up and Running

- start api in test, open [swagger-ui](https://petstore.swagger.io/?url=http://localhost:3000/openapi.json)
  - `task api`
- start ui in test
  - `task ui`

## Testing

- create local test db
  - `APP_ENV=test task db-up`
  - `APP_ENV=test task db-create`
  - `APP_ENV=test task db-migrate-up`
- start api in test
  - `APP_ENV=test task api`
- start ui in test
  - `APP_ENV=test task ui`
