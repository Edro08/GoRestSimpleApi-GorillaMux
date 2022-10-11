package bootstrap

import (
	"GoRestSimpleApiLimpio/internal/task/platform/server/handler"
	"GoRestSimpleApiLimpio/internal/task/platform/storage/mysql"
	"GoRestSimpleApiLimpio/internal/task/read"
	"GoRestSimpleApiLimpio/kit/platform/server"
	"GoRestSimpleApiLimpio/kit/platform/storage"
	"net/http"
)

func Run() error {

	mysqlURI := storage.GetmysqlURI()
	db, err := storage.SetupStorage(mysqlURI)

	if err != nil {
		return err
	}

	taskRepository := mysql.TaskNewRepository(db)
	taskGetIdService := read.NewTaskServiceID(taskRepository)
	taskGetIdHandler := handler.NewTaskHandlerID(taskGetIdService)

	http.HandleFunc("/getIdtask", taskGetIdHandler.TaskReadIDHandler)

	return server.ServerRun()
}
