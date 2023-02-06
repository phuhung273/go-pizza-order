## How to deploy

### Build Docker-Compose with .env file

```bash
docker-compose -f <path_file_docker-compose> --env-file <path_env_file> up -d
```

```bash
docker-compose -f go-pizza-order-docker-compose --env-file local.env up -d
```

### Yourfile.env exmaple <local.env>

```
SERVICE_PORT=8080
SERVICE_MODE=debug
DB_HOST=root
DB_PASSWORD=1234
#must same as container in docker compose file
DB_IP=mysqldb-container
DB_PORT=5506
DB_NAME=go-pizza-order-db

#This variable must be declare after those variables DB_HOST, DB_PASSWORD, DB_IP, DB_PORT, DB_NAME
#If u declare this variable before those variables it will not recognize those variable like this :
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_HOST\" variable is not set. Defaulting to a blank string."
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_PASSWORD\" variable is not set. Defaulting to a blank string."
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_IP\" variable is not set. Defaulting to a blank string."
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_PORT\" variable is not set. Defaulting to a blank string."
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_NAME\" variable is not set. Defaulting to a blank string."
#time="2023-01-14T11:24:04+07:00" level=warning msg="The \"DB_HOST\" variable is not set. Defaulting to a blank string."


CONNECTION_STRING_DB="${DB_HOST}:${DB_PASSWORD}@tcp(${DB_IP}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local"
```
