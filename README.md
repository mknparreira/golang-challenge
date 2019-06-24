# Project Title
This is the challenge code developing with Go Lang. 

The challenge consists of:

Implement a simple API in GO that upon receiving a single request, returns
the following information:

The default parameter-less behavior is to print the numbers from 1 to 100. For multiples
of 3, print the word 'Pé' instead of the number. For multiples of 5, print
the word 'Do' instead of the number. For multiples of both 3 and 5, print 'PéDo'.

## Author
* **Maikon Parreira** - email: mknparreira@gmail.com tel: 913 628 671

### Installing
You can generate a build of the project running the following code inside the challenger/main folder: 

```
go build main.go

```
or just running: 

```
go run main.go
```

After that, the next step is open a navigator on 

The default behavior is to print the numbers from 1 to 100 http://localhost:8080/challenger/api/v1/

or passing params, for instance: http://localhost:8080/challenger/api/v1/from/1/to/100

Replace the number 1 for any integer number. 
Replace the number 100 for any integer number.
Notice that negative numbers aren´t allowed. 

## Note
* I tried extract the documentation of project using godoc. (https://godoc.org/golang.org/x/tools/cmd/godoc). Unfortunanely, I didn´t get the expect. Besides that, I wrote the documentation inside of the code. If you wish, you can check this documentation, you can running the following command:

```
godoc -http=:6060
```
after that, go to the naviagator and open http://localhost:6060/pkg/challenger  

* I used only one third party: (https://github.com/gorilla/mux) because I as my approach, code and implementation will be reviewed, therefore, I believe that it´s not goood to me, user multiple third party to solve the challenge.

* Whatever doubts about the code, approachs or some implementations that I used, you can send email to me (mknparreira@gmail.com) or call me (913 628 671).
* In spite of this position, I would like to say thanks for the opportunity. Opened a new world to me. Go lang is a wonderful language. I going to learn and improve much more. 