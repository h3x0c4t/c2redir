package tcp

import (
	"c2redir/utils/misc"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func ListTCP() error {
	// iptables -t nat -S PREROUTING

	cmd := exec.Command("iptables", "-t", "nat", "-S", "PREROUTING")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	// use regex to parse output `-A PREROUTING -p tcp -m comment --comment c2redir.*--dport.(\d{1,5}).*--to-destination.(.*):(\d{1,5})`

	re := regexp.MustCompile(`-A PREROUTING -p tcp -m comment --comment c2redir.*--dport.(\d{1,5}).*--to-destination.(.*):(\d{1,5})`)
	matches := re.FindAllStringSubmatch(string(stdout), -1)

	for _, match := range matches {
		fmt.Printf("localhost:%s --> %s:%s\n", match[1], match[2], match[3])
	}

	return nil
}

func DelTCP(ip string, lport string, rport string) error {
	// iptables -t nat -D PREROUTING -p tcp  -m comment --comment c2redir --dport 4444 -j DNAT --to-destination xx.xx.xx.xx:4444

	if misc.CheckRuleExists(fmt.Sprintf("PREROUTING -p tcp -m comment --comment c2redir --dport %s -j DNAT --to-destination %s:%s -t nat", lport, ip, rport)) {
		cmd := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-p", "tcp", "-m", "comment", "--comment", "c2redir", "--dport", lport, "-j", "DNAT", "--to-destination", ip+":"+rport)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func AddTCP(ip string, lport string, rport string) error {
	// iptables -t nat -A PREROUTING -p tcp  -m comment --comment c2redir --dport 4444 -j DNAT --to-destination xx.xx.xx.xx:4444

	cmd := exec.Command("sysctl", "-n", "net.ipv4.ip_forward")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if strings.TrimSpace(string(stdout)) != "1" {
		cmd := exec.Command("sysctl", "-w", "net.ipv4.ip_forward=1")
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	if !misc.CheckRuleExists("POSTROUTING -m comment --comment c2redir -j MASQUERADE -t nat") {
		cmd := exec.Command("iptables", "-t", "nat", "-A", "POSTROUTING", "-m", "comment", "--comment", "c2redir", "-j", "MASQUERADE")
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	if !misc.CheckRuleExists(fmt.Sprintf("PREROUTING -p tcp -m comment --comment c2redir --dport %s -j DNAT --to-destination %s:%s -t nat", lport, ip, rport)) {
		cmd := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-m", "comment", "--comment", "c2redir", "--dport", lport, "-j", "DNAT", "--to-destination", ip+":"+rport)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
