package models

import "web-application-project/db"

type Cliente struct {
	Id        int
	Nome      string
	Descricao string
	Cnpj      string
	Tipo      string
}

func BuscaTodosOsProdutos() []Cliente {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsClientes, err := db.Query("select * from clientes")
	if err != nil {
		panic(err.Error())
	}

	p := Cliente{}
	clientes := []Cliente{}

	for selectDeTodosOsClientes.Next() {
		var id int
		var nome, descricao, cnpj, tipo string

		err = selectDeTodosOsClientes.Scan(&id, &nome, &descricao, &cnpj, &tipo)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Cnpj = cnpj
		p.Tipo = tipo

		clientes = append(clientes, p)
	}
	defer db.Close()
	return clientes
}
func CriaNovoCliente(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, cnpj, tipo) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DeletaCliente(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()

}

func EditaCliente(id string) Cliente {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Cliente{}

	for produtoDoBanco.Next() {
		var id int
		var nome, descricao, cnpj, tipo string

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &cnpj, &tipo)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Cnpj = cnpj
		produtoParaAtualizar.Tipo = tipo
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaCliente(id int, nome, descricao, cnpj, tipo string) {
	db := db.ConectaComBancoDeDados()

	AtualizaCliente, err := db.Prepare("update clientes set nome=$1, descricao=$2, cnpj=$3, tipo=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaCliente.Exec(nome, descricao, cnpj, tipo, id)
	defer db.Close()
}
