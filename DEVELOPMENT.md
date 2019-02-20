## Database

You can run a simple Postgres database with:
```bash
podman run -d --name gosip-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:11
```

Once you have a database up and running you can migrate the database.
Migrations are done with a tool called [migrate](https://github.com/golang-migrate/migrate).

Here's an example command to migrate the database running in the container started earlier: 

```bash
migrate -source file://database/migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
```

Now that you have set up the correct tables in the database,
we have prepared some data that you can use to seed the database with:

```bash
podman exec -i gosip-postgres psql -d postgres -U postgres < database/seeds/simple.sql
```


### Creating seed files

```bash
podman exec -it gosip-postgres pg_dump -d postgres -U postgres --exclude-table=schema_migrations --data-only > database/seeds/{something}.sql
```
