package main


func main()  {
	todos := Todos{}
	todos.add("buy milk")
	todos.add("learn Go")
	todos.toggle(0)
	todos.print()

}