cd service/$1/model || exit
goctl model mysql ddl -src ${1}.sql -dir .