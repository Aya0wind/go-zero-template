cd service/$1 || exit
goctl api go -api ${1}.api -dir .