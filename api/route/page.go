package route

import (
	"github.com/eliassama/black-gin/i18n"
	"github.com/eliassama/black-gin/middleware"
	"github.com/eliassama/black-gin/res"
	"github.com/gin-gonic/gin"
)

func welcome(ctx *gin.Context) {
	middleware.SafeFunc(func() {

		htmlContent := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
		  	<link rel="icon" type="image/x-icon" href="/black.gin/page/icon">
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Welcome to Black Gin</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    margin: 0;
                    padding: 0;
                    background-color: #121212;
                    color: #ffffff;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    text-align: center;
                }
                .container {
                    padding: 20px;
                    border: 2px solid #ffffff;
                    border-radius: 10px;
                    background-color: #1e1e1e;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
                }
                h1 {
                    font-size: 3em;
                    margin-bottom: 20px;
                    color: #ff5722;
                }
                p {
                    font-size: 1.2em;
                    margin-bottom: 20px;
                }
                .button {
                    padding: 10px 20px;
                    border: none;
                    border-radius: 5px;
                    background-color: #ff5722;
                    color: #ffffff;
                    font-size: 1em;
                    cursor: pointer;
                    transition: background-color 0.3s;
                }
                .button:hover {
                    background-color: #e64a19;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Welcome to Black Gin</h1>
                <p>If you see this page, the black gin web service is successfully start.</p>
                <p>Thank you for using black gin.</p>
            </div>
        </body>
        </html>`

		res.GetMessage(i18n.SendOKMsg(middleware.GetLanguage(ctx))).SendData(ctx, "text/html; charset=utf-8", []byte(htmlContent))
	})

}

func icon(ctx *gin.Context) {
	middleware.SafeFunc(func() {
		// 写入图片数据到响应体
		res.GetMessage(i18n.SendOKMsg(middleware.GetLanguage(ctx))).SendRedirect(ctx, "https://raw.githubusercontent.com/eliassama/stable-diffusion-webui/master/screenshot.png")
		//res.GetMessage(i18n.SendOKMsg(middleware.GetLanguage(ctx))).SendRedirect(ctx, "https://raw.githubusercontent.com/eliassama/black-gin/main/api/static/image/icon.png")
	})

}
