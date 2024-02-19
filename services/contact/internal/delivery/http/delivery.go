package http

import (
	"architecture_go/services/contact/internal/useCase"
	"fmt"
	"log"
	"net/http"
)

type ContactHTTPDelivery struct {
	contactUseCase useCase.ContactUseCase
	groupUseCase   useCase.GroupUseCase
}

func NewContactHTTP(contactUC useCase.ContactUseCase, groupUC useCase.GroupUseCase) *ContactHTTPDelivery {
	return &ContactHTTPDelivery{contactUseCase: contactUC, groupUseCase: groupUC}
}

func (d *ContactHTTPDelivery) Run(port string) {
	fmt.Println("Delivering...")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("Something up with server delivering")
	}

}
