# cobd

Container toolbox to diagnose connection between databases, http services, etc...

## Usage

```shell
docker run -p 80:80 -e TEST_HTTP_GITHUB="https://github.com/" banst/cobd:latest 
Starting up on port 80
Running 1 test(s)...
Passed   Http https://github.com/
```

This will spinup a webserver which run the tests on each refreshs.

## Environment structure

Each test variable must start with `TEST_`

| Pattern           | Description                                                                                                     | Expected value                                                                                           |
| ----------------- | --------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| TEST_REDIS        | Test a connection to redis                                                                                      | A redis [connection string](https://github.com/ServiceStack/ServiceStack.Redis#redis-connection-strings) |
| TEST_SQL_[DRIVER] | Test a connection to sql database. The driver should be one of a [sql supported driver](#sql-supported-drivers) | A sql connection string (check drivers for details)                                                      |
| TEST_HTTP_[NAME]  | Do a `HEAD` request to target. Succeed if the status code is >=200 && <400                                      | An http/https url, protocol included                                                                     |

## SQL supported drivers

- [MYSQL](github.com/go-sql-driver/mysql)
- [POSTGRES](github.com/jackc/pgx)
- [FIREBIRD](github.com/nakagami/firebirdsql)
- [ATHENA](github.com/segmentio/go-athena)