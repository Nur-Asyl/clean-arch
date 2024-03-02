package http

import (
	"architecture_go/services/contact/configs"
	"architecture_go/services/contact/internal/useCase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ContactHTTPDelivery struct {
	contactUC useCase.ContactUseCase
	groupUC   useCase.GroupUseCase
}

func NewContactHTTP(contactUC useCase.ContactUseCase, groupUC useCase.GroupUseCase) *ContactHTTPDelivery {
	return &ContactHTTPDelivery{contactUC: contactUC, groupUC: groupUC}
}

func (ch *ContactHTTPDelivery) CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestData struct {
		FullName    string `json:"full_name"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newContact, err := ch.contactUC.CreateContact(ctx, requestData.FullName, requestData.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newContact)
}

func (ch *ContactHTTPDelivery) GetContactByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contactID, err := strconv.Atoi(r.URL.Query().Get("contact_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingContact, err := ch.contactUC.ReadContact(ctx, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingContact)
}

func (ch *ContactHTTPDelivery) UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestData struct {
		ContactID   int    `json:"contact_id"`
		FullName    string `json:"full_name"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ch.contactUC.UpdateContact(ctx, requestData.ContactID, requestData.FullName, requestData.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ch *ContactHTTPDelivery) DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contactID, err := strconv.Atoi(r.URL.Query().Get("contact_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ch.contactUC.DeleteContact(ctx, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ch *ContactHTTPDelivery) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestData struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGroup, err := ch.groupUC.CreateGroup(ctx, requestData.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGroup)
}

func (ch *ContactHTTPDelivery) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingGroup, err := ch.groupUC.ReadGroup(ctx, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingGroup)
}

func (ch *ContactHTTPDelivery) AddContactToGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestData struct {
		ContactID int `json:"contact_id"`
		GroupID   int `json:"group_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ch.groupUC.AddContactToGroup(ctx, requestData.GroupID, requestData.ContactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (d *ContactHTTPDelivery) Run(cfg *configs.Config) {
	addr := fmt.Sprintf(":%s", cfg.Port)

	http.HandleFunc("/contact/create", d.CreateContactHandler)
	http.HandleFunc("/contact/get", d.GetContactByIDHandler)
	http.HandleFunc("/contact/update", d.UpdateContactHandler)
	http.HandleFunc("/contact/delete", d.DeleteContactHandler)

	http.HandleFunc("/group/create", d.CreateGroupHandler)
	http.HandleFunc("/group/get", d.GetGroupByIDHandler)
	http.HandleFunc("/group/addContact", d.AddContactToGroupHandler)

	fmt.Println("Delivering... on port:", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panic("Something up with server delivering")
	}
}
