package main

import (
    "fmt"
    "math"
    )

func factorial(n uint) uint {
    if n == 0 {
        return 1
    }

    return n * factorial(n-1)
}

func bernulli (n uint, k uint, p float64) float64 {
    nf := float64(n)
    kf := float64(k)
    facn := float64(factorial(n))
    fack := float64(factorial(k))
    facnk := float64(factorial(n-k))

    return (facn/(fack*facnk)) * math.Pow(p, kf) * math.Pow(1-p, nf-kf)
}

func main() {
    k := uint(6)
    p := 1.0/6.0
    n := uint(12)
    result := float64(0);
    for (k <= 12) {
        result += bernulli(n, k, p)
        k++
    }

    fmt.Println(result)
    fmt.Println(5 * (12 + 6 * 11 + 11 * 20 + 11 * 45 + 11 * 72 + 11 * 12 * 7) / math.Pow(6.0,12.0) )
}