package api

import (
    "encoding/json"
    "fmt"
    "github.com/NyroCodes/NoxFollow/db"
    _ "github.com/NyroCodes/NoxFollow/db"
    "net/http"
)

type Redirect struct {
    Url string
}

type RedirectHandler struct {
    
}

func (h *RedirectHandler) Initialize(mux *http.ServeMux) {
    mux.HandleFunc("GET /api/redirect", h.GetList)
    mux.HandleFunc("POST /api/redirect", h.Post)
    mux.HandleFunc("GET /api/redirect/{id}", h.Get)
    mux.HandleFunc("PUT /api/redirect/{id}", h.Put)
    mux.HandleFunc("DELETE /api/redirect/{id}", h.Delete)
}

func (h *RedirectHandler) GetList(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *RedirectHandler) Post(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *RedirectHandler) Get(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *RedirectHandler) Put(w http.ResponseWriter, r *http.Request)  {
    key := r.PathValue("id")
    
    var redirect Redirect
    
    fmt.Print(r.Body)
    err := json.NewDecoder(r.Body).Decode(&redirect)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Print(redirect)
    _, err = db.WriteUrl(redirect.Url, key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }
    
    w.WriteHeader(http.StatusOK)
}

func (h *RedirectHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}


type FollowHandler struct { }

func (h *FollowHandler) Initialize(mux *http.ServeMux) {
    mux.HandleFunc("GET /follow", h.GetList)
    mux.HandleFunc("POST /follow", h.Post)
    mux.HandleFunc("GET /follow/{id}", h.Get)
    mux.HandleFunc("PUT /follow/{id}", h.Put)
    mux.HandleFunc("DELETE /follow/{id}", h.Delete)
}

func (h *FollowHandler) GetList(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *FollowHandler) Post(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *FollowHandler) Get(w http.ResponseWriter, r *http.Request)  {
    key := r.PathValue("id")
    url, err := db.ReadUrl(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *FollowHandler) Put(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}

func (h *FollowHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("TODO: GetList")
    w.WriteHeader(http.StatusNotImplemented)
}