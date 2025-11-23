package usecase

import (
	"github.com/BMokarzel/clean-arch.git/internal/entity"
	"github.com/BMokarzel/clean-arch.git/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {

	ordersList, err := c.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	orders := make([]OrderOutputDTO, 0, len(ordersList))

	for _, order := range ordersList {
		orders = append(orders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orders, nil
}
