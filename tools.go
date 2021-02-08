package main

import (
	"fmt"
	"github.com/valter4578/vlang/evaluator"
	"github.com/valter4578/vlang/lexer"
	"github.com/valter4578/vlang/object"
	"github.com/valter4578/vlang/parser"
	"io/ioutil"
	"log"
)

func parseFile(fileName string) string {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error: Couldn't read file %v", fileName)
	}

	return string(content)
}

func evalFile(fileData string) {
	l := lexer.New(fileData)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println(p.Errors())
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}
