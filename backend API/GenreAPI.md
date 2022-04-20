
# Genres

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 res | result class
 
## AddGenre
Administrators can add new genre in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/cast/addCast  
**Auth required**: Need token  
If add genre error   
```
{
  "message": "add genre error",
  "code":    404,
}
```  
If add genre success
```
{
 "message": "add genre success",
 "code":    200,
 "genre":   res,
}
```    
The result class will be same as the genre struct.

## DeleteGenre
Administrators can delete genre in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/cast/deleteCast  
**Auth required**: Need token  
If delete genre error   
```
{
  "message": "delete genre error",
  "code":    404,
}
```  
If delete genre success
```
{
 "message": "delete genre success",
 "code":    200,
 "genre":   res,
}
```   
The result class will be same as the genre struct.

## UpdateGenre
Administrators can update genre in the system.  
**Method**:POST   
**URL**:http://localhost:8000/admin/cast/updateCast  
**Auth required**: Need token   
If update genre error   
```
{
  "message": "update genre error",
  "code":    404,
}
```  
If update genre success
```
{
 "message": "update genre success",
 "code":    200,
 "genre":   res,
}
```   
The result class will be same as the genre struct.

## SearchGenreName
search genre by genre name.  
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchGenreName  
**Auth required**: do not need token   
If search genre error   
```
{
  "message": "search genre error",
  "code":    404,
}
```  
If search genre success
```
{
 "message": "search genre success",
 "code":    200,
 "genre":   res,
}
``` 
The result class will be same as the genre struct.

## SearchGenreByGenreName
find all the movies belonging to one specific genre.
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchGenreByGenreName  
**Auth required**: do not need token   
If search genre error   
```
{
  "message": "search genre error",
  "code":    404,
}
```  
If search genre success
```
{
 "message": "search genre success",
 "code":    200,
 "genre":   res,
}
``` 
The result class will be same as the genre struct.

## SearchMovieWithGenre
search movie and genre by movieId
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchMovieWithGenre  
**Auth required**: do not need token   
If search genre error   
```
{
  "message": "search genre error",
  "code":    404,
}
```  
If search genre success
```
{
 "message": "search genre success",
 "code":    200,
 "genre":   res,
}
``` 
The result class will be same as the genre struct.

