package cinema

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ygutara/xsis-test/helper"
	"github.com/ygutara/xsis-test/model"
)

func (cinema Cinema) MovieRoute(r *mux.Router) {
	r.HandleFunc("/Movie/create", cinema.MovieCreate).Methods("POST")
	r.HandleFunc("/Movie/{id:[0-9]+}", cinema.MovieGet).Methods("GET")
	r.HandleFunc("/Movie/update", cinema.MovieUpdate).Methods("PATCH")
	r.HandleFunc("/Movie/{id:[0-9]+}", cinema.MovieDelete).Methods("DELETE")
	r.HandleFunc("/Movie", cinema.MovieList).Methods("GET")
}

func (cinema Cinema) MovieCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	record := model.Movie{}
	if err := decoder.Decode(&record); err == nil {
		if err := cinema.MovieCreate_(&record); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieList(w http.ResponseWriter, r *http.Request) {
	if records, err := cinema.MovieList_(); err == nil {
		helper.WriteResponse(w, http.StatusOK, records, nil)
	} else {
		helper.WriteResponse(w, http.StatusInternalServerError, nil, model.ErrInternalServerError)
	}
}

func (cinema Cinema) MovieGet(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		if record, err := cinema.MovieGet_(id); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, model.ErrBadParamInput)
	}
}

func (cinema Cinema) MovieUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	record := model.Movie{}
	if err := decoder.Decode(&record); err == nil {
		if err := cinema.MovieUpdate_(&record); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieDelete(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		if err := cinema.MovieDelete_(id); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, model.ErrBadParamInput)
	}
}
