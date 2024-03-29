// 13. Roman to Integer
package main

import "fmt"

func match(c byte) int {
   if c == 'I' {
      return 1
   }
   if c == 'V' {
      return 5
   }
   if c == 'X' {
      return 10
   }
   if c == 'L' {
      return 50
   }
   if c == 'C' {
      return 100
   }
   if c == 'D' {
      return 500
   }
   if c == 'M' {
      return 1000
   }
   return 0
}

func romanToInt(s string) int {
   res := 0
   for i := 0; i < len(s)-1; i++ {
      j := match(s[i]);
      k := match(s[i+1]);
      if j < k {
         res -= j;
      } else {
         res += j;
      }
   }
   res += match(s[len(s)-1])
   return res
}

func main() {
   s := "MCMXCIV";
   fmt.Println(romanToInt(s));
}
