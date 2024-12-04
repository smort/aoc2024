package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("input")
}

func part1(filename string) {
	lines := util.GetLines(filename)
	input := strings.Join(lines, "")
	lex := lexer{
		Pos:   0,
		Input: input,
	}

	result := lex.LexMult()
	fmt.Println(result)
}

func part2(filename string) {
	lines := util.GetLines(filename)
	input := strings.Join(lines, "")
	lex := lexer{
		Pos:               0,
		Input:             input,
		ParseConditionals: true,
		Enabled:           true,
	}

	result := lex.LexMult()
	fmt.Println(result)
}

var ErrEOF = errors.New("EOF")

type lexer struct {
	Pos               int
	Input             string
	ParseConditionals bool
	Enabled           bool
}

func (l *lexer) Inc() {
	l.Pos++
}

func (l *lexer) Curr() string {
	return l.Input[l.Pos:]
}

func (l *lexer) Next() (rune, error) {
	l.Pos++
	if l.Pos >= len(l.Input) {
		return 0, ErrEOF
	}
	char := l.Input[l.Pos]
	return rune(char), nil
}

func (l *lexer) SkipWhitespace() error {
	for {
		char, err := l.Next()
		if err != nil {
			return err
		}
		if !unicode.IsSpace(char) {
			l.Pos--
			break
		}
	}

	return nil
}

func (l *lexer) LexMult() int {
	result := 0
	var err error
	for {
		if errors.Is(err, ErrEOF) {
			break
		}

		var beginLexFunc lexBeginFn
		beginLexFunc, err = lexBeginning(l)
		if err != nil {
			continue
		}

		var isMore bool
		isMore, err = beginLexFunc(l)
		if err != nil || !isMore {
			continue
		}

		var num1 int
		num1, err = lexNumber(l)
		if err != nil {
			continue
		}

		err = lexComma(l)
		if err != nil {
			continue
		}

		var num2 int
		num2, err = lexNumber(l)
		if err != nil {
			continue
		}

		err = lexEndMul(l)
		if err != nil {
			continue
		}

		if !l.ParseConditionals || l.Enabled {
			result += num1 * num2
		}
	}

	return result
}

const (
	dont  = "don't()"
	do    = "do()"
	mul   = "mul("
	comma = ","
	end   = ")"
)

type lexBeginFn func(*lexer) (bool, error)

func lexBeginning(lex *lexer) (lexBeginFn, error) {
	err := lex.SkipWhitespace()
	if err != nil {
		return nil, err
	}
	for {
		curr := lex.Curr()

		if strings.HasPrefix(curr, mul) {
			return lexMulStart, nil
		} else if lex.ParseConditionals && strings.HasPrefix(curr, dont) {
			return lexDont, nil
		} else if lex.ParseConditionals && strings.HasPrefix(curr, do) {
			return lexDo, nil
		}

		lex.Inc()
		if reachedEOF(lex) != nil {
			return nil, ErrEOF
		}
	}
}

func lexDont(lex *lexer) (bool, error) {
	lex.Pos += len(dont)
	lex.Enabled = false
	return false, reachedEOF(lex)
}

func lexDo(lex *lexer) (bool, error) {
	lex.Pos += len(do)
	lex.Enabled = true
	return false, reachedEOF(lex)
}

func lexMulStart(lex *lexer) (bool, error) {
	lex.Pos += len(mul)
	return true, reachedEOF(lex)
}

func lexNumber(lex *lexer) (int, error) {
	var count int
	var num int
	for {
		if lex.Pos+count >= len(lex.Input) {
			return 0, ErrEOF
		}
		possibleNum, err := strconv.Atoi(lex.Input[lex.Pos : lex.Pos+count+1])
		if err != nil {
			break
		}
		num = possibleNum
		count++
	}

	if count == 0 {
		return 0, errors.New("no number")
	}

	lex.Pos += count

	return num, reachedEOF(lex)
}

func lexComma(lex *lexer) error {
	if string(lex.Input[lex.Pos]) != comma {
		return errors.New("not a comma")
	}
	lex.Pos += len(comma)

	return reachedEOF(lex)
}

func lexEndMul(lex *lexer) error {
	if string(lex.Input[lex.Pos]) != end {
		return errors.New("not a end paren")
	}

	lex.Pos += len(end)

	return nil // dont care about EOF at this point as we've reached the end of what we care to parse
}

func reachedEOF(lex *lexer) error {
	if lex.Pos >= len(lex.Input) {
		return ErrEOF
	}

	return nil
}
