# myhttp

Simple tool to create in memory organization(simliar to multinde tree) and to find closest manager of 2 emplyees.
This is a common problem to find the common nearest parent for a child.

## Description

- The tool accepts user input to create in memory organization structure
- User can add employees to manager within the organization
- User can find the nearest manager of given 2 employees withing the organization

## Usage

- Clone the repository using git clone.
- Build the project using command ``` go build ```
- Run the program using command ``` ./organization ```
- There are 4 choices for user
  - create organization - should be called only once, this will create new organization with 1 employee(CEO with emp ID - 0)
  - add employees to manager - new employess can be added to manager. Manager should be present in the organization(at the start only CEO will be present).
  - nearest manager for the given 2 employees can be found by entering the emp ID of the users
  - print the organization structure in preorder fashion(tree preorder traversal)
- To run the test cases, run command- ``` go test ./... ```

## Assumptions

- The organization is hierarchical, (i.e. tree structured)
- If emp1 is manager of emp2, then their nearest manager will be emp1
- emp IDs are unique within organization and allowed only number.
- CEO's employee ID will always be 0
- To start adding employees, add employees under CEO and so on
