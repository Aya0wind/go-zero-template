cd service/${1} || exit
goctl api go -api ${1}.api -dir . || exit
cd model || exit
goctl model mysql ddl -src ${1}.sql -dir . || exit
cd ..
goctl docker  --go ${1}.go --port ${2} || exit