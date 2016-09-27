package server

import (
	"io"
	"os"
)

// This file is the driver part of the server. It must be implemented by anyone wanting to use the server.

// Adding the ClientContext concept to be able to handle more than just UserInfo
// Implemented by the server
type ClientContext interface {
	// Get userInfo
	UserInfo() map[string]string

	// Get current path
	Path() string

	// Custom value. This avoids having to create a mapping between the client.Id and our own internal system. We can
	// just store the driver's instance in the ClientContext
	MyInstance() interface{}

	// Set the custom value
	SetMyInstance(interface{})
}

// FileContext to use
type FileContext interface {
	io.Writer
	io.Reader
	io.Closer
	io.Seeker
}

// Server driver
// Implemented by the driver
type Driver interface {
	// Load some general settings around the server setup
	GetSettings() *Settings

	// Welcome a user
	WelcomeUser(cc ClientContext) (string, error)

	// Authenticate an user
	// Returns nil if the user could be authenticated
	CheckUser(cc ClientContext, user, pass string) error

	// Request to use a directory
	ChangeDirectory(cc ClientContext, directory string) error

	// Create a directory
	MakeDirectory(cc ClientContext, directory string) error

	// List the files of a given directory
	ListFiles(cc ClientContext) ([]os.FileInfo, error)

	// Called when a user disconnects
	UserLeft(cc ClientContext)

	// Upload a file
	OpenFile(cc ClientContext, path string, flag int) (FileContext, error)

	// Delete a file
	DeleteFile(cc ClientContext, path string) error

	// Get some info about a file
	GetFileInfo(cc ClientContext, path string) (os.FileInfo, error)
}

// Settings are part of the driver
type Settings struct {
	Host           string // Host to receive connections on
	Port           int    // Port to listen on
	MaxConnections int    // Max number of connections to accept
	MaxPassive     int    // Max number of passive connections per control connections to accept
	MonitorOn      bool   // To activate the monitor
	MonitorPort    int    // Port for the monitor to listen on
}
