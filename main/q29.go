/*执行程序ps -e -opid,ppid,comm，并格式化输出：
Pid 0 has 2 children: [1 2]
Pid 490 has 2 children: [1199 26524]
Pid 1824 has 1 child: [7293]
*/
package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type Pids struct {
	maps map[int][]int
}

func NewPids() *Pids {
	return &Pids{make(map[int][]int)}
}
func (p *Pids) Append(k int, v int) {
	p.maps[k] = append(p.maps[k], v)
}

func (p *Pids) String() (s string) {
	schild := make([]int, len(p.maps))
	i := 0
	for k, _ := range p.maps {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		ch := "child"
		childNum := len(p.maps[ppid])
		if childNum > 1 {
			ch = "children"
		}
		s += fmt.Sprintf("Pid %d has %d %s:%v\n", ppid, childNum, ch, p.maps[ppid])
	}
	return
}

func f(cmdName string, args ...string) {
	pids := NewPids()
	cmd := exec.Command(cmdName, args...)
	buf, _ := cmd.Output()
	outputList := strings.Split(string(buf), "\n")
	for _, line := range outputList[1:] {
		if len(line) == 0 {
			continue
		}
		params := strings.Fields(line)
		ppid, _ := strconv.Atoi(params[1])
		pid, _ := strconv.Atoi(params[0])
		pids.Append(ppid, pid)
	}

	fmt.Println(pids)

}

func main() {
	f("ps", "-e", "-opid,ppid,comm")
}
