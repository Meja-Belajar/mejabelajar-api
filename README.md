# Meja Belajar Web API

This repository serves as the Meja Belajar Web API backend/server codebase, as part of the Software Engineering [[COMP6100001]](https://curriculum.binus.ac.id/course/COMP6100/) course. The project's objective is to equip students with the essential knowledge and skills required for success in software development, covering various aspects such as the software development lifecycle, process models, project management, architecture, and quality assurance. Each contributor in this repository plays a vital role in the Software Engineering project, with every group member contributing based on their assigned roles and responsibilities.


## Tech Stack

**Client:** React, Redux, TailwindCSS, Next.UI, Typescript

**Server:** Go, Gin-Gonic, GORM, PostgreSQL


## Features

For more information about this project, please refer to this [GitHub profile](https://github.com/Meja-Belajar).

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`APP_NAME = "Meja Belajar"`

`PORT = "3000"`

`DB_URL = "host=localhostuser=<username> password=<password> dbname=mejabelajar_db port=5432 ssldmode=disable"`

`JWT_SECRET_KEY = "YOUR_SECRET_KEY"`

`TIMEOUT = "3s"`

## Run Locally

Clone the project

```bash
git clone https://github.com/Meja-Belajar/mejabelajar-api.git
```

Go to the project directory

```bash
cd mejabelajar-api
```

Install dependencies

```bash
go get .
```

Start the server

```bash
go run .
```

For Web Application installation, please refer to [MejaBelajar-Web](https://github.com/Meja-Belajar/mejabelajar-web)

## Demo

Static deployment of this site available at [Here](https://meja-belajar.github.io/mejabelajar-web/)