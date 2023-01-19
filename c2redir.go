package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("Please run this on Linux!")
		os.Exit(0)
	}
}

func checkIPT() {
	cmd := exec.Command("iptables", "-h")
	err := cmd.Run()
	if err != nil {
		fmt.Println("iptables not found")
		os.Exit(0)
	}
}

func checkRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("Please run this as root!")
		os.Exit(0)
	}
}

func checkRuleExists(rule string) bool {
	args := strings.Split(rule, " ")
	args = append([]string{"-C"}, args...)
	cmd := exec.Command("iptables", args...)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func fwdTCP(c2ip string, port string) error {
	// iptables -I INPUT -p tcp --dport 4444 -j ACCEPT
	// iptables -t nat -A PREROUTING -p tcp --dport 4444 -j DNAT --to-destination

	if !checkRuleExists(fmt.Sprintf("INPUT -p tcp --dport %s -j ACCEPT", port)) {
		cmd1 := exec.Command("iptables", "-I", "INPUT", "-p", "tcp", "--dport", port, "-j", "ACCEPT")
		err := cmd1.Run()
		if err != nil {
			fmt.Println("Error adding rule!")
			return err
		}
	}

	if !checkRuleExists(fmt.Sprintf("PREROUTING -p tcp --dport %s -j DNAT --to-destination %s:%s -t nat", port, c2ip, port)) {
		cmd2 := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "--dport", port, "-j", "DNAT", "--to-destination", c2ip+":"+port)
		err := cmd2.Run()
		if err != nil {
			fmt.Println("Error adding rule!")
			return err
		}
	}

	return nil
}

func setBaseRules() error {
	// iptables -t nat -A POSTROUTING -j MASQUERADE
	// iptables -I FORWARD -j ACCEPT
	// iptables -P FORWARD ACCEPT

	if !checkRuleExists("POSTROUTING -j MASQUERADE -t nat") {
		cmd := exec.Command("iptables", "-t", "nat", "-A", "POSTROUTING", "-j", "MASQUERADE")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error adding rule!")
			return err
		}
	}

	if !checkRuleExists("FORWARD -j ACCEPT") {
		cmd := exec.Command("iptables", "-I", "FORWARD", "-j", "ACCEPT")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error adding rule!")
			return err
		}
	}

	cmd := exec.Command("iptables", "-S", "FORWARD")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("Error adding rule!")
		return err
	}

	if !strings.Contains(string(stdout), "-P FORWARD ACCEPT") {
		cmd := exec.Command("iptables", "-P", "FORWARD", "ACCEPT")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error adding rule!")
			return err
		}
	}

	return nil
}

func main() {

	checkOS()
	checkRoot()
	checkIPT()

	m := flag.String("m", "", "TCP, UDP, HTTP")
	c2ip := flag.String("c2ip", "", "c2 ip address")
	port := flag.String("port", "", "port")

	flag.Parse()
	switch *m {
	case "TCP":

		// iptables -I INPUT -p tcp --dport 4444 -j ACCEPT
		// iptables -t nat -A PREROUTING -p tcp --dport 4444 -j DNAT --to-destination 10.33.33.30:4444
		// iptables -t nat -A POSTROUTING -j MASQUERADE
		// iptables -I FORWARD -j ACCEPT
		// iptables -P FORWARD ACCEPT

		if *c2ip == "" || *port == "" {
			flag.Usage()
			os.Exit(0)
		}

		err := setBaseRules()
		if err != nil {
			fmt.Println("Error setting base rules!")
			os.Exit(0)
		}

		err = fwdTCP(*c2ip, *port)
		if err != nil {
			fmt.Println("Error forwarding port!")
			os.Exit(0)
		}

	case "UDP":
		log.Println("UDP")
	case "HTTP":
		log.Println("HTTP")
	default:
		flag.Usage()
	}

}
