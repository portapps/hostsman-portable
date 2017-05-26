//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=HostsManPortable.ico
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/op/go-logging"
)

const (
	NAME = "HostsManPortable"
)

var (
	log       = logging.MustGetLogger(NAME)
	logFormat = logging.MustStringFormatter(`%{time:2006-01-02 15:04:05} %{level:.4s} - %{message}`)
)

func main() {
	// Current path
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("Current path:", err)
	}

	// Log file
	logfile, err := os.OpenFile(path.Join(currentPath, NAME+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Error("Log file:", err)
	}

	// Init logger
	logBackendStdout := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), logFormat)
	logBackendFile := logging.NewBackendFormatter(logging.NewLogBackend(logfile, "", 0), logFormat)
	logging.SetBackend(logBackendStdout, logBackendFile)
	log.Info("--------")
	log.Info("Starting " + NAME + "...")
	log.Info("Current path:", currentPath)

	// Init vars
	var exe = path.Join(currentPath, "hm.exe")
	var progdataPath = path.Join(currentPath, "progdata")
	var appdataPath = path.Join(currentPath, "appdata")
	var symlinkBaseProgdataPath = path.Clean(path.Join(os.Getenv("PROGRAMDATA"), "abelhadigital.com"))
	var symlinkProgdataPath = path.Clean(path.Join(symlinkBaseProgdataPath, "HostsMan"))
	var symlinkBaseAppdataPath = path.Clean(path.Join(os.Getenv("APPDATA"), "abelhadigital.com"))
	var symlinkAppdataPath = path.Clean(path.Join(symlinkBaseAppdataPath, "HostsMan"))
	log.Info("Executable:", exe)
	log.Info("Progdata path:", progdataPath)
	log.Info("Appdata path:", appdataPath)
	log.Info("Symlink Progdata path:", symlinkProgdataPath)
	log.Info("Symlink Appdata path:", symlinkAppdataPath)

	// Create progdata folder
	if _, err := os.Stat(progdataPath); os.IsNotExist(err) {
		log.Info("Create progdata folder...", progdataPath)
		err = os.Mkdir(progdataPath, 777)
		if err != nil {
			log.Error("Create progdata folder:", err)
		}
	}

	// Create appdata folder
	if _, err := os.Stat(appdataPath); os.IsNotExist(err) {
		log.Info("Create appdata folder...", appdataPath)
		err = os.Mkdir(appdataPath, 777)
		if err != nil {
			log.Error("Create appdata folder:", err)
		}
	}

	// Check old progdata folder
	if _, err := os.Stat(symlinkBaseProgdataPath); os.IsNotExist(err) {
		log.Info("Create base progdata folder...", symlinkBaseProgdataPath)
		err = os.Mkdir(symlinkBaseProgdataPath, 777)
		if err != nil {
			log.Error("Create base progdata folder:", err)
		}
	}
	if _, err := os.Stat(symlinkProgdataPath); err == nil {
		fi, err := os.Lstat(symlinkProgdataPath)
		if err != nil {
			log.Error("Symlink lstat:", err)
		}
		if fi.Mode()&os.ModeSymlink != os.ModeSymlink {
			// Copy old data folder
			log.Info("Copy old data from", symlinkProgdataPath)
			err = copyDir(symlinkProgdataPath, progdataPath)
			if err != nil {
				log.Error("Copying old data folder:", err)
			}

			// Rename old data folder
			log.Info("Chmod old data folder...")
			err = os.Chmod(symlinkProgdataPath, 0777)
			if err != nil {
				log.Error("Chmod old data folder:", err)
			}

			log.Info("Rename old data folder to", symlinkProgdataPath+"_old")
			err = os.Rename(symlinkProgdataPath, symlinkProgdataPath+"_old")
			if err != nil {
				log.Error("Renaming old data folder:", err)
			}
		}
	}

	// Check old appdata folder
	if _, err := os.Stat(symlinkBaseAppdataPath); os.IsNotExist(err) {
		log.Info("Create base appdata folder...", symlinkBaseAppdataPath)
		err = os.Mkdir(symlinkBaseAppdataPath, 777)
		if err != nil {
			log.Error("Create base appdata folder:", err)
		}
	}
	if _, err := os.Stat(symlinkAppdataPath); err == nil {
		fi, err := os.Lstat(symlinkAppdataPath)
		if err != nil {
			log.Error("Symlink lstat:", err)
		}
		if fi.Mode()&os.ModeSymlink != os.ModeSymlink {
			// Copy old data folder
			log.Info("Copy old data from", symlinkAppdataPath)
			err = copyDir(symlinkAppdataPath, appdataPath)
			if err != nil {
				log.Error("Copying old data folder:", err)
			}

			// Rename old data folder
			log.Info("Chmod old data folder...")
			err = os.Chmod(symlinkAppdataPath, 0777)
			if err != nil {
				log.Error("Chmod old data folder:", err)
			}

			log.Info("Rename old data folder to", symlinkAppdataPath+"_old")
			err = os.Rename(symlinkAppdataPath, symlinkAppdataPath+"_old")
			if err != nil {
				log.Error("Renaming old data folder:", err)
			}
		}
	}

	// Create progdata symlink
	log.Info("Create symlink", symlinkProgdataPath)
	os.Remove(symlinkProgdataPath)
	cmd := exec.Command("cmd", "/c", "mklink", "/J", strings.Replace(symlinkProgdataPath, "/", "\\", -1), strings.Replace(progdataPath, "/", "\\", -1))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Error("Symlink:", err)
	}

	// Create appdata symlink
	log.Info("Create symlink", symlinkAppdataPath)
	os.Remove(symlinkAppdataPath)
	cmd = exec.Command("cmd", "/c", "mklink", "/J", strings.Replace(symlinkAppdataPath, "/", "\\", -1), strings.Replace(appdataPath, "/", "\\", -1))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Error("Symlink:", err)
	}

	// Launch
	log.Info("Launch HostsMan...")
	cmd = exec.Command(exe)
	cmd.Dir = currentPath

	defer logfile.Close()
	cmd.Stdout = logfile
	cmd.Stderr = logfile

	if err := cmd.Start(); err != nil {
		log.Error("Cmd Start:", err)
	}

	cmd.Wait()
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}

	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
