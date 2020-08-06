package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// RemoteStor a ftp conn or sftp conn
type RemoteStor struct {
	ConnURI  string
	User     string
	Password string
	WorkDir  string
	Timeout  time.Duration
	ConnType string // ftp or sftp
	ftpConn  *ftp.ServerConn
	sshConn  *ssh.Client
}

// Init  a sftp conn or ssh conn
func (r *RemoteStor) Init() (err error) {
	if r.ConnType == "sftp" {
		sshConfig := &ssh.ClientConfig{
			User: r.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(r.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		sshConfig.SetDefaults()
		sshConn, err := ssh.Dial("tcp", r.ConnURI, sshConfig)
		if err != nil {
			return err
		}
		r.sshConn = sshConn
		return nil
	}
	if r.ConnType == "ftp" {
		c, err := ftp.Dial(r.ConnURI, ftp.DialWithTimeout(r.Timeout))
		if err != nil {
			return err
		}
		if err = c.Login(r.User, r.Password); err != nil {
			return err
		}
		if err = c.ChangeDir(r.WorkDir); err != nil {
			return err
		}
		r.ftpConn = c
		return nil
	}
	return fmt.Errorf("not a ftp or sftp conn")
}

// Close  ftp conn or ssh conn
func (r *RemoteStor) Close() (err error) {
	if r.ConnType == "ftp" {
		if err = r.ftpConn.Quit(); err != nil {
			return err
		}
		return
	}
	if r.ConnType == "sftp" {
		if err = r.sshConn.Close(); err != nil {
			return err
		}
		return
	}
	return fmt.Errorf("not a ftp or sftp conn")
}

//Stor a file to remote ftp server
func (r *RemoteStor) Stor(remoteFileName, localFilePath string) (err error) {
	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if r.ConnType == "ftp" {
		if err = r.ftpConn.Stor(remoteFileName, f); err != nil {
			return err
		}
		return nil
	}
	if r.ConnType == "sftp" {
		sftpClient, err := sftp.NewClient(r.sshConn)
		if err != nil {
			return err
		}
		remoteFile, err := sftpClient.Create(path.Join(r.WorkDir, remoteFileName))
		if err != nil {
			return err
		}
		defer remoteFile.Close()

		_, err = io.Copy(remoteFile, f)
		if err != nil {
			return err
		}

	}
	return

}
