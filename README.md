# NFL Looser Pool REST API

REST API for NFL Looser Pool game

# Getting Started

* Install Postgres, Docker, and Go.

* You'll need to set up a ```.env``` file to store your PostgreSQL information. Here is an example:

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
DB_HOST=host.docker.internal

export POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```
> Note: POSTGRESQL_URL is not being used in the program, but we can reference this variable in our db migration commands

*create schemas ```migrate create -ext sql -dir db/migrations -seq create_items_table```

* Run Database Migrations
  * Migrate Up
  ```bash
  migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path db/migrations up
  ```
  or
  ```bash
  export POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  migrate -database ${POSTGRESQL_URL} -path db/migrations up
  ```


  * Migrate Down
  ```bash
  migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path db/migrations down
  ```

* To dig into Postgresql, you can use ```psql```

* run ```docker-compose up --build``` to launch the api server on port ```8080```
  > Before you run this command, be sure your local postgres server is running using the postgres app

* run ```docker-compose down``` to destroy Docker containers

* Create a sample item 
```bash
curl -X POST http://localhost:8080/items -H "Content-type: application/json" -d '{ "name": "swim across the River Benue", "description": "ho ho ho"}'
```

* Get all items
```bash
curl http://localhost:8080/items
```

## References and Tools

* [Tutorial for reference](https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/)

* To troubleshoot initial postgres setup ```sudo pkill -u postgres```
  * I ran this when I got a 'Port 5432 is already in use' error??

* [user does not exist](https://stackoverflow.com/questions/17633422/psql-fatal-database-user-does-not-exist)
