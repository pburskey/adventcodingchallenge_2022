package main

import (
	"adventcodingchallenge_2022/utility"
	"math"
	"strconv"
	"strings"
)

type Part1 struct {
	answer int
}

type Item struct {
	levelOfWorry int
}

type MonkeyContainer struct {
	monkeys []*Monkey
}

func (c *MonkeyContainer) getMonkeyHavingId(id int) *Monkey {
	var interestedMonkey *Monkey
	for _, aMonkey := range c.monkeys {
		if aMonkey.id == id {
			interestedMonkey = aMonkey
			return interestedMonkey
		}
	}
	return interestedMonkey
}

type Operation struct {
	raw      string
	operator string
	operands []interface{}
}

func (o *Operation) process(item *Item) int {
	if o.operator == "" {
		munge := strings.ReplaceAll(o.raw, "new = ", "")
		words := strings.Split(munge, " ")
		o.operator = words[1]
		o.operands = make([]interface{}, 2)
		o.operands[0] = words[0]
		o.operands[1] = words[2]
	}

	a := 0
	b := 0
	if o.operands[0] == "old" {
		a = item.levelOfWorry
	} else {
		aValue, _ := strconv.Atoi(o.operands[0].(string))
		a = aValue
	}
	if o.operands[1] == "old" {
		b = item.levelOfWorry
	} else {
		aValue, _ := strconv.Atoi(o.operands[1].(string))
		b = aValue
	}
	newValue := 0

	if o.operator == "*" {
		newValue = a * b
	} else if o.operator == "+" {
		newValue = a + b
	}
	return newValue

}

type Test struct {
	testValue              int
	monkeyRecipientIfTrue  int
	monkeyRecipientIfFalse int
}

type Monkey struct {
	id              int
	items           []*Item
	operation       *Operation
	test            *Test
	inspectionCount int
	considerRelief  bool
}

func (m *Monkey) inspectItem(index int, item *Item) {
	if m.operation != nil {
		newWorryLevel := m.operation.process(item)
		item.levelOfWorry = newWorryLevel
		m.inspectionCount++
	}
}

func (m *Monkey) testWorryLevel(index int, item *Item) int {
	if m.test != nil {
		divisible := item.levelOfWorry % m.test.testValue
		if divisible == 0 {
			return m.test.monkeyRecipientIfTrue
		} else {
			return m.test.monkeyRecipientIfFalse
		}
	}
	return -1
}

func (m *Monkey) considerItem(container *MonkeyContainer, index int, item *Item) {
	m.inspectItem(index, item)

	/*
		After each monkey inspects an item but before it tests your worry level, your relief that the monkey's inspection didn't damage the item causes your worry level to be divided by three and rounded down to the nearest integer.
	*/
	if m.considerRelief {
		var relief float64
		relief = float64((item.levelOfWorry / 3))
		relief = math.Floor(relief)
		item.levelOfWorry = int(relief)
	}

	nextMonkeyId := m.testWorryLevel(index, item)
	nextMonkey := container.getMonkeyHavingId(nextMonkeyId)

	m.items[index] = nil
	nextMonkey.items = append(nextMonkey.items, item)

}

func (m *Monkey) conductBusiness(container *MonkeyContainer) {
	if m.items != nil && len(m.items) > 0 {
		for index, anItem := range m.items {
			m.considerItem(container, index, anItem)
		}
	}

	if m.items != nil {
		itemsToCleanup := m.items
		m.items = make([]*Item, 0)
		for _, item := range itemsToCleanup {
			if item != nil {
				m.items = append(m.items, item)
			}
		}
	}

}

func (m *Monkey) throwToMonkey(monkeys *MonkeyContainer, monkeyNumber int) {
}

func parseCommands(data []string, considerRelief bool) *MonkeyContainer {

	container := &MonkeyContainer{monkeys: make([]*Monkey, 0)}
	var monkey *Monkey
	for _, aRow := range data {

		if strings.Index(aRow, "Monkey ") > -1 {
			words := strings.Split(aRow, " ")
			aString := words[1]
			aString = strings.Replace(aString, ":", "", 1)
			anID, _ := strconv.Atoi(aString)
			monkey = &Monkey{
				id:             anID,
				items:          make([]*Item, 0),
				operation:      nil,
				test:           nil,
				considerRelief: considerRelief,
			}
			container.monkeys = append(container.monkeys, monkey)
		} else if strings.Index(aRow, "Starting items:") > -1 {
			aRow = strings.Replace(aRow, "Starting items:", "", 1)
			words := strings.Split(aRow, ",")
			for _, aWord := range words {
				aNumber, _ := strconv.Atoi(strings.ReplaceAll(aWord, " ", ""))
				monkey.items = append(monkey.items, &Item{aNumber})
			}
		} else if strings.Index(aRow, "Operation: ") > -1 {
			munge := strings.ReplaceAll(aRow, "Operation: ", "")
			munge = strings.TrimLeft(munge, " ")
			monkey.operation = &Operation{raw: munge}

		} else if strings.Index(aRow, "Test: ") > -1 {
			munge := strings.ReplaceAll(aRow, "Test: ", "")
			munge = strings.TrimLeft(munge, " ")
			munge = strings.ReplaceAll(munge, "divisible by ", "")
			aValue, _ := strconv.Atoi(munge)
			monkey.test = &Test{testValue: aValue}

		} else if strings.Index(aRow, "If true: ") > -1 {
			munge := strings.ReplaceAll(aRow, "If true: throw to monkey", "")
			munge = strings.TrimLeft(munge, " ")
			aValue, _ := strconv.Atoi(munge)
			monkey.test.monkeyRecipientIfTrue = aValue
		} else if strings.Index(aRow, "If false: ") > -1 {
			munge := strings.ReplaceAll(aRow, "If false: throw to monkey", "")
			munge = strings.TrimLeft(munge, " ")
			aValue, _ := strconv.Atoi(munge)
			monkey.test.monkeyRecipientIfFalse = aValue
		}
	}
	return container
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	container := parseCommands(data, true)
	if container != nil {
		for round := 0; round < 20; round++ {
			for _, monkey := range container.monkeys {
				monkey.conductBusiness(container)
			}
		}

		/*
			find two most active monkeys
		*/
		numbers := make([]int, 0)
		for _, monkey := range container.monkeys {
			aValue := monkey.inspectionCount
			numbers = append(numbers, aValue)
		}
		numbers = utility.OrderNumbersSortReversed(numbers)
		alg.answer = numbers[0] * numbers[1]

	}
	return nil, alg.answer
}
