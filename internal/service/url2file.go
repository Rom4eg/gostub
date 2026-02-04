package service

import (
	"os"
	"path/filepath"
)

func (s *Service) Url2File(urlPath string) (string, error) {
	stat, err := os.Stat(urlPath)
	if err != nil {
		return "", err
	}

	if stat.IsDir() {
		idx := filepath.Join(urlPath, "index")
		return s.Url2File(idx)
	}

	return urlPath, nil
}
