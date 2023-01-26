package main

import (
	"context"
	dpfm_api_input_reader "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Processing_Formatter"
	"convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/config"
	"fmt"
	"time"

	convert_complementer "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/convert_complementer"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}
	defer db.Close()

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), c.RMQ.SessionControlQueue(), c.RMQ.QueueTo(), 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		start := time.Now()
		sdc, err := callProcess(ctx, db, msg, c)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		l.JsonParseOut(sdc)
		msg.Success()
		l.Info("process time %v\n", time.Since(start).Milliseconds())
	}
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(ctx context.Context, db *database.Mysql, msg rabbitmq.RabbitmqMessage, c *config.Conf) (dpfm_api_output_formatter.SDC, error) {
	var err error
	l := logger.NewLogger()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("error occurred: %w", e)
			return
		}
	}()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})

	complementer := convert_complementer.NewConvertComplementer(ctx, db, l)

	sdc := dpfm_api_input_reader.ConvertToSDC(msg.Raw())
	psdc := dpfm_api_processing_formatter.ConvertToSDC()
	osdc := dpfm_api_output_formatter.ConvertToSDC(msg.Raw())

	err = complementer.CreateSdc(&sdc, &psdc, &osdc)
	if err != nil {
		return osdc, err
	}

	return osdc, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getAccepter(input *dpfm_api_input_reader.SDC) []string {
	accepter := input.Accepter
	if len(input.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}
	return accepter
}
