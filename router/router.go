package router

import (
    "github.com/viitorags/encurtadorUrl/config"
    "github.com/viitorags/encurtadorUrl/handler"
    "html/template"
    "net/http"
    "os"
)

var (
    logger *config.Logger
)

func InitRoutes() {
    logger = config.GetLogger("router")
    handler.InitializeHandler()
    basePath := "/api/v1/urls"
    router := http.NewServeMux()

    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        logger.Error("erro ao carregar template: ", err)
        os.Exit(1)
    }

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
            return
        }

        if err := tmpl.Execute(w, nil); err != nil {
            http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        }
    })

    router.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            handler.CreateLink(w, r)
        default:
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    })

    if err := http.ListenAndServe(":8080", router); err != nil {
        logger.Error("config initialization error", err)
        os.Exit(1)
    }
}
