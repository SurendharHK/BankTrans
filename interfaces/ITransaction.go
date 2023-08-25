package interfaces

type ITransaction interface{
	Transfer(from string,to string,amount float64)(string,error)
}