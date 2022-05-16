package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Rodrigueslcs/challenge-backend/api/presenter"
	"github.com/Rodrigueslcs/challenge-backend/usecase/video"
	"github.com/gorilla/mux"
)

func listVideos(service video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errorMessage := "Error reading Videos"
		data, err := service.ListVideos()
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

func getVideo(service video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		errorMessage := "Erro ao buscar Videos"
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetVideo(int(id))
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

func createVideo(service video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Erro ao adicionar Video"
		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateVideo(input.Title, input.Description, input.URL)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJSON := &presenter.Video{
			ID:          id,
			Title:       input.Title,
			Description: input.Description,
			URL:         input.URL,
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

func updateVideo(service video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Erro ao Alterar Video"
		var input presenter.Video
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.UpdateVideo(input.ID, input.Title, input.Description, input.URL)
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

func deleteVideo(service video.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Erro ao Excluir Video"
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteVideo(int(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func MakeBookHandlers(r *mux.Router, service video.UseCase) {
	r.Handle("/videos", listVideos(service)).Methods("GET", "OPTIONS").Name("listVideos")
	r.Handle("/videos/{id}", getVideo(service)).Methods("GET", "OPTIONS").Name("getVideo")
	r.Handle("/videos", createVideo(service)).Methods("POST", "OPTIONS").Name("createVideo")
	r.Handle("/videos", updateVideo(service)).Methods("PUT", "OPTIONS").Name("updateVideo")
	r.Handle("/videos/{id}", deleteVideo(service)).Methods("DELETE", "OPTIONS").Name("deleteVideo")

}
