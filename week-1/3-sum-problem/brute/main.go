package main

import "fmt"


func main() {
        myset := []int{10,10,-20,40,-80,40,10,10,10,10,-10,-10}
        indexes := []int{}
        result := [][]int{}

        fmt.Println(myset)
        for i := 0; i < len(myset) ; i++ {
                for j := i+1; j < len(myset) ; j++ {
                        for k := j+1; k < len(myset) ; k++ {
                                if myset[i] + myset[j] + myset[k] == 0 {
                                        if ! contains(myset[i],myset[j],myset[k],indexes) {
                                                result = append(result,[]int{myset[i],myset[j],myset[k]})
                                                indexes = append(indexes, []int{myset[i],myset[j],myset[k]}...)
                                        }
                                }
                        }
                }
        }

        fmt.Println(result)
}

func contains(i,j,k int, indexes []int) bool {
        count := 0
        for _,v := range indexes {
                if v == i || v == j || v == k {
                        count += 1
                }

                if count == 3 {
                        return true
                }
        }

        return false
}

