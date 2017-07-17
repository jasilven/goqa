package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type qa struct {
	answer   string
	question string
}

func loadQA(fn string) (error, []qa) {
	file, err := os.Open(fn)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	retval := []qa{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x qa
		singleQA := strings.Split(scanner.Text(), ",")
		x.question = singleQA[0]
		x.answer = singleQA[1]
		retval = append(retval, x)
	}
	return nil, retval
}

func contains(s string, list []string) bool {
	for _, val := range list {
		if strings.Compare(s, val) == 0 {
			return true
		}
	}
	return false
}

func getQuestionData(count int, qas []qa) (string, string, []string) {
	rnd := rand.Intn(len(qas) - 1)
	question := qas[rnd].question
	answer := qas[rnd].answer
	cands := []string{answer}

	for len(cands) < count {
		i := rand.Intn(len(qas) - 1)
		if !contains(qas[i].answer, cands) {
			cands = append(cands, qas[i].answer)
		}
	}
	sort.Strings(cands)
	return question, answer, cands
}

func printScore(ncorrect int, total int, str string) {
	fmt.Printf("SCORE:[%s] %d/%d (%.0f%s)\n", str, ncorrect, total, 100*float32(ncorrect)/float32(total), "%")
}

func promptAnswer(q string, cands []string, round int) (bool, string) {
	fmt.Println("Question", round)
	for key, val := range cands {
		fmt.Printf("\t\t%d) %s\n", key+1, val)
	}
	for {
		input := 0
		fmt.Print("(0 to exit) ", q, "? ")
		_, err := fmt.Scan(&input)
		if (err != nil) || (input > len(cands)) || (input < 0) {
			continue
		}
		if input == 0 {
			return true, ""
		}
		return false, cands[input-1]
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage:", os.Args[0], "<file>")
	}
	qas := make([]qa, 1)
	err, qas := loadQA(os.Args[1])
	if err != nil {
		log.Fatalln("cannot load data from", os.Args[1], err)
	}

	rounds, correct := 0, 0
	scoreStr := ""
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		rounds++
		question, answer, cands := getQuestionData(4, qas)
		printScore(correct, rounds, scoreStr)
		quit, ans := promptAnswer(question, cands, rounds)
		if quit {
			break
		}
		if strings.Compare(answer, ans) == 0 {
			correct++
			scoreStr = scoreStr + "+"
		} else {
			scoreStr = scoreStr + "-"
		}
	}
	printScore(correct, rounds, scoreStr)
}
