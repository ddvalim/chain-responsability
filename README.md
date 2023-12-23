<h1 align="center">Chain Responsability</h1>

<p align="center">Repositório criado para hospedar o código de estudo sobre o padrão de design Chain Responsability</p>

<p align="center">
  <img align="" alt="gopher" src="https://webartdevelopers.com/blog/wp-content/uploads/2021/11/gopher-golang-css-only-animation.gif" height="70%" width="70%" />
</p>

<p align="center">
 <a href="#Objetivo">Sobre o Projeto</a> •
 <a href="#Tecnologias">Tecnologias</a> • 
 <a href="#Pré-requisitos">Pré-requisitos</a> • 
 <a href="#Rodando o Servidor">Rodando o Servidor</a> • 
 <a href="#Dependências">Dependências</a> •
 <a href="#Autora">Autora</a> •
</p>

<a name="Objetivo"></a>
## 🖊 Sobre o Projeto

<p> 
O objetivo do projeto é desenvolver uma API Rest que implemente o padrão de design Chain of Responsability. Neste contexto, foi desenhado um CRUD para cadastrar, remover e recuperar usuários.
</p>

___
<a name="Tecnologias"></a>

## 🛠 Tecnologias
A API foi construída utilizando a linguagem Go, com a biblioteca nativa http.

___
<a name="Pré-requisitos"></a>

## ✅ Pré-requisitos
Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Git](https://git-scm.com), [Go](https://go.dev/), [MySQL](https://www.mysql.com/) e um editor de código a sua escolha.

___

<a name="Rodando o Servidor"></a>

## 🎲 Rodando o Servidor


- Clone este repositório na sua máquina:

      git clone https://github.com/ddvalim/chain-responsability.git

- Acesse a pasta do projeto no terminal:

      cd chain-responsability

- Instale as dependências do projeto:

      go mod install

- Preencha as variáveis de ambiente do arquivo `env.dist`

- Execute o arquivo SQL `init.sql`, da pasta `sql`

- Execute o servidor:

      go run main.go

O servidor estará sendo escutado na porta indicada na IDE. Geralmente, trata-se da porta 8686.

___
<a name="Dependências"></a>
## 🏁 Dependências

Foram utilizadas no projeto as seguintes dependências:

- 🍃 Mux;

- 🌐 Godotenv

___

<a name="Autora"></a>

## 📝 Autores ##

Desenvolvido com 💛 por Diovana Rodrigues Valim, graduada em Sistemas de Informção pela Universidade Federal de Santa Catarina e desenvolvedora de software no Mercado Livre.