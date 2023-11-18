package views

import (
	"L0/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const Port = ":8080"

func MainPage(resp http.ResponseWriter, req *http.Request) {
	Title := "Main page"

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Получение родительского каталога от wd
	parentDir := filepath.Dir(wd)
	// Полный путь к файлу "index.html" в каталоге "templates"
	templatePath := filepath.Join(parentDir, "templates", "index.html")
	templatePathHeader := filepath.Join(parentDir, "templates", "includes", "_header.html")
	templatePathFooter := filepath.Join(parentDir, "templates", "includes", "_footer.html")
	templatePathStyle := filepath.Join(parentDir, "templates", "includes", "_style.html")

	tmpl, err := template.ParseFiles(templatePath, templatePathHeader, templatePathFooter, templatePathStyle)
	if err != nil {
		log.Fatalf("Error in parsing MAIN page template: %v\n", err)
	}

	var ordersList []models.Order
	for _, v := range models.Cache {
		ordersList = append(ordersList, v)
	}

	err = tmpl.ExecuteTemplate(resp, "index", struct {
		Title string
		List  []models.Order
	}{Title: Title, List: ordersList})
	if err != nil {
		log.Println("Error: can't execute template.")
	}
}

func MessageHandler(resp http.ResponseWriter, req *http.Request) {
	Title := "Message page"
	vars := mux.Vars(req)
	uid := vars["uid"]

	a, ok := models.Cache[uid]

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	parentDir := filepath.Dir(wd)

	if ok == false {
		errorTemplatePath := filepath.Join(parentDir, "templates", "errorPage.html")
		templatePathHeader := filepath.Join(parentDir, "templates", "includes", "_header.html")
		templatePathFooter := filepath.Join(parentDir, "templates", "includes", "_footer.html")
		templatePathStyle := filepath.Join(parentDir, "templates", "includes", "_style.html")
		tmpl, err := template.ParseFiles(errorTemplatePath, templatePathHeader, templatePathFooter, templatePathStyle)
		if err != nil {
			log.Fatalf("Error in parsing ERROR page template: %v\n", err)
		}

		Title = "Not found"

		err = tmpl.ExecuteTemplate(resp, "errorPage", struct {
			Title string
		}{Title: Title})
		if err != nil {
			log.Println("Error: can't execute ERROR page template.")
		}
	} else {

		messageTemplatePath := filepath.Join(parentDir, "templates", "message.html")
		templatePathHeader := filepath.Join(parentDir, "templates", "includes", "_header.html")
		templatePathFooter := filepath.Join(parentDir, "templates", "includes", "_footer.html")
		templatePathStyle := filepath.Join(parentDir, "templates", "includes", "_style.html")
		tmpl, err := template.ParseFiles(messageTemplatePath, templatePathHeader, templatePathFooter, templatePathStyle)
		if err != nil {
			log.Fatalf("Error in parsing MESSAGE page template: %v\n", err)
		}

		err = tmpl.ExecuteTemplate(resp, "message", struct {
			Title string
			Order models.Order
		}{Title: Title, Order: a})
		if err != nil {
			log.Println("Error: can't execute MESSAGE page template.")
		}
	}
}
