package playground

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Product  string
	Quantity int
}

func process(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("🔵 [Order %d] Starting processing of %d x %s\n",
		order.ID, order.Quantity, order.Product)

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("⚙️  [Order %d] Checking inventory...\n", order.ID)

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("💳 [Order %d] Processing payment...\n", order.ID)

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("📦 [Order %d] Preparing for shipment...\n", order.ID)

	fmt.Printf("✅ [Order %d] COMPLETED\n\n", order.ID)
}

func RunGroups() {
	var wg sync.WaitGroup

	orders := []Order{
		{ID: 1, Product: "Laptop", Quantity: 1},
		{ID: 2, Product: "Mouse", Quantity: 2},
		{ID: 3, Product: "Keyboard", Quantity: 1},
		{ID: 4, Product: "Monitor", Quantity: 1},
		{ID: 5, Product: "Webcam", Quantity: 3},
	}

	fmt.Printf("📋 Total orders to process: %d\n\n", len(orders))
	startTime := time.Now()

	for _, order := range orders {
		wg.Add(1)
		go process(order, &wg)
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
	elapsed := time.Since(startTime)

	fmt.Println("================================================")
	fmt.Printf("\n✨ ALL ORDERS PROCESSED ✨\n")
	fmt.Printf("⏱️  Total time: %.2f seconds\n", elapsed.Seconds())
	fmt.Println("\nWithout concurrency it would have taken ~7.5 seconds")
	fmt.Printf("With concurrency it took: %.2f seconds\n", elapsed.Seconds())
}
