package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func ExecutePipeline(work ...job) {
	/* инициализация WaitGroup,
	 * не надо копировать,
	 * поэтому используется указатель*/
	wg := &sync.WaitGroup{}

	/* считывание данных в буферизированный канал */
	workerInput := make(chan interface{}, MaxInputDataLen) /* неэкспортируемая переменная */
	/* вывести данные в буферизированный канал */
	workerOutput := make(chan interface{}, MaxInputDataLen) /* неэкспортируемая переменная */

	for _, myjob := range work {
		wg.Add(1)
		/* запуск горутины */
		go func(workerInput, workerOutput chan interface{}, j job){
			/* ждем, когда воркер закончит работу */
			defer wg.Done()
			/* ждем закрытия канала */
			defer close(workerOutput)
			j(workerInput, workerOutput)
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
		digit, err := item.(int)
		if !err {
			panic("Cannot convert data to int")
		}

		data := fmt.Sprint(digit)
		wg.Add(1)
		go func(workerOutput chan interface{}, wg *sync.WaitGroup, mu *sync.Mutex, data string) {
			defer wg.Done()
			var CRC32 string
			var CRC32MD5 string
			wgCRC := &sync.WaitGroup{}
			wgCRC.Add(1)
			go func() {
				defer wgCRC.Done()
				mu.Lock() // сначала блокируем, чтобы другие не стучались
				MD5 := DataSignerMd5(data)
				mu.Unlock() // затем разблокировали, чтобы можно было работать дальше
				CRC32MD5 = DataSignerCrc32(MD5)
			}()

			CRC32 = DataSignerCrc32(data)
			wgCRC.Wait()
			result := CRC32+"~"+CRC32MD5
			workerOutput <- result
		}(workerOutput, wg, mu, data)
	}
	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for item := range in {
		data, ok := item.(string)
		if !ok {
			panic("cannot convert data to string")
		}

		wg.Add(1)
		go func(out chan interface{}, wg *sync.WaitGroup, data string) {
			defer wg.Done()
			result := make([]string, 6, 6)
			wgCRC := &sync.WaitGroup{}

			for th := 0; th <= 5; th++ {
				wgCRC.Add(1)

				go func(th int) {
					defer wgCRC.Done()
					CRC32 := DataSignerCrc32(fmt.Sprint(th)+data)
					result[th] = CRC32
				}(th)
			}

			wgCRC.Wait()
			out <- strings.Join(result, "")
		}(out, wg, data)
	}

	wg.Wait()
}

func CombineResults(workerInput, workerOutput chan interface{}) {
	var result []string

	for item := range workerInput {
		data, err := item.(string)
		if !err {
			panic(err)
		}
		result = append(result, data)
	}
	sort.Strings(result)
	workerOutput <- strings.Join(result, "_")
}

func main(){}