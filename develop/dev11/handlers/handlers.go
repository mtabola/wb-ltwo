package handlers

import (
	"dev11/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	db *storage.Storage
}

func NewHandler(db *storage.Storage) *Handler {
	return &Handler{db: db}
}

// POST запросы
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		er := ErrorResponse{
			Code:  http.StatusMethodNotAllowed,
			Error: fmt.Sprintf("unsupported method %s", r.Method),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	var req storage.Event

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		er := ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}
	req.Date = time.Now()

	err = h.db.CreateEvent(req)
	if err != nil {
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}
	h.db.SaveToFile()

	sr := SuccessResponse{
		Code:   http.StatusOK,
		Result: "Data is created successfully",
	}
	json.NewEncoder(w).Encode(sr)

}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		er := ErrorResponse{
			Code:  http.StatusMethodNotAllowed,
			Error: fmt.Sprintf("unsupported method %s", r.Method),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	var req storage.Event

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		er := ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	err = h.db.UpdateEvent(req)
	if err != nil {
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}
	h.db.SaveToFile()
	sr := SuccessResponse{
		Code:   http.StatusOK,
		Result: "Data is updated successfully",
	}
	json.NewEncoder(w).Encode(sr)
}

// GET запросы
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		er := ErrorResponse{
			Code:  http.StatusMethodNotAllowed,
			Error: fmt.Sprintf("unsupported method %s", r.Method),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	taskId, _ := strconv.Atoi(r.URL.Query().Get("task_id"))

	err := h.db.DeleteEvent(userId, taskId)
	if err != nil {
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	h.db.SaveToFile()

	sr := SuccessResponse{
		Code:   http.StatusBadRequest,
		Result: "Data is deleted successfully",
	}
	json.NewEncoder(w).Encode(sr)

}

func (h *Handler) ShowEventsByDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		er := ErrorResponse{
			Code:  http.StatusMethodNotAllowed,
			Error: fmt.Sprintf("unsupported method %s", r.Method),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	userId, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
	date, err2 := time.Parse("2006-01-02", r.URL.Query().Get("date"))

	if err1 != nil || err2 != nil {
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	sr := SuccessResponse{
		Code:   http.StatusOK,
		Result: h.db.GetEvents(userId, date, storage.Day),
	}
	json.NewEncoder(w).Encode(sr)
}

func (h *Handler) ShowEventsByWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		er := ErrorResponse{
			Code:  http.StatusMethodNotAllowed,
			Error: fmt.Sprintf("unsupported method %s", r.Method),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	userId, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
	date, err2 := time.Parse("2006-01-02", r.URL.Query().Get("date"))

	if err1 != nil || err2 != nil {
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	sr := SuccessResponse{
		Code:   http.StatusOK,
		Result: h.db.GetEvents(userId, date, storage.Week),
	}
	json.NewEncoder(w).Encode(sr)
}

func (h *Handler) ShowEventsByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
	date, err2 := time.Parse("2006-01-02", r.URL.Query().Get("date"))

	if err1 != nil || err2 != nil {
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		er := ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(er)
		return
	}

	sr := SuccessResponse{
		Code:   http.StatusOK,
		Result: h.db.GetEvents(userId, date, storage.Month),
	}
	json.NewEncoder(w).Encode(sr)
}
