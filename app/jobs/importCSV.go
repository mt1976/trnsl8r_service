package jobs

import (
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/timing"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

type TextImportModel struct {
	Original string `csv:"original"`
	Message  string `csv:"message"`
}

type importCSVData struct {
}

func (job importCSVData) Run() error {

	j := timing.Start(domains.JOBS.String(), "Import", "Translation CSV")

	err := textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		j.Stop(0)
	}

	j.Stop(1)
	return nil
}

func (job importCSVData) Service() func() {
	return func() {
		job.Run()
	}
}

func (job importCSVData) Schedule() string {
	return "10 7 * * *"
}

func (job importCSVData) Name() string {
	return translation.Get("Import Translation CSV", "")
}
