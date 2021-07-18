package ftp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	goftp "github.com/jlaffaye/ftp"
)

// turn on ftp
// sudo launchctl load -w /System/Library/LaunchDaemons/tftp.plist
// turn off ftp
// sudo launchctl unload -w /System/Library/LaunchDaemons/tftp.plist
// check status
// netstat -atp TCP | grep ftp

//ConnnectFTP is connect to ftp connection
func ConnnectFTP(url, port, user, pwd string) *goftp.ServerConn {
	// fmt.Println("connect start!")

	fmt.Println("start to ftp connection : ", url+":"+port, user)

	c, err := ftp.Dial(url+":"+port, goftp.DialWithTimeout(120*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(user, pwd)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("connect success!")
	return c
	// // Do something with the FTP conn

	// if err := c.Quit(); err != nil {
	// 	log.Fatal(err)
	// }
}

//GetFromFTP is for getting data from ftp remote
func GetFromFTP(c *goftp.ServerConn, remotepath, localpath, filename, newname, params string) []string {
	fmt.Println("start to check ftp file : ", remotepath, localpath, filename, newname)
	filelist := []string{}
	// walker := c.Walk(remotepath)
	// for walker.Next() {
	// 	dir := walker.Path()
	// 	dirlist := strings.Split(dir, "/")
	// 	subname := dirlist[len(dirlist)-2]
	// 	entry := walker.Stat()

	// 	fmt.Println("debug check subname: ", subname, dir, entry.Name)
	// 	if strings.Compare(entry.Name, filename) != 0 {
	// 		continue
	// 	}
	// 	res, err := c.Retr(dir)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	rbuf, err := ioutil.ReadAll(res)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	newfilename := path.Join(localpath, params+"_"+newname+"_"+subname+"_"+filename)
	// 	filelist = append(filelist, newfilename)
	// 	file, err := os.OpenFile(newfilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	wbuf := new(bytes.Buffer)
	// 	wbuf.Write(rbuf)
	// 	wbuf.WriteTo(file)

	// 	defer res.Close()
	// }
	err := walk(c, remotepath, func(entry *ftp.Entry, currentPath string) error {
		dir := path.Join(currentPath, entry.Name)
		dirlist := strings.Split(dir, "/")
		subname := dirlist[len(dirlist)-2]

		//fmt.Println("debug check subname: ", subname, dir, entry.Name)

		if strings.Compare(entry.Name, filename) == 0 {

			res, err := c.Retr(dir)
			if err != nil {
				log.Println(err)
			}
			rbuf, err := ioutil.ReadAll(res)
			if err != nil {
				log.Println(err)
			}
			newfilename := path.Join(localpath, params+"_"+newname+"_"+subname+"_"+filename)
			filelist = append(filelist, newfilename)
			file, err := os.OpenFile(newfilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Println(err)
			}
			wbuf := new(bytes.Buffer)
			wbuf.Write(rbuf)
			wbuf.WriteTo(file)

			defer res.Close()
		}
		return nil

	})
	if err != nil {
		log.Println(err)
	}
	return filelist
}

// EntryHandler is function for walk
type EntryHandler func(e *ftp.Entry, currentPath string) error

// 遍歷ftp目錄，獲取文件
func walk(c *goftp.ServerConn, rootDir string, handler EntryHandler) error {
	entries, err := c.List(rootDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {

		switch entry.Type {
		case ftp.EntryTypeFile:
			handler(entry, rootDir)
		case ftp.EntryTypeFolder:
			//	fmt.Println(fmt.Sprintf("%s/%s", rootDir, entry.Name))
			walk(c, fmt.Sprintf("%s/%s", rootDir, entry.Name), handler)
		default:
		}
	}
	return nil
}
