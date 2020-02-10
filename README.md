# SuperHero

## Contexto
Para dar continuidade ao nosso processo temos um desafio! Gostamos muito de super heróis/vilões e queremos que você crie uma API para servir um jogo utilizando a SuperHeroAPI (https://superheroapi.com/).

## Requisitos

A API deve ser escrita em **Golang** e utilizar **PostgreSQL** como armazenamento.

### Gerais
Através da API deve ser possível:
- Cadastrar um Super/Vilão
- Listar todos os Super's cadastrados
- Listar apenas os Super Heróis
- Listar apenas os Super Vilões
- Buscar por nome
- Buscar por 'uuid'
- Remover o Super

### Específicos
- API deve seguir a arquitetura [REST](https://restfulapi.net/)
- Cada super deve ser cadastrado somente a partir do seu `name`.
- A pesquisa por um super deve conter os seguintes campos:
    - uuid
    - name
    - full name
    - intelligence
    - power
    - occupation
    - image
    - work
- A pesquisa por um super também precisa conter:
    - lista de grupos em que tal super está associado
    - número de parentes
    - quantos dos parentes são heróis/vilões

## Avaliação
A ideia aqui é entender como você toma suas decisões e como você desenvolve através de multiplas funcionalidades.

Pontos que vamos avaliar:
- Commits
    - como você evoluiu seu pensamento durante o projeto, pontualidade e clareza.
- Testes
    - Quanto mais testes melhor! Vide https://code.tutsplus.com/pt/tutorials/lets-go-testing-golang-programs--cms-26499 .
- Complexidade
    - Código bom é código legivel e simples (https://medium.com/trainingcenter/golang-d94e16d4b383).
- Dependências
    - O ecosistema (https://github.com/avelino/awesome-go) da linguagem possui diversas ferramentas para serem usadas, use-as bem!
- Documentação
    - Qual versão de Go você usou?
    - Quais bibliotecas e ferramentas usou?
    - Como se utiliza a sua aplicação?
    - Como executamos os testes?
- Considerações
    - as regras de negócio não foram definidas intencionalmente
    - cabe a você decidir como vai manter os cadastros no banco da aplicação
    - cabe a você decidir como vai tratar cadastros repetidos
