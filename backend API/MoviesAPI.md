# Movies

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 
## AddNewMovies
Administrators can add new movies in the system.
**Method**:POST  
**URL**:http://localhost:8000/movie/addMovie 
**Auth required**: Need token  
If add movies error   
```
{
  "message": "add movies error",
	"code":    -1,
}
```  
If add movies success
```
{
 "message": "add movies success",
 "code":    200,
}
```    

## DeleteMovies
Administrators can delete movies in the system.
**Method**:POST  
**URL**:http://localhost:8000/movie/deleteMovie 
**Auth required**: Need token  
If delete movies error   
```
{
  "message": "delete movies error",
	"code":    -1,
}
```  
If delete movies success
```
{
 "message": "delete movies success",
 "code":    200,
}
```   

## UpdateMovies
Administrators can update movies in the system.
**Method**:POST  
**URL**:http://localhost:8000/movie/updateMovie 
**Auth required**: Need token  
If update movies error   
```
{
  "message": "update movies error",
	"code":    -1,
}
```  
If update movies success
```
{
 "message": "update movies success",
 "code":    200,
}
```   
