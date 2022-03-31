# Movies

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 res | result class
 
## AddCast
Administrators can add new cast in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/cast/addCast  
**Auth required**: Need token  
If add movies error   
```
{
  "message": "add cast error",
  "code":    -1,
}
```  
If add cast success
```
{
 "message": "add cast success",
 "code":    200,
 "movie":   res,
}
```    
The result class will be same as the cast struct.

## DeleteCast
Administrators can delete cast in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/cast/DeleteCast  
**Auth required**: Need token  
If delete cast error   
```
{
  "message": "delete cast error",
  "code":    -1,
}
```  
If delete cast success
```
{
 "message": "delete casts success",
 "code":    200,
}
```   
The result class will be same as the cast struct.

## UpdateCast
Administrators can update Cast in the system.  
**Method**:POST   
**URL**:http://localhost:8000/admin/cast/UpdateCast  
**Auth required**: Need token   
If update Cast error   
```
{
  "message": "update Cast error",
  "code":    -1,
}
```  
If update Cast success
```
{
 "message": "update Cast success",
 "code":    200,
 "movie":   res,
}
```   
The result class will be same as the cast struct.

### SearchCastByName
search cast by a specific cast name.  
**Method**:POST   
**URL**:http://localhost:8000/admin/cast/SearchCast  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the Cast",
  "code":    -1,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "movies":   res,
}
``` 
The result class will be same as the cast struct.
