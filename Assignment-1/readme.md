# Student Data CLI

### Description

Student Data CLI is a command-line service used to display data of classmates. This program utilizes os.Args to fetch arguments from the terminal. The displayed data includes name, address, occupation, and reason for choosing the Golang class.

### Usage

To use this service, execute the following command:

```bash
go run main.go <student_number>
```

Replace <student_number> with the student number of the friend whose data you want to see.

### Example

For example, to view the data of a student with student number 1, execute the command:

```bash
go run main.go 1
```

### Program Structure

- The program uses structs to represent friend data, including name, address, occupation, and reason for choosing the Golang class.

- There is a function `printStudentById` used to print friend data based on the provided student number.
