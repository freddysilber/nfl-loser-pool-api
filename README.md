# NFL Looser Pool REST API

REST API for NFL Looser Pool game

# Getting Started

* Install Postgres, Docker, and Go. (The link above should provide that info and if not, please raise an issue)

* You'll need to set up a ```.env``` file to store your PostgreSQL information. Here is an example:

```
POSTGRES_USER=my_postgres_user
POSTGRES_PASSWORD=my_postgres_users_password
POSTGRES_DB=my_postgres_database_name
```

* To dig into Postgresql, you can use ```psql```

* run ```docker-compose up --build``` to launch the api server on port ```8080```

* run ```docker-compose down``` to destroy Docker containers

## References and Tools

* [Tutorial for reference](https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/)

* To troubleshoot initial postgres setup ```sudo pkill -u postgres```
  * I ran this when I got a 'Port 5432 is already in use' error??

* [user does not exist](https://stackoverflow.com/questions/17633422/psql-fatal-database-user-does-not-exist)

## Notes for me

* postgres, postgres, nfl_looser_pool_api_db