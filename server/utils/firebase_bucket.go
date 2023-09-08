package util

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadFile(srcPath string, object int64) {
	logger := GetLogger()
	ctx := context.Background()
	opt := option.WithCredentialsFile(".config/serviceAccountCredentials.json")

	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to create a firebase storage client.")
		return
	}

	config, err := LoadConfig(".")
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot load config")
		return
	}

	bkt := client.Bucket(config.BucketName)
	var fileIdx int8 = 1

	err = filepath.WalkDir(srcPath, func(path string, d fs.DirEntry, err error) error {

		file, err := os.Open(path)
		if err != nil {
			logger.Fatal().Err(err).Msg("Unable to open file for reading.")
			return err
		}
		defer file.Close()

		if d.IsDir() {
			return nil
		}

		currObj := bkt.Object(fmt.Sprintf("%d_%d", object, fileIdx))
		objWriter := currObj.NewWriter(ctx)

		if _, err := io.Copy(objWriter, file); err != nil {
			logger.Fatal().Err(err).Msg("Unable to write file contents to bucket object.")
			return err
		}
		if err := objWriter.Close(); err != nil {
			logger.Fatal().Err(err).Msg("Error closing object writer.")
			return err
		}

		fileIdx++
		return nil
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("Cannot walk over src directory.")
	}

	logger.Info().Msg("Files uploaded successfully.")

}
