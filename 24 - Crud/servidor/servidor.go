package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, er := ioutil.ReadAll(r.Body)

	if er != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	var user user

	if er = json.Unmarshal(requestBody, &user); er != nil {
		w.Write([]byte("Error ao converter user para o struct"))
		return
	}

	db, er := banco.Connect()
	if er != nil {
		w.Write([]byte("Falha ao connectar com o db"))
		return
	}
	defer db.Close()

	statement, er := db.Prepare("insert into usuarios(nome, email) values (?,?)")
	if er != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insert, er := statement.Exec(user.Nome, user.Email)
	if er != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idUser, er := insert.LastInsertId()
	if er != nil {
		w.Write([]byte("Erro ao obter o ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserido com sucesso! Id: %d", idUser)))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db, er := banco.Connect()
	if er != nil {
		w.Write([]byte("Falha ao connectar com o db"))
		return
	}
	defer db.Close()

	rows, er := db.Query("SELECT * FROM usuarios")
	if er != nil {
		w.Write([]byte("Falha ao buscar usuarios"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if er := rows.Scan(&user.ID, &user.Nome, &user.Email); er != nil {
			println(&user.ID, &user.Nome, &user.Email)
			w.Write([]byte("Erro ao ler lista de usuários"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if er := json.NewEncoder(w).Encode(users); er != nil {
		w.Write([]byte("Erro ao converte lista de usuários em JSON"))
		return
	}
}
