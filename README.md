# Desafio Nest.js 

## Informações do desafio
Neste desafio, você deve criar uma aplicação Nest.js com Docker que rode na porta 3000.


Esta aplicação precisa expor 2 rotas de API Rest:


- Listar routes - GET /api/routes
- Criar routes - POST /api/routes


Uma rota tem os seguintes dados:
```

id (gerado automaticamente pelo MongoDB)

name (nome da rota)

source (sub-documento que contém lat e lng)

destination (sub-documento que contém lat e lng)
```

O ORM a ser usado é o Prisma ORM e o banco de dados precisa ser o Mongo, image: bitnami/mongodb:5.0.17 ou similar


Crie o arquivo api.http para fazer as chamadas HTTP. Ao rodar o docker compose up já precisa subir logo de cara o projeto com o Nest.js rodando + o MongoDB.
