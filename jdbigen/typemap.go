package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var defaultTypeMapCfg PgTypeMapConfig

func init() {
	if _, err := toml.Decode(typeMap, &defaultTypeMapCfg); err != nil {
		log.Fatal(err)
	}
}

const typeMap = `
[string]
db_types = ["character", "character varying", "text", "money"]
notnull_go_type = "String"
nullable_go_type = "String"

[time]
db_types = [
    "time with time zone", "time without time zone",
    "timestamp without time zone", "timestamp with time zone", "date"
]
notnull_go_type = "java.sql.Timestamp"
nullable_go_type = "java.sql.Timestamp"

[bool]
db_types = ["boolean"]
notnull_go_type = "boolean"
nullable_go_type = "Boolean"

[smallint]
db_types = ["smallint"]
notnull_go_type = "integer"
nullable_go_type = "Integer"

[integer]
db_types = ["integer"]
notnull_go_type = "integer"
nullable_go_type = "Integer"

[bigint]
db_types = ["bigint"]
notnull_go_type = "long"
nullable_go_type = "Long"

[smallserial]
db_types = ["smallserial"]
notnull_go_type = "integer"
nullable_go_type = "Integer"

[serial]
db_types = ["serial"]
notnull_go_type = "integer"
nullable_go_type = "Integer"

[real]
db_types = ["real"]
notnull_go_type = "double"
nullable_go_type = "Double"

[numeric]
db_types = ["numeric", "double precision"]
notnull_go_type = "double"
nullable_go_type = "Double"

[bytea]
db_types = ["bytea"]
notnull_go_type = "byte"
nullable_go_type = "Byte"

[json]
db_types = ["json", "jsonb"]
notnull_go_type = "byte[]"
nullable_go_type = "byte[]"

[xml]
db_types = ["xml"]
notnull_go_type = "byte[]"
nullable_go_type = "byte[]"

[interval]
db_types = ["interval"]
notnull_go_type = "java.time.Duration"
nullable_go_type = "java.time.Duration"

[default]
db_types = ["*"]
notnull_go_type = "Object"
nullable_go_type = "Object"
`
