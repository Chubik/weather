vendor:
        GO111MODULE=on go mod init
        GO111MODULE=on go mod vendor
run:
        docker-compose stop
        docker-compose up --build -d
dclean:
        docker system prune%