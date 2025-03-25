package main

import "fmt"

// Define the Student struct
type Student struct {
    Name    string
    RollNo  int
    Division string
    CollegeName string
}

func main() {
    // Initialize a Student instance
    student := Student{
        Name:       "John Doe",
        RollNo:     12345,
        Division:   "A",
        CollegeName: "XYZ University",
    }

    // Print the student's details
    fmt.Println("Student Name:", student.Name)
    fmt.Println("Roll Number:", student.RollNo)
    fmt.Println("Division:", student.Division)
    fmt.Println("College Name:", student.CollegeName)
}
