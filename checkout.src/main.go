package main

import (
	"os"
	"flag"
	"fmt"
//	"encoding/json"
	"encoding/base64"
//	"regexp"
	"time"
	"os/exec"
	"strings"
	"syscall"
	"path/filepath"
	"io/ioutil"
	"net/url"
)

var (
    help bool
    interpreter string
    aide string
    output string
    debug bool
)

type FILE struct {
	File string `json:"file"`
	Uid string `json:"uid"`
	Gid string `json:"gid"`
	Modtime string `json:"modtime"`
	State int `json:"state"`
}

func init() {
    flag.BoolVar(&help, "h", false, "this help")
    flag.StringVar(&interpreter, "i", "/bin/bash", "use interpreter")
    flag.StringVar(&aide, "a", "aide", "use aide dir")
    flag.StringVar(&output, "o", "", "output file")
    flag.BoolVar(&debug, "d", false, "use debug")
    flag.Usage = usage
}

func usage() {
   fmt.Println("Options:")
   flag.PrintDefaults()
}

func Filter(arr []string, cond func(string) bool) []string {
   result := []string{}
   for i := range arr {
     if cond(arr[i]) {
       result = append(result, arr[i])
     }
   }
   return result
}

func getOutput(c string,w string) []string {
	if debug {
		fmt.Println(fmt.Sprintf("command: %s",c))
	} 
	cmd := exec.Command(interpreter,"-c",c)
	//args := strings.Split(c, " ")
	//cmd := exec.Command(args[0],args[1:]...)

	if w != "" {
		cmd.Dir = w
	}
	out, err := cmd.CombinedOutput()
	if debug {
		fmt.Println(fmt.Sprintf("stdout: %s",strings.TrimSpace(string(out))))
	}	
	if err != nil {
		if debug {
		 	fmt.Println(fmt.Sprintf("errout: %s",err))
		}	
		return nil
	}
	return  Filter(strings.Split(string(out), "\n"), func(s string) bool {
				return s != ""
			})
}

func fileExists(f string) bool {
    info, err := os.Stat(f)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func dirExists(f string) bool {
    info, err := os.Stat(f)
    if os.IsNotExist(err) {
        return false
    }
    return info.IsDir()
}

func getExecpath() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//path = "/root/monitor"
	return path
}

func grepString(s string) string {
	s = strings.ReplaceAll(url.QueryEscape(s),"%2F","/")
	if debug {
		fmt.Println(fmt.Sprintf("urlencode string: %s",s))
	}

	if strings.Contains(s, "+"){
		s = strings.ReplaceAll(s,"+","%20")
	}
	if strings.Contains(s, "%3F"){
		s = strings.ReplaceAll(s,"%3F","\\?")
	}
	if strings.Contains(s, "%21"){
		s = strings.ReplaceAll(s,"%21","!")
	}
	if strings.Contains(s, "%24"){
		s = strings.ReplaceAll(s,"%24","\\$")
	}
	if strings.Contains(s, "%2B"){
		s = strings.ReplaceAll(s,"%2B","\\+")
	}
	if strings.Contains(s, "%2A"){
		s = strings.ReplaceAll(s,"%2A","\\*")
	}
	if strings.Contains(s, "%3D"){
		s = strings.ReplaceAll(s,"%3D","=")
	}
	if strings.Contains(s, "."){
		s = strings.ReplaceAll(s,".","\\.")
	}
	if strings.Contains(s, "~"){
		s = strings.ReplaceAll(s,"~","%7E")
	}
	if strings.Contains(s, "%28"){
		s = strings.ReplaceAll(s,"%28","\\(")
	}
	if strings.Contains(s, "%29"){
		s = strings.ReplaceAll(s,"%29","\\)")
	}
    return s
}

func fileInfo(s string) FILE {
	info, _ := os.Stat(s)
	//fmt.Println(info.ModTime().Unix())
	return FILE{ File: s,
		Uid: fmt.Sprint(info.Sys().(*syscall.Stat_t).Uid),
		Gid: fmt.Sprint(info.Sys().(*syscall.Stat_t).Gid),
		Modtime: fmt.Sprint(info.ModTime().Unix()),
	}
}

func baseInfo(s string) FILE {

	f := grepString(s)
	//fmt.Println(f)
	if info := getOutput(fmt.Sprintf("cat ./aide.db | grep -E '^%s\\s'",f),fmt.Sprintf("%s/%s",getExecpath(),aide)); info != nil {
		info = strings.Split(info[0]," ")
		modtime, _ := base64.StdEncoding.DecodeString(info[7])
		return FILE{ File: s,
			Uid: info[5],
			Gid: info[6],
			Modtime: string(modtime),
		}	
	}

	now := time.Now()
	today, _ := time.Parse(time.RFC3339, fmt.Sprintf("%04d-%02d-%02dT00:00:00Z",now.Year(),now.Month(),now.Day()))

	return FILE{ File: s,
		Uid: "0",
		Gid: "0",
		Modtime: fmt.Sprint(today.Unix()),
	}	
}

func added(s string) FILE {
	file := fileInfo(s)
	file.State = 1
	if debug {
		fmt.Printf("odj: %+v\n",file)
	}
	return file
}

func removed(s string) FILE {
	file := baseInfo(s)
	file.State = 0
	if debug {
		fmt.Printf("odj: %+v\n",file)
	}
	return file
}

func changed(s string) FILE {
	file := fileInfo(s)
	file.State = 2
	if debug {
		fmt.Printf("odj: %+v\n",file)
	}
	return file
}

func touchFile(f string) {
	file, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
			
func writeFile(m string, f string) {
	message := []byte(m)
	err := ioutil.WriteFile(f, message, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func filesToJson(files []FILE) string {
	s := "["
	n := len(files)
	for i, file := range files {
		f := file.File
		if strings.Contains(f, "\""){
		f = strings.ReplaceAll(f,"\"","\\\"")
		}
		s = s + fmt.Sprintf("{\"file\":\"%s\",\"uid\":\"%s\",\"gid\":\"%s\",\"modtime\":\"%s\",\"state\":%d}",f,file.Uid,file.Gid,file.Modtime,file.State)
		if i < n - 1 {
			s = s + ","
		}
	}
	return s + "]"
}

func main() {
    flag.Parse()

    if help {
        flag.Usage()
        os.Exit(0)
    }

    files := []FILE{}
    for _, str := range getOutput(fmt.Sprint("./aide -c aide.conf | grep -E '(added:)|(removed:)|(changed:)'"),fmt.Sprintf("%s/%s",getExecpath(),aide)) {
        
        file := strings.SplitN(str," ",2)
        if debug {
			fmt.Println(fmt.Sprintf("file: %s, state: %s",file[1],file[0]))
		}
        switch file[0] {
			case "added:":
				files = append(files, added(file[1]))
			case "removed:":
				files = append(files, removed(file[1]))
			case "changed:":
				if !fileExists(file[1]) {
					continue
				}
				files = append(files, changed(file[1]))
		}
    }
    //fmt.Println(files)
    //jsondata, _ := json.Marshal(files)
    jsondata := filesToJson(files)
    if output == "" {
    	fmt.Println(jsondata)
    } else {
    	fmt.Println(fmt.Sprintf("output json %s",output))
    	writeFile(jsondata,output)
    }

}