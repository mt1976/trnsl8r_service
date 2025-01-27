package io

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mt1976/appFrame/fileio"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/id"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
	"github.com/mt1976/trnsl8r_service/app/support/timing"
)

var name = "IO"

func GetDBFileName(name string) string {
	if name == "" {
		panic(fmt.Errorf("db name is required"))
	}

	cfg := config.Get()
	sep := "-"

	name = cfg.ApplicationName() + sep + name

	name = strings.ToLower(name)

	path := "%s" + paths.Seperator() + paths.Database().String() + paths.Seperator() + "%s.db"

	xx := fmt.Sprintf(path, paths.Application().String(), name)
	//logger.InfoLogger.Println("DBN: DB File Name: ", xx)
	return xx
}

func Dump(tableName string, where paths.FileSystemPath, action string, recordID string, yy any) {
	cfg := config.Get()
	sep := "-"

	logger.DatabaseLogger.Printf("[SUPPORT] [%v] Dump to '%v'", strings.ToUpper(tableName), where.String())
	id := id.GetUUID()
	if action != "" {
		id = id + sep + cfg.ApplicationName() + sep + tableName + sep + strings.ToTitle(action) + sep + recordID
	}
	id = id + ".json"

	if where.Is(paths.Backups()) {
		id = tableName + sep + action + ".bk"
	}

	path := where.String()

	b, err := json.Marshal(yy)
	if err != nil {
		logger.WarningLogger.Printf("[SUPPORT] [%v] [Marshalling] Error=[%v]", strings.ToUpper(action), err.Error())
		return
	}
	output := string(b)

	fileio.Write(id, path, output)
}

func Backup(table, location string) {
	//path := BACKUPS.path
	sep := "-"
	cfg := config.Get()
	table = strings.ToLower(cfg.ApplicationName() + sep + table)
	logger.EventLogger.Printf("Backup=[%v] [%v.db] to [%v]", strings.ToLower(table), table, location)

	// sleep for 1 second
	time.Sleep(1 * time.Second)
	timing := timing.Start(table, "Backup", "")
	//dateTime := time.Now().Format("20060102150405")
	toPath := paths.Application().String() + location
	toFile := toPath + paths.Seperator() + table + ".db"

	fromPath := paths.Database().String()
	fromFile := paths.Application().String() + paths.Database().String() + paths.Seperator() + table + ".db"
	logger.EventLogger.Printf("Backup=[%v] Path=[%v]", strings.ToLower(table), paths.Application().String())
	logger.EventLogger.Printf("Backup=[%v] Database=[%v.db]", strings.ToLower(table), table)
	logger.EventLogger.Printf("Backup=[%v] From=[%v]", strings.ToLower(table), fromPath)
	logger.EventLogger.Printf("Backup=[%v] To=[%v]", strings.ToLower(table), toPath)

	// remove last char from path
	toPath = toPath[:len(toPath)-1]
	fromPath = fromPath[:len(fromPath)-1]

	err := CopyFile(fromFile, toFile)
	if err != nil {
		logger.ErrorLogger.Printf("Backup=[%v] [%v] to [%v] Error=[%v]", strings.ToLower(name), fromPath, toPath, err.Error())
		panic(err)
	}
	timing.Stop(1)
	logger.EventLogger.Printf("Backup=[%v] COMPLETE", strings.ToLower(table))
}

// File copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func MkDir(path string) error {
	logger.InfoLogger.Printf("[%v] Creating folder Path=[%v]", strings.ToUpper(name), path)
	return os.MkdirAll(path, os.ModeSticky|os.ModePerm)
}

func Dir(path string) ([]string, error) {
	// Get all folders in the backup directory
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders, nil
}

func DeleteFolder(path string) error {
	// Delete the folder
	logger.InfoLogger.Printf("[DELETE][%v] Deleting folder Path=[%v]", strings.ToUpper(name), path)
	return os.RemoveAll(path)
	//return nil
}
