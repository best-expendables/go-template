FROM golang:1.15

ARG ACCESS_TOKEN_USR="nothing"
ARG ACCESS_TOKEN_PWD="nothing"

# Create a netrc file using the credentials specified using --build-arg
RUN echo ${ACCESS_TOKEN_USR}
RUN echo ${ACCESS_TOKEN_PWD}
RUN printf "machine bitbucket.org\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.bitbucket.org\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

ENV TZ=Asia/Manila
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN set -e; \
    go get github.com/go-playground/overalls; \
    go get github.com/swaggo/swag/cmd/swag

WORKDIR /go/src/${REPO_HOST}/${PROJ_NAME}

ENV GOSUMDB=off

COPY go.mod .
COPY go.sum .

COPY ./docker/Application/bin /usr/local/bin/app

ARG IS_DEV_MODE

RUN chmod +x /usr/local/bin/app/* \
    && if [ -n $IS_DEV_MODE ]; then /usr/local/bin/app/install_dev.sh; fi

COPY ./ /go/src/${REPO_HOST}/${PROJ_NAME}

RUN set -ex; \
    go build -o ./build/http cmd/http/http.go \
    && go build -o ./build/migrate cmd/migrate/migrate.go \
    && go build -o ./build/seeder cmd/seeder/seeder.go

RUN swag init -d cmd/http -g http.go -o ./build --parseDependency

CMD /usr/local/bin/app/http.sh
