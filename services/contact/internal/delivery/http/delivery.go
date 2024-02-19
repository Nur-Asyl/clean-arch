package http

import "architecture_go/services/contact/internal/useCase"

type ContactHTTPDelivery struct {
	contactUseCase useCase.ContactUseCase
	groupUseCase   useCase.GroupUseCase
}

func NewContactHTTP(contactUC useCase.ContactUseCase, groupUC useCase.GroupUseCase) *ContactHTTPDelivery {
	return &ContactHTTPDelivery{contactUseCase: contactUC, groupUseCase: groupUC}
}
