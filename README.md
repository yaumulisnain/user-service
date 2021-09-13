# USER SERVICE

Simple user service. This service is covers user-registration and user-login. The database only support for PostgreSQL.

## Contents
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Quick Start](#quick-start)

## Tech stack
* Go, with modules:
    * [gin](github.com/gin-gonic/gin)
    * [godotenv](github.com/joho/godotenv)
    * [gorm](github.com/jinzhu/gorm)
    * [validator](github.com/go-playground/validator/v10)
* PostgreSQL

## Prerequisites
* Go v1.15
* PostgreSQL

## Setup
1. Clone this project
```
git clone https://github.com/yaumulisnain/user-service
```

2. You must create ```.env``` file, you can copy from ```.env.dev``` and then customize the value.
```
APP_ENV=dev
TZ=UTC
PORT=:8081

DB_HOST=userdbserver
DB_DATABASE=userdb
DB_USERNAME=username
DB_PASSWORD=password
DB_PORT=5432
```

3. Running with Makefile command

## Quick Start
This project using Makefile for simplify your commands.

1. Run project
    ```
    $ make run
    ```

2. Build project
    ```
    $ make build
    ```

3. Run binary file
    ```
    $ make start
    ```

4. Clean up Docker
    ```
    $ make clean
    ```