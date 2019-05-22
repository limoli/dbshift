# DbShift

DbShift is a simple **database migrations tool** with the goal of simplicity. 
You will be able to create migrations, check the current db status, decide to upgrade or downgrade easily.

[![GoDoc](https://godoc.org/limoli/dbshift?status.svg)](https://godoc.org/github.com/limoli/dbshift)
[![Build Status](https://travis-ci.org/limoli/dbshift.svg?branch=master)](https://travis-ci.org/limoli/dbshift)
[![Go Report Card](https://goreportcard.com/badge/github.com/limoli/dbshift)](https://goreportcard.com/report/github.com/limoli/dbshift)
[![Maintainability](https://api.codeclimate.com/v1/badges/76ab43d259213895e8dd/maintainability)](https://codeclimate.com/github/limoli/dbshift/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/76ab43d259213895e8dd/test_coverage)](https://codeclimate.com/github/limoli/dbshift/test_coverage)

## Install

```bash
go get github.com/limoli/dbshift
```

## Tutorial

0. Create your [configuration file](#configuration) and set its path using `DBSHIFT_CONFIG`

1. Initialise your configuration:
```bash
dbshift init
```

2. Create migration. The following command will create two files `$uuid.down.sql` and `$uuid.up.sql` at your migrations folder.
```bash
dbshift create my-migration-description
```
   
3. Check status of your current database.
```bash
dbshift status
```

4. Upgrade migrations.
```bash
dbshift upgrade
```

5. Downgrade migrations.    
```bash
dbshift downgrade
```
6. Get list of available upgrade migrations.
```bash
dbshift migrations-upgrade
```
    
7. Get list of available downgrade migrations.
```bash
dbshift migrations-downgrade
```    

8. Get settings info
```bash
dbshift info
```    

9. Import missing migrations. It is useful when you work on different initialised databases.
```bash
dbshift refresh
```    


## Write good migrations

1. Queries must be database name **agnostic**
2. [SRP](https://en.wikipedia.org/wiki/Single_responsibility_principle) according to your description
3. Write both upgrade and downgrade migrations 

## Configuration

All you need is an environment variable `DBSHIFT_CONFIG` which represents the path where **dbshift** will find:
- a configuration file `dbshift.yaml` written by you
- an autogenerated lockfile managed from dbshift

```bash
export DBSHIFT_CONFIG=/absolutePath/dbshift.yaml
```

The configuration file has the following structure:

```yaml
db:
  type: mysql
  migration:
    path: /absolutePath/migrations
    pathEnv: "MYSQL_MIGRATIONS_PATH"
  connection:
    name:
      env: MYSQL_DATABASE
      value:
    user:
      env: MYSQL_USER
      value:
    password:
      env: MYSQL_PASSWORD
      value:
    host:
      env: MYSQL_HOST
      value:
    port:
      env: MYSQL_PORT
      value: 3306
```      
	
The connection object is composed by fields which values can be set from the `value` field or
from an environment variable declared in the `env` field.
	
The environment variable overrides the value if set.

```bash
export MYSQL_DATABASE=test
export MYSQL_USER=root
export MYSQL_PASSWORD=root
export MYSQL_HOST=127.0.0.1
export MYSQL_PORT=3306
```

## DB compatibility

At the moment the project is focused only on **MySQL**, but it is opened to implementations.

## Future implementations

1. Improve test coverage
2. Add possibility to upgrade/downgrade to a specific migration
3. Add command to restore dbshift in case of transactional errors
4. Add command to remove a migration
5. Add command to get list of all migrations with description and uuid reference
