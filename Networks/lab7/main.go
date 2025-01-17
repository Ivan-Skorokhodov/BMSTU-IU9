package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

func uploadFile(conn *ftp.ServerConn, localPath, remotePath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = conn.Stor(remotePath, file)
	if err != nil {
		return err
	}
	log.Println("Файл успешно загружен:", remotePath)
	return nil
}

func downloadFile(conn *ftp.ServerConn, remotePath, localPath string) error {
	resp, err := conn.Retr(remotePath)
	if err != nil {
		return err
	}
	defer resp.Close()

	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp)
	if err != nil {
		return err
	}
	log.Println("Файл успешно скачан:", localPath)
	return nil
}

func createDirectory(conn *ftp.ServerConn, dirName string) error {
	err := conn.MakeDir(dirName)
	if err != nil {
		return err
	}
	log.Println("Директория создана:", dirName)
	return nil
}


func deleteFile(conn *ftp.ServerConn, filePath string) error {
	err := conn.Delete(filePath)
	if err != nil {
		return err
	}
	log.Println("Файл удален:", filePath)
	return nil
}

func listDirectory(conn *ftp.ServerConn, dirPath string) error {
	entries, err := conn.List(dirPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		log.Printf("Имя: %s, Тип: %v\n", entry.Name, entry.Type)
	}
	return nil
}

func changeDirectory(conn *ftp.ServerConn, dirPath string) error {
	err := conn.ChangeDir(dirPath)
	if err != nil {
		return err
	}
	log.Println("Текущая директория изменена на:", dirPath)
	return nil
}

func deleteEmptyDirectory(conn *ftp.ServerConn, dirPath string) error {
	err := conn.RemoveDir(dirPath)
	if err != nil {
		return err
	}
	log.Println("Пустая директория удалена:", dirPath)
	return nil
}

func deleteDirectoryRecursive(conn *ftp.ServerConn, dirPath string) error {
	entries, err := conn.List(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Name == "." || entry.Name == ".." {
			continue
		}
		
		fullPath := dirPath + "/" + entry.Name
		log.Println("Удаление:", entry.Name)
		if entry.Type == ftp.EntryTypeFolder {
			err = deleteDirectoryRecursive(conn, fullPath)
		} else {
			err = conn.Delete(fullPath)
		}
		if err != nil {
			return err
		}
	}

	err = conn.RemoveDir(dirPath)
	if err != nil {
		return err
	}
	log.Println("Директория и её содержимое удалены:", dirPath)
	return nil
}


func main() {
	host := "students.yss.su:21"
	username := "ftpiu8"
	password := "3Ru7yOTA"

	conn, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Ошибка соединения: %v", err)
	}
	defer conn.Quit()

	err = conn.Login(username, password)
	if err != nil {
		log.Fatalf("Ошибка авторизации: %v", err)
	}

	log.Println("Успешное подключение к серверу FTP!")

	reader := bufio.NewReader(os.Stdin)
	for {
        fmt.Print("Skorokhodov's custom FTP client> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }

        if input == "" {
            continue
        }

		parts := strings.Split(input, " ")

		switch parts[0] {
		case "upload":
			err = uploadFile(conn, parts[1], parts[2]) // upload test1.txt ./test1.txt
			if err != nil {
				log.Printf("Ошибка при загрузке файла: %v", err)
			}
		case "download":
			err = downloadFile(conn, parts[1], parts[2]) // download ./test1.txt test2.txt
			if err != nil {
				log.Printf("Ошибка при скачивании файла: %v", err)
			}
		case "mkdir":
			err = createDirectory(conn, parts[1]) // mkdir test_dir
			if err != nil {
				log.Printf("Ошибка при создании директории: %v", err)
			}
		case "delete":
			err = deleteFile(conn, parts[1]) // delete ./test1.txt
			if err != nil {
				log.Printf("Ошибка при удалении файла: %v", err)
			}
		case "ls":
			err = listDirectory(conn, parts[1]) // ls ./test_dir
			if err != nil {
				log.Printf("Ошибка при выводе списка файлов: %v", err)
			}
		case "cd":
			err = changeDirectory(conn, parts[1]) // cd ./test_dir
			if err != nil {
				log.Printf("Ошибка при изменении текущей директории: %v", err)
			}
		case "rmdir_empty":
			err = deleteEmptyDirectory(conn, parts[1]) // rmdir_empty ./test_dir
			if err != nil {
				log.Printf("Ошибка при удалении пустой директории: %v", err)
			}
		case "rmdir_req":
			err = deleteDirectoryRecursive(conn, parts[1]) // rmdir_req ./test_dir
			if err != nil {
				log.Printf("Ошибка при удалении директории с содержимым: %v", err)
			}
		default:
			log.Printf("Неизвестная команда: %s", parts[0])
		}
    }
}
