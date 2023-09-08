package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(srcDir string, targetDir string){
    logger := GetLogger()
    zippedFile, err := zip.OpenReader(srcDir)
    if err != nil {
        logger.Error().Stack().Msg("Cannot open zip file.")
        return
    }
    defer os.Remove(srcDir)

    defer zippedFile.Close()
    for ind, f := range zippedFile.File {
        logger.Info().Msg("Extracting files...")

        currFile, err := f.Open()
if err != nil {
            logger.Error().Stack().Msg("Cannot open zip file contents.")
            return
        }
        defer currFile.Close()

        unzippedPath := filepath.Join(targetDir, f.Name)

        if f.FileInfo().IsDir() {
            logger.Warn().Msg("Zip cannot contain sub directories")
            return
        }

        uncompressedFile, err := os.Create(unzippedPath)
        if err != nil {
            logger.Fatal().Msg(fmt.Sprintf("impossible to create uncompressed: %s", err))
        }
        _, err = io.Copy(uncompressedFile, currFile)
        if err != nil {
            logger.Fatal().Msg(fmt.Sprintf("impossible to copy file n°%d: %s", ind, err))
        }
    }


    logger.Info().Msg("Succesfully extracted zipped file.")
}
