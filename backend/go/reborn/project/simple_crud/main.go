package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type hr struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

func main() {

	var db []hr

	//响应/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎使用人力资源管理系统")
	})

	//响应/insert，从传入的参数新增人员信息
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "错误的请求")
			} else {
				name := r.FormValue("name")
				// 将string转为int
				age, _ := strconv.Atoi(r.FormValue("age"))
				gender, _ := strconv.Atoi(r.FormValue("gender"))
				db = append(db, hr{Id: len(db) + 1, Name: name, Age: age, Gender: gender})
				fmt.Fprintln(w, "添加了"+name)
			}
		}
	})

	//响应/delete,从传入的参数删除人员信息
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "错误的请求")
			} else {
				id, _ := strconv.Atoi(r.FormValue("id"))
				var ret int
				for k, v := range db {
					if v.Id == id {
						ret = id
						db = append(db[:k], db[k+1:]...)
					}
				}

				if ret == 0 {
					fmt.Fprintln(w, "404")
				} else {
					fmt.Fprintln(w, "删除了id:"+strconv.Itoa(ret))
				}
			}
		}
	})

	//响应/update,从传入的参数更新人员信息
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "错误的请求")
			} else {
				id, _ := strconv.Atoi(r.FormValue("id"))
				var ret int
				name := r.FormValue("name")
				// 将string转为int
				age, _ := strconv.Atoi(r.FormValue("age"))
				gender, _ := strconv.Atoi(r.FormValue("gender"))
				for i := 0; i < len(db); i++ {
					if db[i].Id == id {
						ret = id
						db[i].Name = name
						db[i].Age = age
						db[i].Gender = gender
					}
				}

				if ret == 0 {
					fmt.Fprintln(w, "404")
				} else {
					fmt.Fprintln(w, "更新了id:"+strconv.Itoa(ret))
				}
			}
		}
	})

	//响应/query，获取所有已存在的人员信息。
	//给定format，可按json格式输出，默认格式为字符串
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "错误的请求")
		} else {
			format := r.FormValue("format")
			if format == "json" {
				data, err := json.Marshal(db)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Fprintln(w, string(data))
				}
			} else {
				for i := 0; i < len(db); i++ {
					fmt.Fprintln(w, db[i].Name, db[i].Age, db[i].Gender)
				}
			}
		}
	})

	//启动本地服务器（localhost）
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("启动服务失败，错误信息：", err)
	}
}
