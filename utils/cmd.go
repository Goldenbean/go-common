package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Execute : run BASH Command
func Execute(command string) (string, bool) {

	//fmt.Println("Commands: ", command)
	cmd := exec.Command("/bin/bash", "-c", command)
	// cmd := exec.Command("/bin/bash", command)

	output, err := cmd.Output()

	if err != nil {
		//fmt.Printf("Execute Shell: '[%s]', failed with error: %s\n", command, err.Error())
		return fmt.Sprintf("Error [%s], Output [%s], Commands [%s]",
			err.Error(), string(output), command), false
	}

	ret := string(output)
	//fmt.Printf("Execute Shell: '[%s]', finished with output: %s\n", command, ret)
	return ret, true
}

// RunWithStdin : Execute Commands
func RunWithStdin(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RunWithStdin: %q\n", out.String())
	return out.String()
}

// Run : Execute Commands
func Run(name string, arg ...string) string {
	var out bytes.Buffer

	cmd := exec.Command(name, arg...)
	//cmd.Stdin = strings.NewReader("some input")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Run: %q\n", out.String())
	return out.String()
}

func useStartProcess(message []byte) {

	// // 处理两个PIPE
	// outr, outw, err1 := os.Pipe()
	// if err1 != nil {
	// 	internalError(ws, "stdout:", err1)
	// 	return
	// }
	// defer outr.Close()
	// defer outw.Close()

	var err2 error
	cmdToRun, err2 := exec.LookPath(string(message))
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(cmdToRun)

	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	if process, err3 := os.StartProcess(cmdToRun, []string{""}, procAttr); err3 != nil {
		fmt.Printf("ERROR Unable to run %s: %s \n", cmdToRun, err3.Error())
	} else {
		fmt.Printf("%s running as pid %d \n", cmdToRun, process.Pid)
	}
}

// func internalError(ws *websocket.Conn, msg string, err error) {
// 	log.Println(msg, err)
// 	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
// }
