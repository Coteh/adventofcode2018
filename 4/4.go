package main

import (
	"fmt"
	"bufio"
	"io"
	"regexp"
	"strconv"
	"os"
	"strings"
	"sync"
	"log"
	"sort"
)

const NumMinutes = 60

type GuardInfo struct {
	guardID int
	month int
	day int
	minutes [60]int
}

type GuardMap struct {
	sync.RWMutex
	values map[string]*GuardInfo
}

type SleepWakeTimes struct {
	sleepMinutes []int
	wakeMinutes []int
}

type SleepWakeMap struct {
	sync.RWMutex
	values map[string]*SleepWakeTimes
}

func createGuardMap() *GuardMap {
	return &GuardMap {
		values: make(map[string]*GuardInfo),
	}
}

func createSleepWakeMap() *SleepWakeMap {
	return &SleepWakeMap {
		values: make(map[string]*SleepWakeTimes),
	}
}

func createSleepWakeTimes() *SleepWakeTimes {
	return &SleepWakeTimes {
		sleepMinutes: make([]int, 0, 30),
		wakeMinutes: make([]int, 0, 30),
	}
}

// func (this *GuardMap) GetEntry(key string) GuardInfo {
// 	this.RLock()
// 	defer this.RUnlock()

// 	return this.values[key]
// }

// func (this *GuardMap) SetMinutes(month int, day int, minutes []int) {
// 	this.Lock()
// 	defer this.Unlock()
	
// 	keyStr := createDayMonthString(month, day)
// 	guardInfo := this.values[keyStr]
// 	if guardInfo == (GuardInfo{}) {
// 		log.Fatal("Guard Info cannot be found for specified date")
// 	}

// 	copy(guardInfo.minutes[:], minutes)
// }

func (this *GuardMap) CreateEntry(month int, day int, guardID int) {
	this.Lock()
	defer this.Unlock()

	keyStr := createDayMonthString(month, day)
	this.values[keyStr] = &GuardInfo {
		guardID: guardID,
		month: month,
		day: day,
	}
}

func (this *SleepWakeMap) GetSleepWakeTimes(month int, day int) [60]int {
	var sortedSleepMinutes []int
	var sortedWakeMinutes []int
	var result [60]int

	keyStr := createDayMonthString(month, day)
	sleepItem := this.values[keyStr]
	if sleepItem == nil {
		return result
	}

	// Sort sleep minutes and wake minutes
	sortedSleepMinutes = make([]int, len(sleepItem.sleepMinutes))
	sortedWakeMinutes = make([]int, len(sleepItem.wakeMinutes))
	copy(sortedSleepMinutes, sleepItem.sleepMinutes)
	copy(sortedWakeMinutes, sleepItem.wakeMinutes)
	sort.Ints(sortedSleepMinutes)
	sort.Ints(sortedWakeMinutes)

	// Assume each sleep minute has corresponding wake minute
	for i, sleepMinute := range sortedSleepMinutes {
		wakeMinute := sortedWakeMinutes[i]
		for j := sleepMinute; j < wakeMinute; j++ {
			result[j] = 1
		}
	}

	return result
}

func (this *SleepWakeMap) AddSleepMinute(month int, day int, minute int) {
	this.Lock()
	defer this.Unlock()

	keyStr := createDayMonthString(month, day)
	if this.values[keyStr] == nil {
		this.values[keyStr] = createSleepWakeTimes()
	}

	sleepItem := this.values[keyStr]
	sleepItem.sleepMinutes = append(sleepItem.sleepMinutes, minute)
}

func (this *SleepWakeMap) AddWakeMinute(month int, day int, minute int) {
	this.Lock()
	defer this.Unlock()

	keyStr := createDayMonthString(month, day)
	if this.values[keyStr] == nil {
		this.values[keyStr] = createSleepWakeTimes()
	}

	sleepItem := this.values[keyStr]
	sleepItem.wakeMinutes = append(sleepItem.wakeMinutes, minute)
}

func (this *GuardMap) GetTotalAmountOfSleep() map[int]int {
	sleepMap := make(map[int]int)

	for _, entry := range this.values {
		for i := 0; i < NumMinutes; i++ {
			if entry.minutes[i] == 1 {
				sleepMap[entry.guardID] += 1
			}
		}
	}

	return sleepMap
}

func (this *GuardMap) GetMostAsleepGuard() int {
	sleepMap := this.GetTotalAmountOfSleep()
	var guardID int
	var highestSleep int = 0

	for id, sleep := range sleepMap {
		if sleep > highestSleep {
			highestSleep = sleep
			guardID = id
		}
	}

	return guardID
}

func (this *GuardMap) GetMostCommonSleepMinute(guardID int) int {
	// save count of each minute asleep to hash table
	minuteMap := make(map[int]int, 60)
	highestVal := 0
	highestMinute := 0
	
	// iterate through all Guard infos,
	// increment minuteMap value for corresponding minute
	// if it's a minute that guard was sleeping
	for _, guardInfo := range this.values {
		if guardInfo.guardID == guardID {
			for i, val := range guardInfo.minutes {
				if val == 1 {
					minuteMap[i] += 1
				}
			}
		}
	}

	// return the minute with highest count
	for i, value := range minuteMap {
		if value > highestVal {
			highestMinute = i
			highestVal = value
		}
	}

	return highestMinute
}

func createDayMonthString(month int, day int) string {
	var resultBuilder strings.Builder
	newMonth := strconv.FormatInt(int64(month), 10)
	newDay := strconv.FormatInt(int64(day), 10)

	resultBuilder.WriteString(newMonth)
	resultBuilder.WriteString("-")
	resultBuilder.WriteString(newDay)

	return resultBuilder.String()
}

func shiftDayMonth(month int, day int) (int, int) {
	day += 1

	switch month {
	case 1:
		if day > 31 {
			month = 2
			day = 1
		}
	case 2:
		if day > 28 {
			month = 3
			day = 1
		}
	case 3:
		if day > 30 {
			month = 4
			day = 1
		}
	case 4:
		if day > 30 {
			month = 5
			day = 1
		}
	case 5:
		if day > 31 {
			month = 6
			day = 1
		}
	case 6:
		if day > 30 {
			month = 7
			day = 1
		}
	case 7:
		if day > 31 {
			month = 8
			day = 1
		}
	case 8:
		if day > 31 {
			month = 9
			day = 1
		}
	case 9:
		if day > 30 {
			month = 10
			day = 1
		}
	case 10:
		if day > 31 {
			month = 11
			day = 1
		}
	case 11:
		if day > 30 {
			month = 12
			day = 1
		}
	default:
		// There are no dates in December
		// we don't need to worry about flipping years
	}

	return month, day
}

func populateSleepTimes(guardMap *GuardMap, sleepWakeMap *SleepWakeMap) {
	for _, guardInfo := range guardMap.values {
		month, day := guardInfo.month, guardInfo.day
		minutes := sleepWakeMap.GetSleepWakeTimes(month, day)
		copy(guardInfo.minutes[:], minutes[:])
	}
}

func parseTimestamp(input string, guardMap *GuardMap, sleepWakeMap *SleepWakeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	re, err := regexp.Compile("\\[([0-9]+)-([0-9]{2})-([0-9]{2})\\s([0-9]{2}):([0-9]{2})]\\s(.+)")
	if err != nil {
		return
	}

	matches := re.FindStringSubmatch(input)

	month, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal("There was an error getting month")
	}
	day, err := strconv.Atoi(matches[3])
	if err != nil {
		log.Fatal("There was an error getting day")
	}
	hour, err := strconv.Atoi(matches[4])
	if err != nil {
		log.Fatal("There was an error getting hour")
	}
	minute, err := strconv.Atoi(matches[5])
	if err != nil {
		log.Fatal("There was an error getting minute")
	}
	message := matches[6]

	if strings.Contains(message, "begins shift") {
		re, err := regexp.Compile("Guard #([0-9]+)\\sbegins shift")
		if err != nil {
			return
		}

		matches := re.FindStringSubmatch(message)

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatal("There was an error getting guard id")
		}

		if hour == 23 {
			month, day = shiftDayMonth(month, day)
		}

		guardMap.CreateEntry(month, day, id)
	} else {
		if message == "falls asleep" {
			sleepWakeMap.AddSleepMinute(month, day, minute)
		} else if message == "wakes up" {
			sleepWakeMap.AddWakeMinute(month, day, minute)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	guardMap := createGuardMap()
	sleepWakeMap := createSleepWakeMap()

	var wg sync.WaitGroup

	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal("Encountered an error with input")
				os.Exit(1)
			}
			break;
		}
		wg.Add(1)
		go parseTimestamp(input, guardMap, sleepWakeMap, &wg)
	}

	wg.Wait()

	populateSleepTimes(guardMap, sleepWakeMap)

	mostAsleepID := guardMap.GetMostAsleepGuard()
	mostCommonSleepMinute := guardMap.GetMostCommonSleepMinute(mostAsleepID)
	answer := mostAsleepID * mostCommonSleepMinute

	fmt.Println(answer)
}