# Dekiteru

Dekiteru is a tool that checks if a service is ready to use.

## Examples

`Postgresql`:

```bash
dekiteru check --service postgresql \
               --parameter "dsn=postgres://localhost:5432?connect_timeout=5"
```

`Redis`:

```bash
dekiteru check --service redis --parameter "url=redis://localhost:16379/1"
```

`ElasticSearch`

```bash
dekiteru check --service elasticsearch \
               --parameter "url=http://localhost:9200"
```

## Help

```
$ dekiteru help
NAME:
   Dekiteru - Check if a service is ready to use

USAGE:
   dekiteru [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     check, c  Check if a service is ready to use
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
