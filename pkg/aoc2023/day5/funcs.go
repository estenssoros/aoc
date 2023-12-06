package day5

import (
	"bufio"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

type Almanac struct {
	Seeds []int
	Maps  map[string]*Map
}

func (a *Almanac) SeedDestinations() []int {
	out := make([]int, len(a.Seeds))
	bar := progressbar.New(len(a.Seeds))
	for i, seed := range a.Seeds {
		out[i] = a.seedDestination(seed)
		bar.Add(1)
	}
	bar.Close()
	return out
}

func (a *Almanac) SeedDestinationsAsync(numWorkers int) []int {
	bar := progressbar.New(len(a.Seeds))
	workChan := make(chan int)
	resultChan := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workChan <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
			defer wg.Done()
			for seed := range workChan {
				resultChan <- a.seedDestination(seed)
			}
		}(workChan, resultChan, &wg)
	}
	go func() {
		for _, seed := range a.Seeds {
			workChan <- seed
		}
		close(workChan)
	}()
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	out := []int{}
	for result := range resultChan {
		out = append(out, result)
		bar.Add(1)
	}
	bar.Close()
	return out
}

func (a *Almanac) seedDestination(seed int) int {
	mp := a.Maps["seed"]
	for mp.To != "location" {
		seed = mp.Next(seed)
		mp = a.Maps[mp.To]
	}
	seed = mp.Next(seed)
	return seed
}

type Map struct {
	Name   string
	To     string
	Ranges []*Range
}

func (m *Map) Next(seed int) (out int) {
	for _, rng := range m.Ranges {
		if rng.Has(seed) {
			return rng.Translate(seed)
		}
	}
	return seed
}

type Range struct {
	Destination int
	Source      int
	RangeLength int
}

func (r Range) String() string {
	ju, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(ju)
}

func (r *Range) Has(seed int) bool {
	return r.Source <= seed && seed <= r.Source+r.RangeLength-1
}

func (r *Range) Translate(seed int) int {
	return seed - r.Source + r.Destination
}

func parseAlmanac(input string) *Almanac {
	almanac := &Almanac{
		Maps: map[string]*Map{},
	}
	scanner := bufio.NewScanner(strings.NewReader(input))
	queue := NewQueue(scanner)
	line, err := queue.Pop()
	if err != nil {
		panic("queue.Pop")
	}
	almanac.Seeds = parseSeeds(line)
	_, err = queue.Pop()
	if err != nil {
		panic("queue.Pop")
	}
	var lines []string
	for {
		lines, err = queue.PopUntilBlank()
		if err != nil {
			break
		}

		mp := parseMap(lines)
		almanac.Maps[mp.Name] = mp
	}
	mp := parseMap(lines)
	almanac.Maps[mp.Name] = mp

	return almanac
}

var mapRegex = regexp.MustCompile(`([a-z]+)-to-([a-z]+) map:`)

func parseMap(lines []string) *Map {
	matches := mapRegex.FindStringSubmatch(lines[0])
	if len(matches) != 3 {
		panic("could not parse: " + lines[0])
	}
	return &Map{
		Name:   matches[1],
		To:     matches[2],
		Ranges: parseRanges(lines[1:]),
	}
}

func parseRanges(rangeStrings []string) []*Range {
	ranges := make([]*Range, len(rangeStrings))
	for i, rangeString := range rangeStrings {
		var dest, src, length int
		_, err := fmt.Sscanf(rangeString, "%d %d %d", &dest, &src, &length)
		if err != nil {
			panic(errors.Wrap(err, rangeString))
		}
		ranges[i] = &Range{dest, src, length}
	}
	return ranges
}

func parseSeeds(line string) []int {
	out := []int{}
	fields := strings.Split(line, ":")
	fields = strings.Split(fields[1], " ")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		i, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

type Queue struct {
	scanner *bufio.Scanner
}

func NewQueue(scanner *bufio.Scanner) Queue {
	return Queue{scanner: scanner}
}

func (q Queue) Pop() (string, error) {
	has := q.scanner.Scan()
	if !has {
		return "", errors.New("queue empty")
	}
	text := strings.TrimSpace(q.scanner.Text())

	return text, nil
}

func (q Queue) PopN(n int) ([]string, error) {
	out := make([]string, n)
	for i := 0; i < n; i++ {
		next, err := q.Pop()
		if err != nil {
			return nil, errors.Wrapf(err, "missing: %d", i)
		}
		out[i] = next
	}
	return out, nil
}

func (q Queue) PopUntilBlank() ([]string, error) {
	out := []string{}
	for {
		next, err := q.Pop()
		if err != nil {
			return out, err
		}
		if next == "" {
			return out, nil
		}
		out = append(out, next)
	}
}

func (q Queue) PopUntilMatch(match string) error {
	for {
		next, err := q.Pop()
		if err != nil {
			return err
		}
		if next == match {
			return nil
		}
	}
}
