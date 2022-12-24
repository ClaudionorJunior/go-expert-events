# Descrição do projeto:
Esse projeto faz parte do curso Go Expert da [FullCycle](https://fullcycle.com.br/). Foi desevolvido a partir dos aprendizados de como uma integração com RabbitMQ.
<br><br>
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/ClaudionorJunior/go-expert-events">

# Rodando o projeto:
- Para rotar o consumidor (receber os eventos)
```sh
cd cmd/consumer && go run main.go
```

- Para rotar a produtora (criar e enviar eventos)
```sh
cd cmd/producer && go run main.go
```

# Acessar RabbitMQ com Docker:
- Subindo o container
```sh
docker-compose up -d
```

- Página do RabbitMQ `http://localhost:15672`

## Autor
<table>
  <tr>
    <th><img src="https://avatars.githubusercontent.com/u/82416762?v=4" width=60></th>
  </tr>
  <tr>
    <td><a href="https://github.com/ClaudionorJunior">Github</a></td>
  </tr>
  <tr>
    <td><a href="https://www.linkedin.com/in/claudionorsilva">Linkedin</a></td>
  </tr>
</table>
