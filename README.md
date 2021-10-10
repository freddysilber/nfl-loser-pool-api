# nfl-looser-pool-api
REST API for NFL Looser Pool game

[Tutorial I build this from](https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/)

## Tools to Troubleshoot...
* To troubleshoot initial postgres setup ```sudo pkill -u postgres```
  * I ran this when I got a 'Port 5432 is already in use' error??
* https://stackoverflow.com/questions/17633422/psql-fatal-database-user-does-not-exist

# Getting Started

* You'll need a ```.env``` file with ```POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB``` values
* To start the server, run ```docker-compose up --build``` or ```npm run dev```
* To dig into Postgresql, you can use ```psql```

# TODO

* Debug the DB connection... ```go run main.go``` wont even complete