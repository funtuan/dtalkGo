FROM golang:1.12

WORKDIR /dtalkGO

COPY . /dtalkGO

RUN go mod download

ENTRYPOINT ["go", "run", "./"]