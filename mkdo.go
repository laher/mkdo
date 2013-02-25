package main
/*
   Copyright 2013 Am Laher

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)


var (
	//use ldflags (or goxc) to populate main.VERSION during build
	VERSION	    = "unknown"
	flagSet     = flag.NewFlagSet("mkdo", flag.ExitOnError)
	verbose     bool
	interactive bool
	isHelp      bool
	isVersion   bool
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdoDirs(args []string) {
	if verbose {
		log.Printf("checking dirs: %s", args)
	}
	for _, arg := range args {
		if slashy(arg) {
			//get last index of slash
			dirpart := dirpart(arg)
			exists, err := exists(dirpart)
			//TODO: Recursively check permissions before attempting to mkdir ..
			if !exists {
				var cont bool
				if interactive {
					fmt.Printf("Make dir '%s'? [Y/n]", dirpart)
					var confirm string
					fmt.Scanf("%s", &confirm)
					confirm = strings.TrimSpace(confirm)
					if len(confirm) == 0 || strings.ToLower(confirm) == "y" {
						cont = true
					} else {
						cont = false
					}
				} else {
					cont = true
				}
				if cont {
					if err != nil {
						log.Printf("Error checking if directory '%s' exists: %s", dirpart, err)
					}
					if verbose {
						log.Printf("Making dir '%s'", dirpart)
					}
					os.MkdirAll(dirpart, 0755)
				} else {
					if verbose {
						log.Printf("Not making dir '%s'", dirpart)
					}
				}
			} else {
				if verbose {
					log.Printf("Dir '%s' exists already", dirpart)
				}
			}
		} else {
			if verbose {
				log.Printf("Skipping non-slashy %s", arg)
			}
		}
	}
}

func dirpart(arg string) string {
	return arg[0:strings.LastIndex(arg, string(os.PathSeparator))]
}

func slashy(arg string) bool {
	//slashes present
	if strings.LastIndex(arg, string(os.PathSeparator)) > 1 {
		return true
	}
	return false
}

func redirectIOStandard(cmd *exec.Cmd) (*os.File, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
	}
	if verbose {
		log.Printf("Redirecting output")
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	//direct. Masked passwords work OK!
	cmd.Stdin = os.Stdin
	return nil, err
}

func redirectIO(cmd *exec.Cmd) (*os.File, error) {
	f, err := redirectIOStandard(cmd)
	return f, err
}
func run(args []string) (int, error) {
	p, err := exec.LookPath(args[0])
	if err != nil {
		log.Printf("Couldn't find exe %s - %s", p, err)
	}
	if true {
		cmd := exec.Command(args[0])
		cmd.Args = args
		if verbose {
			log.Printf("Running cmd: %s", args)
		}

		f, err := redirectIO(cmd)
		if err != nil {
			log.Printf("Error redirecting IO: %s", err)
		}
		if f != nil {
			defer f.Close()
		}

		err = cmd.Start()
		if err != nil {
			log.Printf("Launch error: %s", err)
			return 1, err
		} else {
			if verbose {
				log.Printf("Waiting for command to finish...")
			}
			err = cmd.Wait()
			if err != nil {
				if verbose {
					log.Printf("Command exited with error: %v", err)
				}
			} else {
				if verbose {
					log.Printf("Command completed without error")
				}
			}
		}
		if err != nil {
			if e2, ok := err.(*exec.ExitError); ok { // there is error code
				processState, ok2 := e2.Sys().(syscall.WaitStatus)
				if ok2 {
					errcode := processState.ExitStatus()
					//TODO: Check on windows. Google groups suggests Windows uses processState.ExitCode instead, but it's not in the docs...
					log.Printf("%s returned exit status: %d", cmd.Args[0], errcode)
					return errcode, err
				}
			}
			return 1, err
		}
	}
	return 0, nil
}

func printHelp() {
	fmt.Fprint(os.Stderr, "`mkdo` [options] <cmd> <paths"+string(os.PathSeparator)+"with"+string(os.PathSeparator)+"slashes"+string(os.PathSeparator)+">\n")
	fmt.Fprintf(os.Stderr, " Version %s. Options:\n", VERSION)
	flagSet.PrintDefaults()
	fmt.Fprint(os.Stderr, "\nTip 1: mkdo recognises folders by the last trailing"+string(os.PathSeparator)+"slash"+string(os.PathSeparator)+" in an argument.\n")
}

func printVersion() {
	fmt.Fprintf(os.Stderr, " mkdo version %s\n", VERSION)
}

func Mkdo(call []string) (int, error) {
	if len(call) < 2 { //no options - invalid
		printHelp()
		return 1, nil
	}

	//no options - just pass all args to cli
	if strings.Index(call[1], "-") != 0 {
		mkdoDirs(call[2:])
		return run(call[1:])
	} else {
		//options - go to parse
		e := flagSet.Parse(call[1:])
		if e != nil {
			return 1, e
		}
		remainder := flagSet.Args()
		if isHelp {
			printHelp()
			return 0, nil
		} else if isVersion {
			printVersion()
			return 0, nil
		} else if len(remainder) < 1 {
			printHelp()
			return 1, nil
		} else {
			mkdoDirs(remainder[1:])
			return run(remainder)
		}
	}
	return 1, nil
}

func main() {
	log.SetPrefix("[mkdo] ")
	flagSet.BoolVar(&isHelp, "h", false, "Show this help")
	flagSet.BoolVar(&isVersion, "version", false, "version info")
	flagSet.BoolVar(&verbose, "v", false, "verbose")
	flagSet.BoolVar(&interactive, "i", false, "interactive mode")
	errcode, _ := Mkdo(os.Args)
	if errcode != 0 {
		if verbose {
			log.Printf("Exiting with status: %d.", errcode)
		}
		os.Exit(errcode)
	}
}
