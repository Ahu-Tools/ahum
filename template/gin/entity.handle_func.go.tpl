func (h Handler) {{.HandlerName}}(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello World!",
    })
}