# Lab Stress Test

Ferramenta CLI em Go para realizar testes de carga em serviços web. Permite configurar a URL alvo, o numero total de requests e o nivel de concorrencia.

## Parametros

| Flag | Alias | Descricao |
|------|-------|-----------|
| `--url` | `-u` | URL do servico a ser testado |
| `--requests` | `-r` | Numero total de requests |
| `--concurrency` | `-c` | Numero de chamadas simultaneas |

## Relatorio

Apos a execucao, o sistema exibe:

- Tempo total gasto na execucao
- Quantidade total de requests realizados
- Quantidade de requests com status HTTP 200
- Distribuicao de outros codigos de status HTTP (404, 500, etc.)
- Quantidade de requests com erro

## Executando localmente

```bash
# Request GET simples
go run main.go request --url=http://google.com --requests=100 --concurrency=10

# Endpoint que retorna JSON (status 200)
go run main.go request --url=https://httpbin.org/get --requests=50 --concurrency=5

# Simulando respostas com status 404
go run main.go request --url=https://httpbin.org/status/404 --requests=30 --concurrency=3

# Simulando respostas com status 500
go run main.go request --url=https://httpbin.org/status/500 --requests=30 --concurrency=3

# Simulando respostas com delay de 1 segundo
go run main.go request --url=https://httpbin.org/delay/1 --requests=20 --concurrency=10
```

## Executando os testes

```bash
go test ./...
```

## Docker

### Build da imagem

```bash
docker build -t stress-test .
```

### Executando via Docker

```bash
# Request basico
docker run stress-test --url=http://google.com --requests=1000 --concurrency=10

# Testando com httpbin
docker run stress-test --url=https://httpbin.org/get --requests=100 --concurrency=10

# Testando respostas 404
docker run stress-test --url=https://httpbin.org/status/404 --requests=50 --concurrency=5
```

### Removendo a imagem

```bash
docker rmi stress-test
```
