package main

import "fmt"

func main(){
    var ages = [5]int{17,16,20,40,50}
    fmt.Println(ages)
    // slice
    var score= []int{2,3,4,5,6}
    fmt.Println(score)
    fmt.Println(score, len(score), cap(score))
    score = append(score,1,7,8)
    fmt.Println(score, len(score), cap(score))
    }