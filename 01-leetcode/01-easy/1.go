// 1. Two Sum
package main

import "fmt"

func twoSum(nums [3]int, target int) []int {
   res := []int{}
   for i:=0; i < len(nums)-1; i++ {
      for j:=i+1; j < len(nums); j++ {
         if target - nums[i] == nums[j] {
            res = append(res, i);
            res = append(res, j);
         }
      }
   }
   return res
}


func main() {
   nums := [3]int{3,2,4};
   target := 6;
   res := twoSum(nums, target);
   fmt.Println(res[0], res[1]);
}
