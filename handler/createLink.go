package handler

import (
    "encoding/json"
    "net/http"
)

func CreateLink(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        sendError(w, http.StatusMethodNotAllowed, "Método não permitido")
        return
    }

    body := LinkRequest{}

    if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
        logger.Error("Erro ao decodificar JSON: ", err)
        sendError(w, http.StatusBadRequest, "JSON inválido")
        return
    }

    stmt, err := db.Prepare("INSERT INTO urls (original_url, short_url) VALUES( ?, ? )")
    if err != nil {
        logger.Error("Error creating link", err.Error())
        sendError(w, http.StatusInternalServerError, "Erro ao preparar query")
    }
    defer stmt.Close()

    _, err = stmt.Exec(body.OriginalURL, body.ShortUrl)
    if err != nil {
        logger.Error("Error ao executar query: ", err.Error())
        sendError(w, http.StatusInternalServerError, "Erro ao inserir no banco")
        return
    }

    sendSuccess(w, "create_link", map[string]string{
        "original_url": body.OriginalURL,
        "short_url":    body.ShortUrl,
    })
}
