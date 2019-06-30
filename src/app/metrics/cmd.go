package metrics

import (
	"app/dindin"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)


func unlabel(name string) {
	defer panics()
	_ = shell("kubectl label node " + name + "   test- --kubeconfig=/etc/config")
}

func label(name string, v int) {
	defer panics()
	s := "  kubectl label node " + name + "  --overwrite test=" + strconv.Itoa(v) + " --kubeconfig=/etc/config"
	export := shell(s)
	fmt.Println(export)
	fmt.Println(name, " label ： test=", v)
}

func cmd(name string) {
	defer panics()
	_ = shell("kubectl taint node " + name + "  status=down:NoSchedule  --kubeconfig=/etc/config")
	s := "echo kubectl drain  " + name + "  --ignore-daemonsets --delete-local-data --grace-period=0 --force --kubeconfig=/etc/config"
	export := shell(s)
	fmt.Println(export)
	dindin.Cordon(name, export)
	fmt.Println(name, "： 紧急移除")

}


func panics() {
	if x := recover(); x != nil {
		log.Printf("caught panic: %v", x)
		dindin.Worring(x)
	}
}
func shell(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s)
	output, _ := cmd.Output()
	return string(output)
}
