# Instagam

Instagam is a RestFull API Social Media project. This project is made using the Go programming language and using the Gin framework. In the implementation of programming and layout, it uses Clean Architecture

The project structure is as follows:
| Layer | Directory |
|----------------------|----------------|
| Frameworks & Drivers | Infrastructures|
| Interface | Interfaces |
| Usecases | Usecases |
| Entities | Domain |

## Getting Started

To start running this project locally,

```bash
git clone https://github.com/adiprrassetyo/instagam.git
```

Open mygram-api folder and install all required dependencies

```bash
cd instagam && go mod tidy
```

Copy the example env file and adjust the env file

```
cp .env.example .env
```

Start the server

- The First Way
  To run directly the main.go file, you can use the following command:
  ```bash
  go run ./app/main.go
  ```
- The Second Way
  To run using Makefile, you can use the command as follows:
  ```bash
  Make run
  ```
- The Third Way
  To run using nodemon, you can use the command as follows:

  ```bash
  nodemon --exec go run ./app/main.go
  ```

  Or you can use a Makefile that has provided the nodemon command:

  ```bash
  Make run-nodemons
  ```

  For complete details on how to run using Makefile can be seen in the Makefile file.

Check the Instagam API documentation

```
[https://documenter.getpostman.com/view/23538862/2s9YC5wX3x](https://documenter.getpostman.com/view/23538862/2s9YC5wX3x)
```
