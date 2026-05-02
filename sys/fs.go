package sys

import (
  "log"
  "os"
  "path/filepath"
)

func Rootify(path string) string {
  root, err := GetProjectRoot()
  if err != nil {
    log.Fatal(err)
  }

  return filepath.Join(root, path)
}

func GetProjectRoot(dir... bool) (string, error) {
  exeDir, err := os.Executable()
  if err != nil {
    return "", err
  }

  exeDir = filepath.Dir(exeDir)

  for i, dir := range dir {
    if dir {
      exeDir = filepath.Dir(exeDir)
    }

    // up to 3 
    if i > 3 {
      break
    }
  }

  return exeDir, nil
}


