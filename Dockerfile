# Use a imagem base do Go
FROM golang:1.23 AS builder

# Instalar o protoc
RUN apt-get update && \
    apt-get install -y \
    curl \
    unzip \
    && curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.18.0/protoc-3.18.0-linux-x86_64.zip \
    && unzip protoc-3.18.0-linux-x86_64.zip -d /usr/local/ \
    && rm protoc-3.18.0-linux-x86_64.zip

# Instalar o plugin protoc-gen-go (responsável por gerar código Go a partir de arquivos .proto)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adiciona o diretório bin do Protoc e do Go ao PATH
ENV PATH="/usr/local/bin:$PATH"
ENV PATH="$PATH:$(go env GOPATH)/bin"

# Definir o diretório de trabalho
WORKDIR /app

# Copia o arquivo proto para o contêiner
COPY proto /app/proto

# Baixar as dependências do projeto Go
COPY go.mod go.sum ./
RUN go mod download

# Gerar os arquivos Go a partir do arquivo .proto
RUN protoc --go_out=. --go-grpc_out=. proto/order.proto

# Copia o código-fonte do aplicativo
COPY . .

# Compila o código Go
RUN go build -o main .

# Imagem final minimalista
FROM alpine:latest

WORKDIR /root/

# Copia o binário compilado
COPY --from=builder /app/main .

# Expõe as portas
EXPOSE 8080 50051 4000

# Comando para rodar a aplicação
CMD ["./main"]
