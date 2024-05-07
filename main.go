package main

func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}

func main(){ 
	resultado, soma := soma(10, 20)
	print("olaaa", resultado, soma)
	
}