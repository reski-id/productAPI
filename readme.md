

<h3 align="center">Product Rest API <br>
<h5 align="center" >Golang Echo Clean Architecture <h5>
<br>
</h4>
<p align="left">
<h2>
  Content <br></h2>
  • Key Features <br>
  • Installing Using Github<br>
  • Installing Using Docker<br>
  • End Point<br>
  • Technologi that i use<br>
  • Contact me<br>
</p>

## 📱 Features

* Product
* Category
* Images


## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/productAPI.git

# Go into the repository
$ cd productAPI

# Install dependencies
$ go get

# load .env env 
# use bash cmd and type this..
$ source .env

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allready create database mysql `dbproduct` for this app.more info in local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/productapi

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name productapi -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="productapi" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/productapi

# Run the app
$ docker logs productapi
```

## 📜 End Point  

Product
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------                  
| `GET`         | /products            | Get all products      
| `DELETE`      | /products/:id         | Delete product  
| `PUT`         | /products/:id         | Update data product
| `POST`        | /products         | Insert product 

Category
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------                  
| `GET`         | /category            | Get all categories       
| `DELETE`      | /category/:id         | Delete category  
| `PUT`         | /category/:id         | Update data category
| `POST`        | /category         | Insert category 

Images
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------      
| `GET`         | /file            | Get all Images       
| `DELETE`      | /file/:id         | Delete Image 
| `PUT`         | /file/:id         | Update data Image
| `POST`        | /file         | Insert Image

Product-Category
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------      
| `GET`         | /productcategory            | Get all product categories     
| `GET`         | /productcategory/:id           | Get one product categories    
| `DELETE`      | /productcategory/:id         | Delete product category 
| `PUT`         | /productcategory/:id         | Update data product category
| `POST`        | /productcategory         | Insert product category

Product-Images
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------      
| `GET`         | /productimages            | Get all product images     
| `GET`         | /productimages/:id           | Get one product images    
| `DELETE`      | /productimages/:id         | Delete product images 
| `PUT`         | /productimages/:id         | Update data product images
| `POST`        | /productimages         | Insert product images


## 🛠️ Technology

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Echo Framework](https://echo.labstack.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Git Repository](https://github.com/reski-id)
- [Trunk Base Development](https://trunkbaseddevelopment.com/)

## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>
