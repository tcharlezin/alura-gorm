package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome      string
	Area      string
	Professor string
}

func main() {

	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao acessar o banco de dados")
	}

	migrate(db)
	// seeder(db)
	// var curso = finder(db)

	curso := Curso{
		Nome:      "Nome Teste",
		Area:      "Area Teste",
		Professor: "Professor Teste",
	}

	save(db, curso)

	// fmt.Println(curso.ID)
	// db.Delete(&Curso{}, 1)
	// db.Model(&curso).Update("Nome", "Teste de atualizacao de model")
}

func finder(db *gorm.DB) Curso {
	var curso Curso
	db.First(&curso, 1)
	return curso
}

func save(db *gorm.DB, curso Curso) {
	db.Save(&curso)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Curso{})
}

func seeder(db *gorm.DB) {
	db.Create(&Curso{
		Nome:      "Go: Fundamentos de uma aplicacao web",
		Area:      "Programacao",
		Professor: "Guilherme Lima",
	})
}
