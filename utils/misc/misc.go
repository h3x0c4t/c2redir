package misc

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CheckOS() bool {
	return runtime.GOOS == "linux"
}

func CheckIPT() bool {
	cmd := exec.Command("iptables", "-h")
	err := cmd.Run()
	return err == nil
}

func CheckApache() bool {
	// check if file /etc/apache2/apache2.conf exists
	_, err := os.Stat("/etc/apache2/apache2.conf")
	return err == nil
}

func CheckRoot() bool {
	return os.Geteuid() == 0
}

func CheckSupport() {
	if !CheckOS() {
		log.Fatal("This program only runs on Linux")
	}
	if !CheckRoot() {
		log.Fatal("This program requires root privileges")
	}
	if !CheckIPT() {
		log.Fatal("This program requires iptables")
	}
	if !CheckApache() {
		log.Fatal("This program requires apache2")
	}
}

func CheckRuleExists(rule string) bool {
	args := strings.Split(rule, " ")
	args = append([]string{"-C"}, args...)
	cmd := exec.Command("iptables", args...)
	err := cmd.Run()
	return err == nil
}
