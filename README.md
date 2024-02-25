# Scan IPv4

Este é um scanner IPv4 para Android. Ele foi projetado para escanear uma variedade de portas e fornecedores CDN.

## Como usar

**ANDROID - TERMUX**
```bash
wget https://github.com/adfastltda/scanv4/raw/main/scanv4 && chmod +x && ./scanv4 -h
```


### Opções

- `-cdn string`: O fornecedor CDN para escanear.
- `-limparcache`: Limpar o arquivo em cache dos fornecedores CDN.
- `-max int`: Número máximo de threads. O padrão é 1000.
- `-output string`: Arquivo para resultados de saída.
- `-port string`: Porta para escanear. O padrão é "443".
- `-timeout duration`: Duração do tempo limite de conexão. O padrão é 1s.

## Fornecedores CDN salvos
- cloudflare
- cloudfront
- fastly
- akamai
- edcast
- azure
- imperva
- securi
- openx
- facebook
- dosarrest
- maxcdn

## Exemplo

Aqui está um exemplo de como usar o scanner:

```bash
./scanv4 -cdn openx -output resultados.txt -port 80 -timeout 2s
