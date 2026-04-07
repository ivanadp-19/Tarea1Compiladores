package main

import (
	"fmt"

	"datastructures/ds"
)

func main() {
	fmt.Println("Demo de estructuras de datos")

	demoHashMap()
	demoStack()
	demoQueue()
}
	
func demoHashMap() {
	fmt.Println("\n--- HashMap (bucket = key % 10; Set/Get/Remove) ---")
	h := ds.NewHashMap()
	h.Set(1, 10)
	h.Set(2, 20)
	h.Set(11, 110) // colisión con clave 1 (mismo bucket índice 1)
	fmt.Println("Tabla tras Set(1,10), Set(2,20), Set(11,110):")
	h.Print()
	fmt.Println("Get(1) — primer valor en el bucket 1:", h.Get(1))
	fmt.Println("Get(2):", h.Get(2))
	h.Remove(1)
	fmt.Println("Tras Remove(1) en bucket 1 (quita primer elemento del slice):")
	h.Print()
	fmt.Println("Get(1) ahora (siguiente en bucket):", h.Get(1))
}

func demoStack() {
	fmt.Println("\n--- Stack (LIFO: último en entrar, primero en salir) ---")
	s := &ds.Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Print("Tras Push 100, 200, 300 → ")
	s.Print()
	fmt.Println("Pop (debe ser 300):", s.Pop())
	fmt.Print("Tras un Pop → ")
	s.Print()
	fmt.Println("Pop (debe ser 200):", s.Pop())
	fmt.Println("Pop (debe ser 100):", s.Pop())
	fmt.Print("Stack vacío → ")
	s.Print()
}

func demoQueue() {
	fmt.Println("\n--- Queue (FIFO: primero en entrar, primero en salir) ---")
	q := &ds.Queue{}
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Print("Tras Enqueue 100, 200, 300 → ")
	q.Print()
	fmt.Println("Dequeue (debe ser 100):", q.Dequeue())
	fmt.Print("Tras un Dequeue → ")
	q.Print()
	fmt.Println("Dequeue (debe ser 200):", q.Dequeue())
	fmt.Println("Dequeue (debe ser 300):", q.Dequeue())
	fmt.Print("Queue vacía → ")
	q.Print()
}