package main

<<<<<<< HEAD
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello my friends!"
	fmt.Println(strings.Contains(greeting, "dog"))
	fmt.Println(strings.ReplaceAll(greeting, "my", "mine"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "my"))
	fmt.Println(strings.Split(greeting, "!"))
	ages := []int {50, 80, 10}
	sort.Ints(ages)
	fmt.Println(ages)
index := sort.SearchInts(ages,50)
fmt.Println(index)
names :=[]string{"Caroline", "Maicon", "Diego"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names, "caroline"))
x:=0

for x < 5 {
	fmt.Println(x)
	x++
}

for i:=0; i <5; i++{
fmt.Println("for 2: ", i )
}

for i:=0; i <len(names); i++{
	fmt.Println(names[i])
}

for index, value := range names {
	fmt.Println ("o indice é", index, "e o valor", value)

}


for i:=0; i <len(ages); i++{
	fmt.Println(ages[i])
}


for index, value := range ages{
	fmt.Println ("o indice é", index, "e o valor", value)

}

superherois := []string{"deadpool", "homem aranha", "batman"}

for index, value := range superherois {
	fmt.Println("o numero do herói:", index, "o nome do herói", value)

}

for i:=0; i < len(superherois); i++{
fmt.Println("o numero do herói", i, "e o herói", superherois[i])
}
}
=======
import "fmt"                                                      

func main(){

    var nomes = []string{"yasmim", "leandro", "junia", "arthur", "ane"}
    fmt.Println(nomes)
     nomesOne := nomes [:2]
     fmt.Println(nomesOne)
     nomesTwo := nomes[3:]
     fmt.Println(nomesTwo)
     rangeOne := nomes[2]
     fmt.Println(rangeOne)
}                                                                                                                                                                                                                                                                                    
>>>>>>> b6f418b4f42e0c98ff128d40e99855809c734f5a
