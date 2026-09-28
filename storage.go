package main

type Storage [T any]{
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T]{
	return &Storage[T]{FileName: fileName}
}

func 

