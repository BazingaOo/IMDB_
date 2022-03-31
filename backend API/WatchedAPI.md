

# Watched List

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 res | result class
 
## AddWatched
Users can add movies they have watched in the watched list.  
**Method**:POST  
**URL**:http://localhost:8000/user/watchedList/addWatchedList  
**Auth required**: Need token  
If add watched error   
```
{
  "message": "add watched error",
  "code":    -1,
}
```  
If add watched success
```
{
 "message": "add watched success",
 "code":    200,
}
```    
The result class will be same as the watched_list struct.

## DeleteWatched
Users can delete movies they have watched in the watched list. 
**Method**:POST  
**URL**:http://localhost:8000/user/watchedList/deleteWatchedList  
**Auth required**: Need token  
If delete watched error   
```
{
  "message": "delete watched error",
  "code":    -1,
}
```  
If delete watched success
```
{
 "message": "delete watched success",
 "code":    200,
}
```   
The result class will be same as the review struct.

## ReadWatched
Users can see all the movies they have watched in the watched list. 
**Method**:POST   
**URL**:http://localhost:8000/user/watchedList/readWatchedList  
**Auth required**: Need token   
If read watched list error   
```
{
  "message": "read watched list error",
  "code":    -1,
}
```  
If read watched list success
```
{
 "message": "read watched list success",
 "code":    200,
}
```   
