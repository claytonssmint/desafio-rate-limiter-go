# Rate Limiter em Go com Redis

Este projeto implementa um Rate Limiter em Go, projetado para controlar o número de requisições a uma API com base em dois critérios: 
endereço IP ou token de acesso. Ele usa o Redis como mecanismo de persistência para armazenar os limites temporários


# Como o Rate Limiter Funciona
1. Conceitos Principais
Limitação por IP: Restringe o número de requisições por segundo de um único endereço IP.
Limitação por Token de Acesso: Permite diferentes limites por segundo com base em um token de acesso informados no cabeçalho da requisição API_KEY.

# Como Funciona
Cada requisição é interceptada pelo Middleware do rate limiter.
O middleware identifica o cliente pelo endereço IP ou pelo cabeçalho API_KEY.
O Rate Limiter verifica no Redis o número de requisições feitas pelo IP ou Token no intervalo configurado.
Se o cliente não excedeu o limite:
A requisição é permitida e os contadores são atualizados.
Caso o cliente exceda o limite:
O middleware retorna um código HTTP 429 (Too Many Requests) com a mensagem:

{
"error": "you have reached the maximum number of requests or actions allowed within a certain time frame"
}

# Para rodar o projeto localmente siga os passos abaixo:

### 1. Clone o repositório
### 2. Execute o docker-compose up --build
### 3. O servidor web estará disponível em GET http://localhost:8080/