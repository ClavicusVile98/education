package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	//common "common"
)

func ExecutePipeline(hashSignJobs ...job){
	/* инициализация WaitGroup,
	 * не надо копировать,
	 * поэтому используется указатель*/
	wg := &sync.WaitGroup{}

	/* считывание данных в буферизированный канал */
	workerInput := make(chan interface{}, MaxInputDataLen) /* неэкспортируемая переменная */
	/* вывести данные в буферизированный канал */
	workerOutput := make(chan interface{}, MaxInputDataLen) /* неэкспортируемая переменная */

	for _, myjob := range hashSignJobs {
		wg.Add(1)
		/* запуск горутины */
		go func(workerInput, workerOutput chan interface{}, jfunc job){
			/* ждем, когда воркер закончит работу */
			defer wg.Done()
			/* ждем закрытия канала */
			defer close(workerOutput)
			jfunc(workerInput, workerOutput)
		}(workerInput, workerOutput, myjob)
		workerInput = workerOutput
	}
	/* ожидаем, пока не отработает горутина*/
	wg.Wait()
}

func SingleHash(workerInput, workerOutput chan interface{}) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for item := range workerInput {
		digit := strconv.Itoa(item.(int))
		//data := fmt.Sprint(digit)
		wg.Add(1)

		/* горутина */
		// workerOutput chan interface{}, wg *sync.WaitGroup, mu *sync.Mutex,
		go func(digit string) {
			defer wg.Done()
			var CRC32 string // объявление переменных
			var CRC32MD5 string
			wgCRC := &sync.WaitGroup{}
			wgCRC.Add(1)

			/* еще одна горутина */
			go func() {
				defer wgCRC.Done()
				mu.Lock() // сначала блокируем, чтобы другие не стучались
				MD5 := DataSignerMd5(digit)
				mu.Unlock() // затем разблокировали, чтобы можно было работать дальше
				CRC32MD5 = DataSignerCrc32(MD5)
			}()

			CRC32 = DataSignerCrc32(digit)
			wgCRC.Wait()
			result := CRC32+"~"+CRC32MD5
			workerOutput <- result
		}(digit)
	}
	wg.Wait()
}

func MultiHash(workerInput, workerOutput chan interface{}) {
	wg := &sync.WaitGroup{}

	for item := range workerInput {
		data := item.(string)
		wg.Add(1)

		// out chan interface{}, wg *sync.WaitGroup,
		go func(data string) {
			defer wg.Done()
			result := make([]string, 6)
			wgCRC := &sync.WaitGroup{}

			for th := 0; th < 6; th+=1 {
				wgCRC.Add(1)

				go func(th int) {
					defer wgCRC.Done()
					CRC32 := DataSignerCrc32(fmt.Sprint(th)+data)
					result[th] = CRC32
				}(th)
			}

			wgCRC.Wait()
			workerOutput <- strings.Join(result, "")
		}(data)
	}

	wg.Wait()
}

func CombineResults(workerInput, workerOutput chan interface{}) {
	var result []string

	for item := range workerInput {
		data := item.(string)
		result = append(result, data)
	}
	sort.Strings(result)
	workerOutput <- strings.Join(result, "_")
}

func main(){}
