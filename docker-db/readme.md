## Roda a aplicação
docker-compose up -d

## Acessar o banco dentro do container 
docker exec -it container-mariadb mariadb -u root -p ordersdb
    SENHA: root