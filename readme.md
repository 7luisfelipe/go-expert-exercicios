# Executar aplicação
* Terminal:
    * go mod tidy
    * docker-compose up --build -d

# Testar aplicação
* Na raiz do projeto tem um arquivo requisicoes.http

* zipkin: http://localhost:9411/api/v2/spans
* Prometheus: localhost:9090
* Metricas: http://localhost:8080/metrics