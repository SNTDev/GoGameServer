package main

import (
    "bufio"
    "net"
)

type Position struct(
    X int
    Y int
)

type User struct(
    pos Position
    reader *bufio.Reader
    writer *bufio.Writer
)

var users []*User

func makeServer() {
    users = make([]*User, 0)
}

func acceptUsers() {
    in, _ := net.Listen("tcp", ":8081")

    for {
        conn, err := in.Accept()

        if err == nil {

            user := User{
                reader : bufio.NewReader(conn),
                writer : bufio.NewWriter(conn),
            }

            users = append(users, &user)

            go userTreatMent(&user)
        }
    }
}

func userTreatMent(user *User) {
   
}

func main() {
    makeServer()

    acceptUsers()
}
