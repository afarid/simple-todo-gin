## Simple todo application using golang gin framework and sqlc


### Prerequisites:
- Install [go](https://golang.org/doc/install)
- Install [docker](https://docs.docker.com/get-docker/)
- Install [docker-compose](https://docs.docker.com/compose/install/)
- Install [sqlc](https://github.com/kyleconroy/sqlc)
- Install [golang migrate](https://github.com/golang-migrate/migrate)

### Setup:
#### Verify that tools are installed and create directories
```bash 
./hack/init.sh
```

#### Create database schema 
- This is an [example](https://dbdiagram.io/d/632784350911f91ba5d9e0bc) of the schema I used for this project
- Export the sqlfile from dbdiagrams.io and save it in the migrations folder
```bash
migrate create --dir db/migrations -ext sql -seq init
cp ~/Dowwnloads/todo.sql db/migrations/000001_init.up.sql
```
- Write the down migration
```sql
echo "DROP TABLE IF EXISTS todo;" > db/migrations/000001_init.down.sql
```
- Run the migrations
```dotenv
make migrate-up
```

#### Generate db access layer code
- Define your queries in `db/query/todo.sql`
- Generate the code
```bash
sqlc generate
```

#### 