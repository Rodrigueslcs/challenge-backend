package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Rodrigueslcs/challenge-backend/api/presenter"
	"github.com/Rodrigueslcs/challenge-backend/usecase/category"
	"github.com/gorilla/mux"
)

func listCategories(service category.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errorMessage := "Erro ao escrever list de Categorias"
		data, err := service.ListCategories()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func getCategory(service category.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		errorMessage := "Erro ao buscar Categoria"
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetCategory(int(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createCategory(service category.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Erro ao adicionar Video"
		var input struct {
			Title string `json:"title"`
			Color string `json:"color"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateCategory(input.Title, input.Color)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		toJSON := &presenter.Category{
			ID:    id,
			Title: input.Title,
			Color: input.Color,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateCategory(service category.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Erro ao Alterar Categoria"
		var input presenter.Category
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.UpdateCategory(input.ID, input.Title, input.Color)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(input); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func deleteCategory(service category.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Erro ao Excluir Categoria"
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteCategory(int(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func MakeCategoryHandlers(r *mux.Router, service category.UseCase) {
	r.Handle("/categories", listCategories(service)).Methods("GET", "OPTIONS").Name("listCategories")
	r.Handle("/categories/{id}", getCategory(service)).Methods("GET", "OPTIONS").Name("getCategory")
	r.Handle("/categories", createCategory(service)).Methods("POST", "OPTIONS").Name("createCategory")
	r.Handle("/categories", updateCategory(service)).Methods("PUT", "OPTIONS").Name("updateCategory")
	r.Handle("/categories/{id}", deleteCategory(service)).Methods("DELETE", "OPTIONS").Name("deleteCategory")

}
