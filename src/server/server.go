package server

import (
    "github.com/gin-gonic/gin"
    "controller"
)

func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    m := r.Group("/memos")
    {
        ctrl := memo.Controller{}
        m.GET("", ctrl.Index)
        m.GET("/:id", ctrl.Show)
        m.POST("", ctrl.Create)
        m.PUT("/:id", ctrl.Update)
        m.DELETE("/:id", ctrl.Delete)
    }

    return r
}