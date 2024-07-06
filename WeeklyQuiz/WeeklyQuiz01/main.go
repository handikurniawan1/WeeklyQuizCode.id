package main

import (
	"fmt"
	"sort"
	"strings"
) 

func main() {
	
	fmt.Println(oddBeforeEven([]int{1,2,3,4,5,6,7,8,9,10})) //5
	fmt.Println(removeDuplicateLetter("bananas")) //1
	fmt.Println(removeDuplicate([]int{1,2,2,2,3,3,4,4,4})) //4
	fmt.Println(isAnagram("Otto", "Toto")) //2
	fmt.Println(capitalize([]string{"this", "is", "a", "kanoha"}, []string{"is", "a"})) //3
	fmt.Println(plusOne([]int{1, 2, 3, 4})) //6
	fmt.Println(sameElement([]int{1, 2, 4, 7, 8}, []int{2, 3, 7, 9})) //7
	fmt.Println(countFrequentElement([]int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4})) //9
	fmt.Println(addDigits([]int{1, 2, 3}, []int{4, 5, 6})) //10

	//8
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}
	fmt.Println("Same:", sliceFruits(fruits1, fruits2, 1))
	fmt.Println("Difference:", sliceFruits(fruits1, fruits2, 2))

	//11
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{2, 4, 6, 7}
	fmt.Println("Difference1:", sliceOperation(nums1, nums2, 1))
	fmt.Println("Difference2:", sliceOperation(nums1, nums2, 2))
	fmt.Println("Union:", sliceOperation(nums1, nums2, 3))
	fmt.Println("Intersection:", sliceOperation(nums1, nums2, 4))
}

func oddBeforeEven(nums []int)[]int{
	var odd, even []int

	for _, num := range nums{
		if num %2 == 0{
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	sort.Ints(even)
	sort.Ints(odd)

	result := append(odd, even...)
	return result
}

func removeDuplicateLetter(words string) string {
	seen := make(map[rune]bool)
	result := []rune{}

	for _, char := range words {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}

	return string(result)
}

func isAnagram(word1 string, word2 string) bool {
	w1 := strings.Split(strings.ToLower(word1), "")
	w2 := strings.Split(strings.ToLower(word2), "")

	if len(w1) != len(w2) {
		return false
	}

	sort.Strings(w1)
	sort.Strings(w2)

	return strings.Join(w1, "") == strings.Join(w2, "")
}

func removeDuplicate(nums []int)[]int{
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums{
		if !seen[num]{
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func capitalize(words []string, except []string) []string {
	exceptionMap := make(map[string]bool)
	for _, word := range except {
		exceptionMap[strings.ToLower(word)] = true
	}

	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if !exceptionMap[lowerWord] {
			words[i] = strings.Title(lowerWord)
		}
	}

	return words
}

func plusOne(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	return append([]int{1}, nums...)
}

func sameElement(nums1 []int, nums2 []int) []int {
	elements := make(map[int]bool)
	for _, num := range nums2 {
		elements[num] = true
	}

	var result []int
	for _, num := range nums1 {
		if elements[num] {
			result = append(result, num)
		}
	}
	return result
}

func countFrequentElement(nums []int) map[int]int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	return frequency
}

func addDigits(nums1 []int, nums2 []int) []int {

	result := make([]int, len(nums1))
	for i := range nums1 {
		result[i] = nums1[i] + nums2[i]
	}
	return result
}

func sliceFruits(fruits1, fruits2 []string, operationType int) []string {
	set := func(fruits []string) map[string]bool {
		s := make(map[string]bool)
		for _, fruit := range fruits {
			s[fruit] = true
		}
		return s
	}

	set1 := set(fruits1)
	set2 := set(fruits2)
	result := []string{}

	switch operationType {
	case 1: 
		for fruit := range set1 {
			if set2[fruit] {
				result = append(result, fruit)
			}
		}
	case 2: 
		for fruit := range set1 {
			if !set2[fruit] {
				result = append(result, fruit)
			}
		}
		for fruit := range set2 {
			if !set1[fruit] {
				result = append(result, fruit)
			}
		}
	}
	return result
}

func sliceOperation(nums1, nums2 []int, typeOperation int) []int {
	set := func(nums []int) map[int]bool {
		s := make(map[int]bool)
		for _, num := range nums {
			s[num] = true
		}
		return s
	}

	set1 := set(nums1)
	set2 := set(nums2)
	result := []int{}

	switch typeOperation {
	case 1: 
		for _, num := range nums1 {
			if !set2[num] {
				result = append(result, num)
			}
		}
	case 2: 
		for _, num := range nums2 {
			if !set1[num] {
				result = append(result, num)
			}
		}
	case 3: 
		for num := range set1 {
			result = append(result, num)
		}
		for num := range set2 {
			if !set1[num] {
				result = append(result, num)
			}
		}
	case 4: 
		for num := range set1 {
			if set2[num] {
				result = append(result, num)
			}
		}
	}
	return result
}

