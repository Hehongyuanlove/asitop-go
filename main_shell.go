package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main_shell() {
	cmd := exec.Command("sh", "-c", "cat") // 使用 `sh -c` 来执行一个 shell 命令

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe: ", err)
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe: ", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // 打印从命令输出的数据
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading stdout: ", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1] // 去掉换行符

		if text == "exit" {
			break
		}

		_, err := stdin.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Println("Error writing to stdin: ", err)
			break
		}
	}

	stdin.Close()
	cmd.Wait()
	fmt.Println("Command finished.")
}
