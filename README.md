# go-weather-cli

Uma ferramenta de linha de comando feita em Go para consultar o clima de qualquer cidade do mundo, direto no terminal.

## Funcionalidades

- Temperatura atual e sensação térmica
- Umidade
- Descrição do clima

## Como usar

### Instalação

```bash
git clone https://github.com/MuriloZR/go-weather-cli.git
cd go-weather-cli
go build -o weather-cli .
```

### Executando

```bash
# Cidade simples
./weather-cli London

# Cidade com espaço
./weather-cli "São Paulo"
./weather-cli "New York"

# Direto com go run
go run main.go Tokyo
go run main.go "New York"
```

### Exemplo de saída
```
Tempo em Santa Maria:
Última Atualização em: 2026-03-25 01:54 AM
Partly cloudy
Temperatura: 19 °C
Sensação Térmica: 19 °C
Umidade: 91%
```
<!-- 
Em algum momento eu vou deixar assim
```
╔══════════════════════════════════════╗
║  Santa Maria
╠══════════════════════════════════════╣
║  Light rain shower
║  Temperatura:   22°C
║  Sensação:      20°C
║  Umidade:       85%
║  Vento:         18 km/h
╠══════════════════════════════════════╣
║  Próximos dias:
║    2025-01-20    ↑24°C  ↓18°C
║    2025-01-21    ↑27°C  ↓19°C
║    2025-01-22    ↑21°C  ↓17°C
╚══════════════════════════════════════╝
``` -->

## Tecnologias

- **Go** — linguagem principal
- **wttr.in** — API de clima gratuita, sem necessidade de chave de API
- Apenas bibliotecas da standard library (`net/http`, `encoding/json`)

## Como funciona

1. Recebe o nome da cidade como argumento
2. Faz uma requisição HTTP para a API [wttr.in](https://wttr.in)
3. Faz o parse do JSON retornado
4. Formata e exibe as informações no terminal

## 📄 Licença

MIT