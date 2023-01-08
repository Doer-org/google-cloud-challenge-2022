package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	json_res "github.com/Doer-org/google-cloud-challenge-2022/utils/http/json"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userUC usecase.IUserUsecase
	eventUC usecase.IEventUsecase
}

func NewUserHandler(uuc usecase.IUserUsecase,euc usecase.IEventUsecase) *UserHandler {
	return &UserHandler{
		userUC: uuc,
		eventUC: euc,
	}
}

// POST /users
func (h *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("request body is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}
	j := &json_res.UserJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()

	user, err := h.userUC.CreateNewUser(
		r.Context(),
		j.Name,
		j.Authenticated,
		j.Mail,
		j.Icon,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusCreated,
	)
}

// GET /users/{id}
func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.userUC.GetUserById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusOK,
	)
}

// DELETE /users/{id}
func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.userUC.DeleteUserById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		fmt.Sprintf("delete user success"), //TODO: Response Objectとかあったらよさそう
		http.StatusOK,
	)
}

// PATCH /users/{id}
func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("request body is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}
	j := &json_res.UserJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	idParam := chi.URLParam(r, "id")
	user, err := h.userUC.UpdateUserById(
		r.Context(),
		idParam,
		j.Name,
		j.Authenticated,
		j.Mail,
		j.Icon,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusCreated,
	)
}

// GET /users/{id}/events
func (h *UserHandler) GetUserEvents(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	events, err := h.eventUC.GetUserEvents(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonEvents(events),
		http.StatusOK,
	)
}


// TODO: openapiに追加する
// GET /users?mail=<user mail>
func (h *UserHandler) GetUserByMail(w http.ResponseWriter, r *http.Request) {
	mailQuery := r.URL.Query().Get("mail")
	user, err := h.userUC.GetUserByMail(r.Context(), mailQuery)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusOK,
	)
}
