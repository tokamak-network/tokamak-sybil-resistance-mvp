FROM golang:1.23.0-alpine

WORKDIR /app

ARG APP_MODE
ARG AWS_REGION
ARG TEST_AWS_DB_SECRET_NAME
ARG MAIN_AWS_DB_SECRET_NAME
ARG WAIT_TIME_SECONDS
ARG MAXTX
ARG NLEVEL

# TODO: a new address is created on every deployment
ARG KEYSTORE_PATH
ARG KEYSTORE_PASSWORD
ARG FORGER_ADDRESS

# TODO: ROLLUP_START_BLOCK_NUM is an awkward way of starting the sequencer
# It should instead save the latest processed block number in the database
# and start from there or if the database is empty, start from the latest block
ARG ROLLUP_START_BLOCK_NUM

ENV APP_MODE=${APP_MODE}
ENV AWS_REGION=${AWS_REGION}
ENV TEST_AWS_DB_SECRET_NAME=${TEST_AWS_DB_SECRET_NAME}
ENV MAIN_AWS_DB_SECRET_NAME=${MAIN_AWS_DB_SECRET_NAME}
ENV KEYSTORE_PATH=${KEYSTORE_PATH}
ENV KEYSTORE_PASSWORD=${KEYSTORE_PASSWORD}
ENV FORGER_ADDRESS=${FORGER_ADDRESS}
ENV ROLLUP_START_BLOCK_NUM=${ROLLUP_START_BLOCK_NUM}
ENV WAIT_TIME_SECONDS=${WAIT_TIME_SECONDS}
ENV MAXTX=${MAXTX}
ENV NLEVEL=${NLEVEL}

# for internal health check
RUN apk add --no-cache curl 

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 5000

CMD ["go", "run", "main.go", "run", "--cfg", "cfg.toml"]