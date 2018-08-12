package files

import (
	"log"
	"os"
	"testing"
)

func TestIsARepoAlready(t *testing.T) {
	t.Run("having a groot folder", func(t *testing.T) {
		grootFolderPath := "a_valid_groot_repo/.groot"
		err := os.MkdirAll(grootFolderPath, 0777)
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
		if !IsARepoAlready(currentDir + "/a_valid_groot_repo") {
			t.Fail()
		}

		err = os.RemoveAll(currentDir + "/a_valid_groot_repo")
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
	})
	t.Run("having a groot file", func(t *testing.T) {
		grootFolderPath := "folder_with_groot_file"
		err := os.Mkdir(grootFolderPath, 0777)
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}

		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}

		f, err := os.OpenFile(currentDir+"/"+grootFolderPath+"/.groot", os.O_RDONLY|os.O_CREATE, 0666)
		defer f.Close()
		if IsARepoAlready(currentDir + "/folder_with_groot_file") {
			t.Fail()
		}

		err = os.RemoveAll(currentDir + "/" + grootFolderPath)
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
	})
	t.Run("not having a groot folder", func(t *testing.T) {
		folderPath := "not_a_groot_repo"
		err := os.MkdirAll(folderPath, 0777)
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}

		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}

		if IsARepoAlready(currentDir + "/" + folderPath) {
			t.Fail()
		}

		err = os.RemoveAll(currentDir + "/" + folderPath)
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
	})
}
