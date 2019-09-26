package veiculo

import "database/sql"

//é um contrato de implementação
type Storage interface {
	GetVeiculo() ([]Veiculo, error)
	CreateVeiculo(nome, marca string, ano, modelo int) error
	UpdateVeiculo(id int, veiculo *Veiculo) error
	DeleteVeiculo(id int) error
}

type MySQLStorage struct {
	dbConn *sql.DB
}

//func (dono - owner = mys=SQLstorage) nomeFuncao = GetVeiculo (parametros = não tem )(retorno = )
func (s *MySQLStorage) GetVeiculo() ([]Veiculo, error) {
	sql := "select id, nome, marca, ano, modelo from veiculos"
	rows, err := s.dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	//garante que será fechada a conexao ao término do método
	defer rows.Close()
	//define um slice - lista - de veiculos
	var veiculos []Veiculo

	for rows.Next() {
		//define variavel do tipo Veículo
		var veiculo Veiculo
		//1-pega o resultset (linhas - rows que retornara do banco)
		//2-pega o ponteiro & da variavel Veiculo e armazena os dados nela
		rows.Scan(&veiculo.ID, &veiculo.Nome, &veiculo.Marca, &veiculo.Ano, &veiculo.Modelo)
		//3-Pega o item (variavel veiculo) e adiciona no slice
	}
	return veiculos, nil
}
func (s *MySQLStorage) CreateVeiculo(nome, marca string, ano, modelo int) error {
	insert := "insert into veiculos (nome, marca, ano, modelo) values (?,?,?,?);"

	//prepara o banco de dados para receber os parametros
	stmt, err := s.dbConn.Prepare(insert)
	if err != nil {
		return err
	}
	//garante que será fechada a conexão
	defer stmt.Close()

	//executa a query que estava preparada com os parametros
	_, err = stmt.Exec(nome, marca, ano, modelo)
	if err != nil {
		return err
	}
	//se tudo correr bem retornará nil pointer, ou seja, sem erro

	return nil
}

func (s *MySQLStorage) UpdateVeiculo(id int, veiculo *Veiculo) error {
	update := "update veiculos set nome=?, marca=?, ano=? modelo=? where id=?;"
	stmt, err := s.dbConn.Prepare(update)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(veiculo.Nome, veiculo.Marca, veiculo.Ano, veiculo.Modelo, id)
	return err

	return nil

}

func (s *MySQLStorage) Delete(id int) error {
	deleteSQL := "delete from veiculos where id=?"
	stmt, err := s.dbConn.Prepare(deleteSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func NewStorage(conStr string) MySQLStorage {
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		panic("MySQL connection has faled!")
	}
	return MySQLStorage{
		dbConn: conn,
	}
}
