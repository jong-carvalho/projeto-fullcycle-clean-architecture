# Usa a versão mais recente do Go
FROM golang:1.23 AS builder

WORKDIR /app

# Copia apenas os arquivos de dependências primeiro
COPY go.mod go.sum ./

# Baixa as dependências do projeto antes de copiar o restante dos arquivos
RUN go mod download

# Agora copia os demais arquivos do projeto
COPY . .

# Depuração: Exibe os arquivos dentro do container antes de compilar
RUN ls -la /app

# Depuração: Mostra a versão do Go no container
RUN go version

# Tenta compilar e exibir mais informações em caso de erro
RUN go build -o main . || (echo "Erro na compilação" && exit 1)

# Imagem final minimalista
FROM alpine:latest
WORKDIR /root/

# Copia o binário compilado
COPY --from=builder /app/main .

EXPOSE 8080 50051 4000

CMD ["./main"]
