# Projeto Leilão
## Objetivo
Criar funcionalidade que após um tempo não seja possível cadastrar novas propostas.

## Validar
Na raiz do projeto:

1 - Executar aplicação
docker-compose up --build -d

2 - No terminal: go run internal/test/test.go
Vai exibir no terminal o passo a passo de cadastro e mostrar que a situação do leilão é alterada (nesse caso após 20 segundos)


## Alternativa
1 - acessar banco de dados: (Terminal)
docker exec -it mongodb_container mongosh -u admin -p admin --authenticationDatabase admin

2 - selectionar banco de dados
use auctions;

3 - Consultas

//consultar total de propostas
db.bids.count();

//Ver propostas cadastradas
db.bids.find();



