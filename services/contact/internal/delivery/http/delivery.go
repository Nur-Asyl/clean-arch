package http

import (
	"architecture_go/services/contact/configs"
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

func (d *ContactHTTPDelivery) Run(cfg *configs.Config) {
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Println("Delivering... on port:", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panic("Something up with server delivering")
	}
}
