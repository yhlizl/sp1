package ftp

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dutchcoders/goftp"
)

// turn on ftp
// sudo launchctl load -w /System/Library/LaunchDaemons/tftp.plist
// turn off ftp
// sudo launchctl unload -w /System/Library/LaunchDaemons/tftp.plist
// check status
// netstat -atp TCP | grep ftp

//ConnnectFTP is connect to ftp connection
func ConnnectFTP(url, port, user, pwd string) *goftp.FTP {
	// fmt.Println("connect start!")

	c, err := goftp.Connect(url + ":" + port)
	// connConfig is the address configuration, a string of ip: port, such as: localhost:2121

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
func GetFromFTP(c *goftp.FTP, remotepath, localpath string) error {
	file, err := os.OpenFile(localpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Fatal(err)
		return err

	}
	path, err := c.Pwd()
	if err != nil {
		log.Fatal(err)
		file.Close()
		return err

	}
	log.Println("now location:", path)
	c.Retr(remotepath, func(r io.Reader) error {

		wr, err := io.Copy(file, r)
		if err != nil {
			log.Fatal(err)
			file.Close()

			return err
		}
		fmt.Println("check copy ok : ", wr)

		return nil
	})

	return file.Close()

}
