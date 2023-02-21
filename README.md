# Backend-Golang

Codificação de Backend em Golang usando biblioteca do Gorila Mux para desenvolvimento de API que definindo as rotas para cada endpoint da API.
E juntamente com a implementação de Teste Unitários para Golang que ajuda a cobertura de teste para toda aplicação. A aplicação está conteirnerizada para rodar em Docker


#  Realize a consulta dos principais Endpoints da aplicação:

/users método GET<br>
/users/{id} método Get por Id<br>
/users método POST<br>
/users/{id} método PUT<br>
/users/{id} método por Id DELETE<br>


# Rode o comando subir a imagem:
docker build -t Backend-Golang

# Para usar a imagem use esse comando:
docker run -p 8018:8018 Backend-Golang
