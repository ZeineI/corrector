package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/ZeineI/corrector/config"
	"go.uber.org/zap"
)

type RequestData struct {
	Texts []string `json:"texts"`
}

type Response [][]struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
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
	// var result []string

	for _, sentence := range text {
		urL, err := url.Parse(cfg.API.Url)
		if err != nil {
			logger.Info("Parse url error")
			return nil, err
		}

		values := urL.Query()
		values.Add("text", sentence)
		values.Add("lang", "ru")
		values.Add("format", "plain")

		urL.RawQuery = values.Encode()

		req, err := http.NewRequest(http.MethodGet, urL.String(), nil)
		if err != nil {
			logger.Info("Cant form request")
			return nil, err
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Info("Send request error")
			return nil, err
		}

		if response.StatusCode != http.StatusOK {
			logger.Info(response.StatusCode)
			return nil, err
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logger.Info("response body io error")
			return nil, err
		}

		defer response.Body.Close()

		var resp Response
		if err := json.Unmarshal(body, &resp); err != nil {
			logger.Info(err)
			return nil, err
		}
		fmt.Println(resp)
	}

	return nil, nil
}
