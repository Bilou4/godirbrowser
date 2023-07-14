package main

import (
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func handlerSelection(next, assetsHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		if r.URL.Query().Has("8bf729ffe074caee622c02928173467e658e19e28233cff8a445819e3cae4d50") {
			assetsHandler.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
		log.Printf("%s - %s [%s] %s %s %s total time=%s\n", r.RemoteAddr, start.Format("2006-01-02 15:04:05.000"), r.Method, r.URL.Path, r.URL.Query(), r.UserAgent(), time.Since(start))
	})
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var osFullPath string = rootDir
	var err error
	if path != "/" {
		osFullPath, err = checkPath(filepath.Join(rootDir, path))
		if err != nil {
			log.Println(err)
			return
		}

		if !strings.HasSuffix(path, "/") {
			fileInfo, err := os.Stat(osFullPath)
			if err != nil {
				log.Println(err)
				return
			}

			if fileInfo.IsDir() {
				http.Redirect(w, r, scheme+net.JoinHostPort(host, port)+path+"/", http.StatusMovedPermanently)
				return
			}

			http.ServeFile(w, r, osFullPath)
			return
		}
	}
	dirs, err := os.ReadDir(osFullPath)
	if err != nil {
		log.Println(err)
		return
	}

	data := struct {
		Dirs        []fs.DirEntry
		RootDir     string
		CurrentPath string
	}{
		Dirs:        dirs,
		RootDir:     rootDir,
		CurrentPath: path,
	}
	err = mainTemplate.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}
