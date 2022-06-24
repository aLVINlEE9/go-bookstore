package controllers

import {
	"encoding./json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/aLVINlEE9/go-bookstore/pkg/utils"
	"github.com/aLVINlEE9/go-bookstore/pkg/models"
}

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Requset){
	vars := mux.Vars(r)
	bookId := vars[["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
}