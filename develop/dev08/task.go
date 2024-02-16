package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*/

func CommandExecute(args []string) error {
	switch args[0] {
	case "cd":
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Printf("Cannot change directory.\nCheck your current working directory with the help \"pwd\"")
		}

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current directory")
		}
		fmt.Printf("Current directory: %s\n", dir)

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "kill":
		for _, arg := range args[1:] {
			if arg[0] != '-' {
				_, err := strconv.Atoi(arg)
				if err != nil {
					return errors.New("wrong PID. To find PID you can type \"ps -e\"")
				}
			}
		}

		out, err := exec.Command("kill", args[1:]...).Output()
		if err != nil {
			return errors.New("could not kill process. Check your PID")
		} else {
			fmt.Printf("Process %s is killed!\n", string(out))
		}
	case "ps":
		process, err := exec.Command("ps", args[1:]...).Output()
		if err != nil {
			return err
		}
		fmt.Println(string(process))

	case "exec":
		if len(args[1:]) > 1 {
			fmt.Println("Too many arguments.")
			fmt.Println("Exec usage: exec firefox")
			return errors.New("so much arguments")
		}
		cmd := exec.Command("open", "-a", args[1])
		err := cmd.Run()
		if err != nil {
			fmt.Println("Wrong argument for exec.")
			return errors.New("wrong argument for exec")

		}
	//Fork не был реализован по причине отсутствия данной команды на macOS
	default:
		return errors.New("wrong command")
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		if args[0] == "quit" {
			return
		}

		fmt.Printf("%v\n", CommandExecute(args))

	}
}
