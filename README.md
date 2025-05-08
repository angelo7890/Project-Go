# API de Ingressos 🎟️

Seja bem-vindo à API de Ingressos! Este projeto foi desenvolvido como parte da disciplina de Banco de Dados na faculdade, e o tema escolhido foi o gerenciamento da venda de ingressos para um show, com diferentes setores e toda a estrutura de vendas. O foco principal do trabalho foi implementar uma solução com código personalizado, que incluísse a manipulação de transações no banco de dados sem o uso de ORM, garantindo a integridade e o controle total das transações diretamente no banco.

## 🚀 Funcionalidades

- ✅ Criação de setores e eventos.
- ✅ Criação de usuários.
- ✅ Compra de ingressos.
- ✅ Atualização de setor dos ingressos.
- ✅ Uso de transações para garantir a integridade dos dados.

## 📂 Estrutura do Projeto

- `handler/` - Controladores responsáveis por lidar com as requisições HTTP, implementando diretamente a lógica de negócios e interagindo com o banco de dados.
- `repository/` - Manipulação direta dos dados no banco.
- `dto/` - Objetos de Transferência de Dados (Data Transfer Objects).
- `responses/` - Modelos de respostas padronizadas.
- `database/` - Configuração e gerenciamento da conexão com o banco de dados.
- `router/` - Definição e organização das rotas da API.

## 🚀 Como Executar

1. Clone o repositório.
2. Instale as dependências do Go com `go mod tidy`.
3. Configure o banco de dados.
4. Inicie o servidor com o comando:

```bash
 go run main.go
```

## 🌐 Como Usar

- Acesse a API em `http://localhost:8080/api`.
- Utilize um cliente HTTP como Postman ou Insomnia para testar os endpoints.
