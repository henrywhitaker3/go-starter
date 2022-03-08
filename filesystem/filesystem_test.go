package filesystem_test

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/henrywhitaker3/go-starter/filesystem"
)

type directory struct {
	path        string
	create      bool
	shouldExist bool
}

var directories []directory

func init() {

}

func TestDetrminingIfADirectoryExists(t *testing.T) {
	setupSampleDirs()

	for _, dir := range directories {
		directory := filesystem.Directory{Path: dir.path}
		exists := directory.Exists()

		if dir.shouldExist {
			if !exists {
				t.Error("direcotry does not exist when it is expected to")
			}
		} else {
			if exists {
				t.Error("directory exists when it shouldn't")
			}
		}
	}

	teardownSampleDirs()
}

func TestItCanCreateDirectories(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	dir := filesystem.Directory{Path: "/tmp/" + strconv.Itoa(random)}

	if _, err := os.Stat(dir.Path); err == nil {
		t.Error("directory already exists")
	}

	if err := dir.Create(); err != nil {
		t.Error(err)
	}
}

func setupSampleDirs() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	directories = []directory{
		{path: "/tmp/" + strconv.Itoa(random.Int()), create: true, shouldExist: true},
		{path: "/tmp/" + strconv.Itoa(random.Int()), create: false, shouldExist: false},
	}

	for _, dir := range directories {
		if dir.create {
			err := os.Mkdir(dir.path, 07555)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func teardownSampleDirs() {
	for _, dir := range directories {
		if dir.create {
			err := os.Remove(dir.path)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
