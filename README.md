# Dekiteru

Dekiteru is a tool that check if a service is ready to use.

## Examples

`Postgresql`:

    DEKITERU_POSTGRESQL_DSN="postgres://localhost:5432?connect_timeout=5" dekiteru --service postgresql

`Redis`:

    DEKITERU_REDIS_URL="redis://localhost:16379/1" dekiteru --service redis

`ElasticSearch`

    DEKITERU_ELASTICSEARCH_URL="http://localhost:9200" dekiteru --service elasticsearch

## Help

```
$ ./bin/dekiteru --help
NAME:
   Dekiteru - Check if a service is ready to use

USAGE:
   dekiteru [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --service value, -s value   Service name to check (required)
   --interval value, -i value  Interval between retries in second (default: 1)
   --retry value, -r value     Number of retry (default: 10)
   --help, -h                  show help
   --version, -v               print the version
```