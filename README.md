# cobd

Container toolbox to diagnose connection between databases, http services, etc...

## Usage

```shell
docker run -p 80:80 -e TEST_HTTP_GITHUB="https://github.com/" banst/cobd:latest 
Starting up on port 80
Running 1 test(s)...
Passed   Http https://github.com/
```

This will spinup a webserver on port 80 by default.

### Paths

- [/](localhost) Whoami similar to [containerous/whoami](https://github.com/containous/whoami).
- [/test](localhost/test) Run configured tests.
- [/env](localhost/env) Display current env, only if `COBD_ENABLE_ENV` is set. **Take extra care before exposing /env path publicly, as it could leak some critical informations (secrets, kube config, etc...).**

### Options

Options are given as env vars prefixed with `COBD_`

| Name            | Description                                         | Default |
| --------------- | --------------------------------------------------- | ------- |
| COBD_PORT       | Listening port.                                     | 80      |
| COBD_ENABLE_ENV | Enable the `/env` path. Set to any value to enable. | _Unset_ |


## Test environment variables structure

Each test variable must start with `COBD_T_`

| Pattern             | Description                                                                                                     | Expected value                                                                                           |
| ------------------- | --------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| COBD_T_REDIS        | Test a connection to redis                                                                                      | A redis [connection string](https://github.com/ServiceStack/ServiceStack.Redis#redis-connection-strings) |
| COBD_T_SQL_[DRIVER] | Test a connection to sql database. The driver should be one of a [sql supported driver](#sql-supported-drivers) | A sql connection string (check drivers for details)                                                      |
| COBD_T_HTTP_[NAME]  | Do a `HEAD` request to target. Succeed if the status code is >=200 && <400                                      | An http/https url, protocol included                                                                     |

## SQL supported drivers

- [MYSQL](github.com/go-sql-driver/mysql)
- [POSTGRES](github.com/jackc/pgx)
- [FIREBIRD](github.com/nakagami/firebirdsql)
- [ATHENA](github.com/segmentio/go-athena)