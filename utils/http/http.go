package http

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func EnableRewriteEngine() bool {
	input, err := os.ReadFile("/etc/apache2/apache2.conf")
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`<Directory /var/www/>[\w\s]*?</Directory>`)
	fullTag := r.FindAllString(string(input), -1)

	r = regexp.MustCompile(`AllowOverride\s\w*`)
	allowOverride := r.FindAllString(fullTag[0], -1)

	if allowOverride[0] == "AllowOverride None" {
		newValue := strings.Replace(fullTag[0], "AllowOverride None", "AllowOverride All", 1)
		newValue = strings.Replace(string(input), fullTag[0], newValue, 1)
		os.WriteFile("/etc/apache2/apache2.conf", []byte(newValue), 0644)
		return true
	} else if allowOverride[0] == "AllowOverride All" {
		log.Println("AllowOverride is already set to All")
		return false
	} else {
		log.Println(allowOverride[0], "-> AllowOverride All")
		newValue := strings.Replace(fullTag[0], allowOverride[0], "AllowOverride All", 1)
		newValue = strings.Replace(string(input), fullTag[0], newValue, 1)
		os.WriteFile("/etc/apache2/apache2.conf", []byte(newValue), 0644)
		return true
	}
}

func EnableMod() {
	cmd := exec.Command("a2enmod", "rewrite", "proxy", "proxy_http")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RestartApache() {
	cmd := exec.Command("service", "apache2", "restart")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func AddHtaccess() {
	// check if .htaccess exists
	_, err := os.Stat("/var/www/html/.htaccess")
	if os.IsNotExist(err) {
		// create .htaccess with 644 permissions
		os.WriteFile("/var/www/html/.htaccess", []byte("RewriteEngine On"), 0644)
	}
}
