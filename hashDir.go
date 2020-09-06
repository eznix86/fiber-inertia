package inertia

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

func hashByte(contentPtr *[]byte)string{

	contents := *contentPtr
	hasher := md5.New()
	hasher.Write(contents)
	return hex.EncodeToString(hasher.Sum(nil))

}

func hashDir(dir string) string {
	var err error
	var finHash = dir
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		sbyte := []byte(finHash)
		concatBytes := hashByte(&sbyte)
		nameByte := []byte(path)
		nameHash := hashByte(&nameByte)
		fileBytes,_ := ioutil.ReadFile(path)
		fileHash := hashByte(&fileBytes)
		finHash  = concatBytes+fileHash+nameHash
		return nil
	})
	if err!=nil{
		os.Exit(1)
	}
	c := []byte(finHash)
	m := hashByte(&c)
	return m
}