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
func GetFromFTP(c *goftp.ServerConn, remotepath, localpath, filename, newname string) {
	fmt.Println("start to check ftp file : ", remotepath, localpath, filename, newname)

	walker := c.Walk(remotepath)
	for walker.Next() {
		entry := walker.Stat()
		dir := walker.Path()
		dirlist := strings.Split(dir, "/")
		subname := dirlist[len(dirlist)-2]
		if entry.Name != filename {
			continue
		}
		res, err := c.Retr(dir)
		if err != nil {
			log.Println(err)
		}
		rbuf, err := ioutil.ReadAll(res)
		if err != nil {
			log.Println(err)
		}
		newfilename := path.Join(localpath, newname+"_"+subname+"_"+filename)
		file, err := os.OpenFile(newfilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Println(err)
		}
		wbuf := new(bytes.Buffer)
		wbuf.Write(rbuf)
		wbuf.WriteTo(file)

		defer res.Close()
	}

}
