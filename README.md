<div align="center">
  <h1> Go Api Pet </h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/go-api-pet"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/go-api-pet"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/go-api-pet"/>
</div>
<div align="center">
 <a href="#-description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-runner">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</div>

## ⭐️ Description

The application is an API to manage the registry of pets.

## 📍 Requirement

You will need to have installed golang on your machine, my version is 1.19. Docker to run of database or you can use another way.

## 🚀 run

In project root:

To initial the database run the command below:

```sh
docker-compose up -d
```

To run the project use the command below:

```sh
go run main.go
```

The project uses port _3000_, on your browser access **[documentation](http://localhost:3000/api/v1/pets/swagger/index.html)**

## ⚙️ Technologies

- golang
- postgres
- sqlite
- gin
- swagger

<br>
<span align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</span>
