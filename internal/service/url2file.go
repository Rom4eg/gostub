package service

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *Service) Url2File(urlPath string) (string, error) {
	s.l.Debug("Url2File - Start")
	stat, err := os.Stat(urlPath)

	s.l.Debug(fmt.Sprintf("Path exists %s exists: %t", urlPath, err == nil))
	if err != nil {
		return "", err
	}

	if stat.IsDir() {
		s.l.Debug("Path is directory")
		idx := filepath.Join(urlPath, "index")
		return s.Url2File(idx)
	}

	s.l.Debug(fmt.Sprintf("Exact path - %s", urlPath))
	s.l.Debug("Url2File - End")
	return urlPath, nil
}
