## goqa
goqa is a simple terminal based quiz tool that can be used for example when learning 
foreign language words.
goqa reads questions and answers (e.g. local and foreign language words) from csv file 
given as argument, csv file format is:
```
question1,answer1
question2,answer2
.
.
questionn,answern
````
It then polls user randomly questions and prints real time score until user wants to quit.  

goqa is written in Go.

### Installation
Install and update with go: 
`go get -u github.com/jasilven/goqa`

### Usage
run: `$GOPATH/bin/goqa csv_file_containing_questions_and_answers.txt`

### Dependencies
No dependencies.
