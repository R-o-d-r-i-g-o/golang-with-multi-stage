# Projeto Go com Docker Multi-Stage

Este é um exemplo simples de um projeto Go usando Docker multi-stage build. <br>
O projeto consiste em um aplicativo Go que imprime "Hello, World!" na tela.

De fato é bem simples o projeto, mas seu diferencial se encontra no tamanho da imagem gerada (como sugere o fragmento de código abaixo). <br>

Se desejar contribuir para a performance deste pequeno projeto, sinta-se avontade para abrir issues e sugerir PRs. <br>
Conto com sua participação :D

### Dados da imagem:
 ```bash
REPOSITORY                 TAG           IMAGE ID       CREATED          SIZE
rodrigomarq/my-golang-app  latest        fa72b8aac68e   51 seconds ago   9.14MB
 ```

## Como usar

1. Clone o repositório:

   ```bash
   git clone https://github.com/R-o-d-r-i-g-o/golang-with-multi-stage.git
   cd golang-with-multi-stage
   ```

2. Faça o build local da imagem:

   ```bash
   docker build -t my-golang-app .
   docker run my-golang-app
   ```

## Como usar (DockerHub)

1. Execute a imagem:

   ```bash
   docker run rodrigomarq/my-golang-app
   ```
