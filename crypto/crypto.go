package crypto

import (
  "crypto/sha256"
  "crypto/md5"
  "encoding/hex"
  "io"
  "os"
)

func sha256sum(filePath string) (result string, err error) {
  f, err := os.Open(filePath)
  if err != nil {
    return
  }
  defer f.Close()

  h := sha256.New()
  if _, err = io.Copy(h, f); err != nil {
    return
  }

  result = hex.EncodeToString(h.Sum(nil))
  return 
}

func md5sum(filePath string) (result string, err error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 
    }
    defer file.Close()

    hash := md5.New()
    _, err = io.Copy(hash, file)
    if err != nil {
        return 
    }

    result = hex.EncodeToString(hash.Sum(nil))
    return 
}
