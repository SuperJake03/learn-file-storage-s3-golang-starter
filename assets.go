package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func getAssetPath(mediaType string) (string, error) {
	base := make([]byte, 32)
	rand.Read(base)
	id := base64.RawURLEncoding.EncodeToString(base)
	ext, err := mediaTypeToExt(mediaType)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", id, ext), nil
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func mediaTypeToExt(mediaType string) (string, error) {
	ext, err := mime.ExtensionsByType(mediaType)
	if err != nil {
		return "", err
	}
	return ext[0], nil
}
