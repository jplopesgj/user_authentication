# User Authentication Template.
Basic Template for user authentication

## Installation
1 - Clone this repository to your local machine:

```
git clone <repository-url>
```
2 - Inside the project directory, run the following command to install the required dependencies:

```bash
go get
```

3 - Next, you need to install two external libraries: make and migrate.

To install make, you can use Chocolatey (a package manager for Windows) by running the following command:


```bash
choco install make
```

4 - To install migrate, you can use Scoop (a command-line installer for Windows) by running the following commands:

```bash
scoop install migrate
```
You can find more information about migrate library on its GitHub repository: https://github.com/golang-migrate/migrate.

PS:

`Update the database connection configuration in database/config file with your own database connection details.`

`If you need to make any changes to the database schema, modify the .sql files in the database/migrations directory.`

Now you're ready to use the user authentication template. Feel free to explore the code and customize it according to your needs.

### This template is open-source. Feel free to update and send to the repository.


## Authors

- [@jplopesgj](https://www.github.com/jplopesgj)