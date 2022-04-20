# Description of IMDB
IMDB_Clone is a web application that we mimic the IMDB website. User could browse and search movie and cast information and detail. Meanwhile, they can write a comment or rate the movie they like, but prerequisite is they need to login. We also provide a sign up function for users who like our website. For users who already login, they can update or delete their account. For Administrator, they can conduct the movie and cast data, such as add, update or delete data. For some inappropriate comments, admin could delete the commemt, even delete user account.

Demo video functionality - 3min max

Cypress test video - 90sec max

Backend unit test video - 90sec max

[Link to API Documentation](https://github.com/BazingaOo/IMDB_Clone/tree/main/backend%20API)  

[Link to Project board](https://github.com/BazingaOo/IMDB_Clone/projects)

[Link to Sprint4 deliverables](https://github.com/BazingaOo/IMDB_Clone.git) Note: The whole project is in the frontend and backend folders.
# what we do
## front end
### Adding more Details in HomePage
For aesthetic design of our homepage, we refill the picture size to 
make movie posters showing like IMax screen. \
Moreover, we find more suitable-size pictures as our Top-Movies' covers.

### For 
implemented the actor, comment, write comment page\
embellished all the pages previously created\
redesigned some pages\
cypress test for login

## back end
### For Comment model
basic CURD for comment    
add a pagination to display comments for each movie  
add jwt middleware to the router  
### For Rating model
basic CURD for rating    
add a function that could display the trending movie  
add a computing average socore function for each function  
add jwt middleware to the router  
### For Genre model
basic CURD for genre  
add genre to the search result for each movie  
add jwt middleware to the router  
# How to run
## Install npm package
```
npm install
```
## Run the front server
```
npm run serve
```
the homepage is localhost:8080
## Run the backend server
open in the GoLand software and config the go build, then click the start button, the backend address is localhost:8000

# Frontend and backend team members
## Group Members
### front-end
Yuxuan Wang, 7566-9009<br/>
Shihuan Wang, 9715-8829<br/>
### back-end
Tao Zhang, 7636-6624<br/>
Zihan Guo, 8615-3487<br/>
