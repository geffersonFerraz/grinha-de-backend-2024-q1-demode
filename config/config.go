package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_MAX_POOL_SIZE       uint64
	DB_MIN_POOL_SIZE       uint64
	DB_MAX_CONNECTIONS     uint64
	DB_COMPRESSORS         string
	DB_ZLIB_COMPRESS_LEVEL int
	DB_ZSTD_COMPRESS_LEVEL int
	DB_PORT                string
	DB_HOST                string
	DB_NAME                string
	DB_USER                string
	DB_PASSWORD            string
	SERVER_HOST            string
	SERVER_PORT            string
}

var CFG *Config

func setFieldValue(obj interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByName(fieldName)

	if f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(value))
	}
}

func getEnvOrThrow(config *Config) (*Config, bool) {
	v := reflect.ValueOf(*config)
	typeOfs := v.Type()
	hasError := false
	for i := 0; i < v.NumField(); i++ {
		fieldType := typeOfs.Field(i).Type
		fieldName := typeOfs.Field(i).Name
		if fieldType.Kind() == reflect.String {
			setFieldValue(config, fieldName, os.Getenv(fieldName))
		} else if fieldType.Kind() == reflect.Int {
			valueInt, _ := strconv.Atoi(os.Getenv(fieldName))
			setFieldValue(config, fieldName, valueInt)
		} else if fieldType.Kind() == reflect.Int64 {
			valueInt, _ := strconv.Atoi(os.Getenv(fieldName))
			setFieldValue(config, fieldName, uint64(valueInt))
		}
	}

	return config, hasError
}

func InitConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file - use only env from os")
	}

	config := &Config{}

	config, hasErr := getEnvOrThrow(config)

	if hasErr {
		log.Println("fatal: fail to get all envs")
		os.Exit(1)
	}

	CFG = config
}
