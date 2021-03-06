env=".env.example"

if [ -e "$1" ]
then
    env=$1
fi

export $(cat $env | grep -v ^# | xargs)
go test ./... -coverprofile=coverage.out -race -coverpkg=./... -tags musl