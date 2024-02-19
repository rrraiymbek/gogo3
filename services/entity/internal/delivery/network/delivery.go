package network

import (
	contact2 "cleanArch/services/entity/internal/useCase/contact"
	"cleanArch/services/entity/internal/useCase/group"
	"net/http"
)

type ContactDelivery struct {
	useCase contact2.ContactUseCase
}
type GroupDelivery struct {
	useCase group.GroupUseCase
}

func NewContactDelivery(useCase contact2.ContactUseCase) *ContactDelivery {
	return &ContactDelivery{
		useCase: useCase,
	}
}
func NewGroupDelivery(useCase group.GroupUseCase) *GroupDelivery {
	return &GroupDelivery{
		useCase: useCase,
	}
}
func (d *ContactDelivery) HandleRequests(w http.ResponseWriter, r *http.Request) {
}
func (d *GroupDelivery) HandleRequests(w http.ResponseWriter, r *http.Request) {
}
