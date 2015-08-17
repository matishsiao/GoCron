package main

import (
	"fmt"	
	"flag"
	"time"	
	"strings"
	"io/ioutil"
	"os/exec"
	"bytes"
	"strconv"
)
const layout = "2006-01-02-15-04-05"

var (
	version string = "0.0.1"
	help bool
	script string
	now []string
	weekday int
)

func main() {	
	flag.StringVar(&script,"load", "cron.cs", "load script to run.")
	flag.BoolVar(&help,"help", false, "Show help information")
	flag.Parse()
	fmt.Printf("GoCron Version:%s\n", version)
	if help {
		fmt.Println("========================")
		fmt.Printf("GoCron Version:%s\n", version)
		fmt.Println("-load\tload cronjob script,cronjob script like crontab script.")
		fmt.Println("========================")
	}
	loadScript(script)
	
	for {
		getTime()
		time.Sleep(250 * time.Millisecond)
	}	
}

func getTime() {
	now = strings.Split(time.Now().Format(layout),"-")
	weekday = int(time.Now().Weekday())
}

func loadScript(FileName string) {
	getTime()
	file, e := ioutil.ReadFile(FileName)
	if e != nil {
		fmt.Printf("[Error]: %v\n", e)
		return
	}
	cmdlines := strings.Split(string(file),"\n")

	for _,v := range cmdlines {
		if v == "" {
			continue
		}
		cron := parserCron(v)
		if cron != nil && len(cron) == 6 {
			
			output := "default.log"
			idx := strings.Index(v,">")
			if idx != -1 {
				output = v[idx+2:]
				v = v[0:idx]
			}
			args := strings.Split(v," ")
			args = args[6:]
			if args[0] == "" {
				fmt.Println("command is nil.")			
				continue
			} else {
				fmt.Println("cron:",cron,"args:",args,"output:",output)
				go addCron(cron,args,output)
			}
		}
	}
}

func addCron(cron []string,cmd []string,output string){
	var running string = ""

	for {
		check := true
		for k,v := range cron {
			switch k {
				case 0://sec
					if check {
						check = runCron(5,v)
					} else {
						break
					}
				case 1://min
					if check {
						check = runCron(4,v)
					} else {
						break
					}
				case 2://hour
					if check {
						check = runCron(3,v)
					} else {
						break
					}
				case 3://day
					if check {
						check = runCron(2,v)
					} else {
						break
					}
				case 4://month
					if check {
						check = runCron(1,v)
					} else {
						break
					}
				case 5://weekday
					if check {
						if v == "*" {
							check = true
						} else if v == strconv.Itoa(weekday) {
							check = true
						}
					} else {
						break
					}	
			}
			
		}
		if check && now[5] != running {
			running = now[5]
			runCmd(cmd,output)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func runCron(key int,v string) bool {
	
	if v == "*" {
		return true
	} else if strings.Index(v,"-") != -1 {
		start,err := strconv.Atoi(string(v[0]))
		if err != nil {
			return false
		}
		end,err := strconv.Atoi(string(v[2]))
		if err != nil {
			return false
		}
		iv,err := strconv.Atoi(now[key])
		if err != nil {
			return false
		}
		for i:= start;i <= end;i++{
			if i == iv {
				return true
			}
		}
	 } else if strings.Index(v,"/") != -1 {
		end,err := strconv.Atoi(string(v[1]))
		if err != nil {
			return false
		}
		iv,err := strconv.Atoi(now[key])
		if err != nil {
			return false
		}
		if iv%end == 0 {
			return true
		}
	} else {
		i,err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		iv,err := strconv.Atoi(now[key])
		if err != nil {
			return false
		}
		if i == iv {
			return true
		}
	}
	return false
}

func runCmd(args []string,output string){
		var cmd *exec.Cmd
		cmdstr := args[0]
		if len(args) > 1 {
			args = args[1:]
			cmd = exec.Command(cmdstr,args...)
		} else {
			cmd = exec.Command(cmdstr)
		}
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		msg := ""	
		if err != nil {
			msg = fmt.Sprintf("error %v",fmt.Sprint(err) + "-" + stderr.String())		
		} else {
			msg = out.String()
		} 
		SaveFile("log/",output,[]byte(msg))
}
func parserCron(cron string) []string {
	cronstr := strings.Split(cron," ")
	if len(cronstr) < 6 {
		fmt.Println("args length wrong.",len(cronstr),cron)
		return nil
	}
	cronInfo := cronstr[0:6]
	return cronInfo
	
}