package main

import "fmt"

// using concept of Interfaces in GO Lang (polymorphism)
type Person interface {
	getDetails()
}

type Student struct {
	Name       string
	RollNumber int
	Courses    []string
}

type Professor struct {
	Name          string
	Department    string
	CoursesTaught []string
}

type Course struct {
	Name             string
	Credit           int
	EnrolledStudents []string
}

// By this I can ask to student to choose thier subjects
func (s *Student) Enroll(subject string) {
	s.Courses = append(s.Courses, subject)
}

// printing details of student and professor Interface
func (s *Student) getDetails() {
	fmt.Printf("Student details: Name: %s, Roll Number: %d, Subjects: %v\n", s.Name, s.RollNumber, s.Courses)
}

func (p *Professor) getDetails() {
	fmt.Printf("Professor details: Name: %s, Department: %s, Subjecttaught: %v\n", p.Name, p.Department, p.CoursesTaught)
}
//Enrolling students in the courses accordingly
func (c *Course) addStudents(student *Student) {
	courseTakenByStudnet := student.Courses
	for _, str := range courseTakenByStudnet {
		if str == c.Name {
			c.EnrolledStudents = append(c.EnrolledStudents, student.Name)
			fmt.Println(student.Name, "added")
		}
	}
}
func enrollStudentInCourse(students []Student, courses []Course) {
	for i := range courses {
		for j := range students {
			courses[i].addStudents(&students[j]) //Method of course
		}
	}
}
func main() {
	student1 := Student{Name: "Prem Kumar", RollNumber: 1}
	student2 := Student{Name: "Shreyas Katiyar", RollNumber: 2}

	prof1 := Professor{Name: "Dr. Raj Kumar Pathak", Department: "Chemcial Engineering", CoursesTaught: []string{"Heat Transfer", " ", "Mathematics"}}
	prof2 := Professor{Name: "Dr. Deepu Jha", Department: "Computer Science", CoursesTaught: []string{"DSA", " ", "OOPS"}}

	//Task1 : student choosing subjects
	student1.Enroll("Heat Transfer")
	student2.Enroll("Fluid Mechanics")
	student2.Enroll("DSA")

	//Task2 : Printing the details of student
	//Task3 : Printing the details of professor
	// list of student and professor as a Person interface
	people := []Person{&student1, &student2, &prof1, &prof2}
	fmt.Println("\n-----Getting detalis of Student and professors-----")
	for _, Person := range people {
		Person.getDetails()
	}
	//Task4: Course instance creation and adding student accordingily if they have taken this course
	course1 := Course{Name: "Heat Transfer", Credit: 4}
	course2 := Course{Name: "DSA", Credit: 8}

	courses := []Course{course1, course2}
	enrollStudentInCourse([]Student{student1, student2}, courses)

	fmt.Println("\n-----Printing details of courses-----")
	for _, course := range courses {
		fmt.Printf("Course Name: %s, Credit: %d, EnrollStudents: %v\n", course.Name, course.Credit, course.EnrolledStudents)
	}
}
