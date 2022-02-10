package main

// E:\4_project\5_Go\bin\fresh
// gox -os "windows linux" -arch amd64
import (
	"database/sql"
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// //go:embed templates\index.html
// var index []byte

//go:embed templates
var html embed.FS

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// box := packr.NewBox("./templates")
	// r.StaticFS("/", box)
	t, _ := template.ParseFS(html, "templates/*.html")
	r.SetHTMLTemplate(t)

	r.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index", gin.H{
		// 	"title": "Gin",
		// })

		// if _, err := c.Writer.Write(index); err != nil {
		// 	fmt.Println(err)
		// }

		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/api/get_homework", func(c *gin.Context) {
		//输出json结果给调用方
		var student_num int
		// var student_name string
		var homework_num int
		var homework_done_num int
		var homework_last_done_time int
		rows, err := db.Query("SELECT * FROM homework_210702_22A WHERE student_num = 2021070247")
		checkErr(err)
		defer rows.Close()
		rows.Next()
		err = rows.Scan(&student_num, &homework_num, &homework_done_num, &homework_last_done_time)
		checkErr(err)
		err = rows.Err()
		checkErr(err)

		var seat_num = strconv.Itoa(student_num)[8:10]
		var class_id = strconv.Itoa(student_num)[2:8]
		c.JSON(200, gin.H{
			"student_num":             student_num,
			"class_id":                class_id,
			"seat_num":                seat_num,
			"homework_num":            homework_num,
			"homework_done_num":       homework_done_num,
			"homework_rema_num":       homework_num - homework_done_num,
			"homework_last_done_time": homework_last_done_time,
		})
	})

	r.Run(":5505")
}
