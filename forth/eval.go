//go:build !solution

package main

import (
	"errors"
	"strconv"
	"strings"
)

// support all numbers - done
// support all arithm operations - done
// support all commands - done
// support custom defined commands creation - evaluate when adding
// support custom defined commands usage - done
// support no default cmd redefine after usage - done

type Command func(e *Evaluator) error

var defaultCommands = map[string]Command{
	"+": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid operation: stack length is less than 2")
		}

		a := e.stack[len(e.stack)-2]
		b := e.stack[len(e.stack)-1]

		e.stack = e.stack[:len(e.stack)-2]
		e.stack = append(e.stack, a+b)
		return nil
	},

	"-": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid operation: stack length is less than 2")
		}

		a := e.stack[len(e.stack)-2]
		b := e.stack[len(e.stack)-1]

		e.stack = e.stack[:len(e.stack)-2]
		e.stack = append(e.stack, a-b)
		return nil
	},

	"*": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid operation: stack length is less than 2")
		}

		a := e.stack[len(e.stack)-2]
		b := e.stack[len(e.stack)-1]

		e.stack = e.stack[:len(e.stack)-2]
		e.stack = append(e.stack, a*b)
		return nil
	},

	"/": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid operation: stack length is less than 2")
		}

		a := e.stack[len(e.stack)-2]
		b := e.stack[len(e.stack)-1]

		if b == 0 {
			return errors.New("invalid operation: division by zero")
		}

		e.stack = e.stack[:len(e.stack)-2]
		e.stack = append(e.stack, a/b)
		return nil
	},

	"dup": func(e *Evaluator) error {
		if len(e.stack) == 0 {
			return errors.New("invalid dup operation: stack is empty")
		}

		num := e.stack[len(e.stack)-1]
		e.stack = append(e.stack, num)

		return nil
	},

	"over": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid over operation: stack is less than 2")
		}

		num := e.stack[len(e.stack)-2]
		e.stack = append(e.stack, num)

		return nil
	},

	"drop": func(e *Evaluator) error {
		if len(e.stack) == 0 {
			return errors.New("invalid drop operation: stack is empty")
		}

		e.stack = e.stack[:len(e.stack)-1]
		return nil
	},

	"swap": func(e *Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("invalid swap operation: stack is less than 2")
		}

		e.stack[len(e.stack)-1], e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2], e.stack[len(e.stack)-1]
		return nil
	},
}

type Evaluator struct {
	stack        []int
	custom       map[string][]string
	usedDefaults map[string]bool
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	usedDefaults := make(map[string]bool)

	for k := range defaultCommands {
		usedDefaults[k] = false
	}

	return &Evaluator{custom: make(map[string][]string), usedDefaults: usedDefaults}
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	tokens := strings.Fields(row)

	if tokens[0] == ":" {
		err := e.recordUserCommand(tokens)
		return nil, err
	}

	err := e.processExpression(tokens)
	if err != nil {
		return nil, err
	}

	return e.stack, nil
}

func (e *Evaluator) recordUserCommand(tokens []string) error {
	start := 0

	if start+1 >= len(tokens) || tokens[start+1] == ";" {
		return errors.New("custom command error: no command name")
	}
	customName := strings.ToLower(tokens[start+1])

	// no default command overwrite if they were used
	if e.usedDefaults[customName] {
		return nil
	}

	if _, err := strconv.Atoi(customName); err == nil {
		return errors.New("custom command error: cannot redefine number")
	}

	end := start
	for end < len(tokens) && tokens[end] != ";" {
		end++
	}

	if end >= len(tokens) {
		return errors.New("custom command error: unclosed command definition")
	}

	if start+2 >= len(tokens) {
		return errors.New("custom command error: no custom command payload found")
	}

	payload := tokens[start+2 : end]
	payloadCopy := make([]string, len(payload))
	copy(payloadCopy, payload)

	evaluatedPayload, err := e.evaluateUserCommandPayload(payloadCopy)
	if err != nil {
		return err
	}

	e.custom[customName] = evaluatedPayload
	return nil
}

func (e *Evaluator) evaluateUserCommandPayload(tokens []string) ([]string, error) {
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if customCmd, ok := e.custom[t]; ok {
			tokens = append(tokens[:i], append(customCmd, tokens[i+1:]...)...)
			i--
			continue
		}

		if _, err := strconv.Atoi(t); err == nil {
			continue
		}

		if _, ok := defaultCommands[strings.ToLower(t)]; ok {
			e.usedDefaults[strings.ToLower(t)] = true
			continue
		}

		return nil, errors.New("custom command error: unknown token")
	}

	return tokens, nil
}

func (e *Evaluator) processExpression(tokens []string) error {
	for i := 0; i < len(tokens); i++ {
		tLow := strings.ToLower(tokens[i])

		if customCmd, ok := e.custom[tLow]; ok {
			tokens = append(tokens[:i], append(customCmd, tokens[i+1:]...)...)
			i--
			continue
		}

		if num, err := strconv.Atoi(tLow); err == nil {
			e.stack = append(e.stack, num)
			continue
		}

		if cmd, ok := defaultCommands[tLow]; ok {
			err := cmd(e)
			if err != nil {
				return err
			}
			continue
		}

		return errors.New("invalid token")
	}

	return nil
}
