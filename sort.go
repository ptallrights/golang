package main

import (
	"fmt"
)

var globalArray = []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}

/*
冒泡排序(bubbleSort)
1、对比相邻的两个元素：如果第一个比第二个大，就交换它们两个
2、对每一对相邻的两个元素进行同样的比较：从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数
3、针对所有的元素重复以上的步骤，除了最后一个
4、重复步骤1~3，直到排序完成
*/
func bubbleSort(array []int) []int {
	ln := len(array)
	for i := 0; i < ln - 1; i++ {
		for j := 0; j < ln - 1 - i; j++ {
			if array[j] > array[j + 1] {
				tmpElement := array[j + 1]
				array[j + 1] = array[j]
				array[j] = tmpElement
			}
		}
	}
	return array
}

/*
选择排序(selectionSort)
1、首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
2、再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
3、以此类推，直到所有元素均排序完毕。
PS：n个记录的直接选择排序可经过n-1趟直接选择排序得到有序结果。具体算法描述如下：
	1、初始状态：无序区为R[1…n]，有序区为空；
	2、第i趟排序(i=1,2,3…n-1)开始时，当前有序区和无序区分别为R[1…i-1]和R(i…n）
	3、该趟排序从当前无序区中-选出关键字最小的记录R[k]，将它与无序区的第1个记录R交换，使R[1…i]和R[i+1…n)分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区
	4、n-1趟结束，数组有序化了。
*/
func selectionSort(array []int) []int {
	ln := len(array)
	for i := 0; i < ln -1; i++ {
		minIndex := i
		for j := minIndex; j < ln; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		tmpElement := array[i]
		array[i] = array[minIndex]
		array[minIndex] = tmpElement
	}
	return array
}

/*
插入排序(insertionSort)
1、从第一个元素开始，该元素可以认为已经被排序
2、取出下一个元素，在已经排序的元素序列中从后向前扫描
3、如果该元素（已排序）大于新元素，将该元素移到下一位置
4、重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
5、将新元素插入到该位置
6、重复步骤2~5。
*/
func insertionSort(array []int) []int {
	ln := len(array)
	for i := 1; i < ln; i++ {
		preIndex := i - 1
		currentElement := array[i]
		for (preIndex >= 0 && array[preIndex] > currentElement) {
			array[preIndex + 1] = array[preIndex]
			preIndex--
		}
		array[preIndex + 1] = currentElement
	}
	return array
}

/*
希尔排序(shellSort)
1、选择增量 ：gap=length/2，缩小增量：gap = gap/2
2、增量序列：用序列表示增量选择，{n/2, (n/2)/2, …, 1}
3、希尔排序的增量序列的选择与证明是个数学难题，我们选择的这个增量序列是比较常用的，也是希尔建议的增量，称为希尔增量，但其实这个增量序列不是最优的。此处我们做示例使用希尔增量
4、先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：
	1、选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
	2、按增量序列个数k，对序列进行k趟排序；
	3、每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序;
	4、仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func shellSort(array []int) []int {
	ln := len(array)
	gap := ln / 2
	for ; gap > 0; gap /= 2 {
		for i := gap; i < ln; i++ {
			for j := i - gap; j >= 0 && array[j + gap] < array[j]; j -= gap {
				array[j], array[j + gap] = array[j + gap], array[j]
			}
		}
	}
	return array
}

/*
归并排序(mergeSort)
1、把长度为n的输入序列分成两个长度为n/2的子序列
2、对这两个子序列分别采用归并排序
3、将两个排序好的子序列合并成一个最终的排序序列
*/
func mergeSort(array []int) []int {
	ln := len(array)
	if ln <= 1 {
		return array
	}
	num := ln / 2
	left := mergeSort(array[:num])
	right := mergeSort(array[num:])
	return merge(left, right)
}

func merge(left, right []int)  (result []int) {
	r, l := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/*
快速排序(quickSort)
1、从数列中挑出一个元素，称为 “基准”（pivot），即枢纽元
2、重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边），在这个分区退出之后，该基准就处于数列的中间位置，称为分区（partition）操作
3、递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序
*/
func quickSort(array []int) (result []int) {
	result = array
	ln := len(result)
	if ln <= 1 {
		return
	}
	mid, i := result[0], 1
	head, tail := 0, len(result) - 1
	for head < tail {
		if result[i] > mid {
			result[i], result[tail] = result[tail], result[i]
			tail--
		} else {
			result[i], result[head] = result[head], result[i]
			head++
			i++
		}
	}
	result[head] = mid
	quickSort(result[:head])
	quickSort(result[head+1:])
	return
}

/*
堆排序(heapSort)
1、将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区
2、将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]
3、由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)
4、不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
*/
func heapSort(array []int) []int{
	for i := len(array) / 2; i >= 0; i -- {
		heap(array, i)
	}
	end := len(array) - 1
	for end > 0 {
		array[0], array[end] = array[end], array[0]
		heap(array[:end], 0)
		end--
	}
	return array
}

func heap(array []int, pos int) {
	end := len(array) - 1
	left := 2 * pos + 1
	if left > end {
		return
	}

	right := 2 * pos + 2
	temp := left
	if right <= end && array[right] > array[temp] {
		temp = right
	}
	if array[temp] <= array[pos] {
		return
	}
	array[temp], array[pos] = array[pos], array[temp]
	heap(array, temp)
}

/*
计数排序(countingSort)
1、找出待排序的数组中最大和最小的元素
2、统计数组中每个值为i的元素出现的次数，存入数组C的第i项
3、对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
4、反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
 */
func countingSort(array []int) []int{
	var maxValue = 0
	ln := len(array)
	for i := 0; i < ln - 1; i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	for i := 0; i < ln; i++ {
		bucket[array[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			array[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return array
}

/*
桶排序(bucketSort)
桶排序 (Bucket sort)或所谓的箱排序的原理是将数组分到有限数量的桶子里，然后对每个桶子再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序），最后将各个桶中的数据有序的合并起来。
 */
func bucketSort(array []int) []int{
	ln := len(array)
	if ln == 0 {
		return array
	}

	var minValue, maxValue = array[0], array[0]
	for i := 1; i < ln; i++ {
		if array[i] < minValue  {
			minValue = array[i]
		} else if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	var defaultBucketSize = 5
	bucketCount := (maxValue - minValue) / defaultBucketSize + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < ln; i++ {
		buckets[(array[i] - minValue) / defaultBucketSize] = append(buckets[(array[i] - minValue) / defaultBucketSize], array[i])
	}

	key := 0
	for _, bucket := range buckets {
		if len(bucket) < 0 {
			continue
		}
		insertionSort(bucket)
		for _, value := range bucket {
			array[key] = value
			key += 1
		}
	}

	return array
}

/*
基数排序(radixSort)
1、取得数组中的最大数，并取得位数
2、arr为原始数组，从最低位开始取每个位组成radix数组
3、对radix进行计数排序（利用计数排序适用于小范围数的特点）
 */
func radixSort(array []int) []int{
	var maxValue = 0
	for _, value := range array {
		if value > maxValue {
			maxValue = value
		}
	}
	var maxDigit int = 0
	for maxValue % 10 > 0 {
		maxValue = maxValue / 10
		maxDigit++
	}

	var mod, dev = 10, 1
	for i := 0; i < maxDigit; i++ {
		counter := make([][]int, 10)
		for j := 0; j < len(counter); j++ {
			bucket := (array[j] % mod) / dev
			counter[bucket] = append(counter[bucket], array[j])
		}
		pos := 0
		for j := 0; j < len(counter); j++ {
			if counter[j] != nil {
				for k := 0; k < len(counter[j]); k++ {
					array[pos] = counter[j][k]
					pos++
				}
			}
		}
		mod *= 10
		dev *= 10
	}
	return array
}

func main() {
	fmt.Println("initial array:", globalArray)
	// bubble sort
	bubbleSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(bubbleSortArray)
	bubbleSortResult := bubbleSort(bubbleSortArray)
	fmt.Println("bubble sort:", bubbleSortResult)

	//selection sort
	selectionSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(selectionSortArray)
	selectionSortResult := selectionSort(selectionSortArray)
	fmt.Println("selection sort:", selectionSortResult)

	//insertion sort
	insertionSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(insertionSortArray)
	insertionSortResult := insertionSort(insertionSortArray)
	fmt.Println("insertion sort:", insertionSortResult)

	//shell sort
	shellSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(shellSortArray)
	shellSortResult := shellSort(shellSortArray)
	fmt.Println("shell sort:", shellSortResult)

	//merge sort
	mergeSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(mergeSortArray)
	mergeSortResult := mergeSort(mergeSortArray)
	fmt.Println("merge sort:", mergeSortResult)

	//quick sort
	quickSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(quickSortArray)
	quickSortResult := quickSort(quickSortArray)
	fmt.Println("quick sort:", quickSortResult)

	//heap sort
	heapSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(heapSortArray)
	heapSortResult := heapSort(heapSortArray)
	fmt.Println("heap sort:", heapSortResult)

	//counting sort
	countingSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(countingSortArray)
	countingSortResult := countingSort(countingSortArray)
	fmt.Println("counting sort:", countingSortResult)

	//bucket sort
	bucketSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(bucketSortArray)
	bucketSortResult := bucketSort(bucketSortArray)
	fmt.Println("bucket sort:", bucketSortResult)

	//radix sort
	radixSortArray := []int{3, 2, 4, 3, 10, 14, 11, 9, 7, 8}
	fmt.Println(radixSortArray)
	radixSortResult := radixSort(radixSortArray)
	fmt.Println("radix sort:", radixSortResult)
}