# Proxy de notificações push Skymail Talk

Projeto para comunicação entre Google Firebase + Aplicações Mobile Skymail Talk

## Como subir a aplicação

Suba o container, ele compilar a imagem Docker e já compilar também a aplicação

```
docker-compose up -d --build
```

O container já utiliza um build watcher automático (https://github.com/cosmtrek/air), e a medida que você altera o código ele já compila com as modificações realizadas.

## Instruções VSCode

Para que o VSCode não compile automaticamente os modulos utilizando o go da sua máquina e assim gere conflito com a compilação do container, é necessário inserir uma configuração global no seu ambiente

1. No seu VSCode, vá em Configurações
2. Complete com "go tools env vars" e clique em "Editar em settings.json"
3. Insira GO111MODULE para "off" em go.toolsEnvVars

```json
{
    ...
    "go.toolsEnvVars": {
        "GO111MODULE": "off"
    }
}
```

---

Para mais instruções, visite a página do projeto original (https://github.com/mattermost/mattermost-push-proxy)

