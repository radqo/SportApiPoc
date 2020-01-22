package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/radqo/SportApiPoc/model"

	"github.com/gorilla/mux"
)

//Server - rest api
type Server struct {
	PlayerInfo model.PlayerInfoFinder
	httpServer *http.Server
}

//GetPlayerInfo - player info handler
func (s *Server) GetPlayerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	surname := vars["surname"]

	info, err := s.PlayerInfo.FindPlayer(surname)

	if err != nil {
		http.Error(w, err.Error(), err.Code)

	} else {

		body, err := json.Marshal(info)

		if err != nil {
			http.Error(w, "Serialization error", 500)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(body)
		}

	}
}

// Run - starts server
func (s *Server) Run(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/player/{surname}", s.GetPlayerInfo).Methods(http.MethodGet)

	s.httpServer = &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: r}

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
