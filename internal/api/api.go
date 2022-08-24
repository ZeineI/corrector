package api

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ZeineI/corrector/config"
	"go.uber.org/zap"
)

type RequestData struct {
	Texts []string `json:"texts"`
}

func GetResponse(cfg *config.Config, logger *zap.SugaredLogger) ([]string, error) {
	//get text
	text, err := getText(cfg, logger)
	if err != nil {
		logger.Info("GetText function errors")
		return nil, err
	}

	result, err := doHTTP(cfg, logger, text)
	if err != nil {
		logger.Info("DoHTTP request function errors")
		return nil, err
	}

	return result, nil
}

func getText(cfg *config.Config, logger *zap.SugaredLogger) ([]string, error) {

	jsonFile, err := os.Open(cfg.API.DataPath)
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	logger.Info("Successfully opened file")
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	var reqData RequestData
	if err := json.Unmarshal(data, &reqData); err != nil {
		logger.Info(err)
		return nil, err
	}

	return reqData.Texts, nil
}

func doHTTP(cfg *config.Config, logger *zap.SugaredLogger, text []string) ([]string, error) {
	var result []string
}
