package main

type idsFlag []string

func(ids idsFlag)String() string{
	return ""
}

func(ids idsFlag)Set(id string) error{
	return nil
}

func main(){
	
}