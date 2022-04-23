jdbigen
=======

`jdbigen` generates Java interfaces and classes for simple DAO functions from PostgreSQL table metadata.
This based off of the awesome [dgw](https://github.com/achiku/dgw) project.

> **NOTE:** jdbigen only supports PostgreSQL at the moment 

## Installation

```
go install github.com/nndi-oss/jdbi-utils/jdbigen@latest
```

## How to use

```
usage: jdbigen [<flags>] <conn>

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -s, --schema="public"      PostgreSQL schema name
  -p, --package=""           package name
  -t, --typemap=TYPEMAP      column type and go type map file path
  -x, --exclude=EXCLUDE ...  table names to exclude
      --template=TEMPLATE    custom template path
  -o, --output=OUTPUT        output file path
      --sqlhandle            output as SQL handlers

Args:
  <conn>  PostgreSQL connection string in URL format
```

```
jdbigen postgres://dbuser@localhost/dbname?sslmode=disable 
```

## ROADMAP

0. Fix and improve the tests
1. I am planning on re-writing this in Java and compiling a native image with GraalVM.
2. Support other database engines?
