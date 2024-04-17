// priority based queue -> ecommerce

package main

import (
	"fmt"
)

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue []*Order

// func (order *Order) New(priority int, quantity int, product string, customerName string) {
// }

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, appendOrder := range *queue {
			if order.priority > appendOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	queue := make(Queue, 0)
	order := &Order{1, 5, "juice", "John Doe"}
	appendOrder := &Order{2, 4, "milk", "Clint Eastwood"}
	queue.Add(order)
	queue.Add(appendOrder)
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
