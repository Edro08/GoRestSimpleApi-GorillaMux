package bootstrap

import (
	"GoRestSimpleApi/internal/task/platform/server/handler"
	"GoRestSimpleApi/internal/task/platform/storage/mysql"
	"GoRestSimpleApi/internal/task/read"
	"GoRestSimpleApi/kit/platform/server"
	"GoRestSimpleApi/kit/platform/storage"
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
