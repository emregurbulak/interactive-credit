FROM golang:1.9

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/emregurbulak/interactive-credit/api
WORKDIR /go/src/github.com/emregurbulak/interactive-credit/api

RUN go get ./vendor/database

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	api; \
	fi
	
EXPOSE 8080