
# Review

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 res | result class
 
## AddNewReview
Users can add new reviews in any movies.  
**Method**:POST  
**URL**:http://localhost:8000/user/review/addReview  
**Auth required**: Need token  
If add reviews error   
```
{
  "message": "add reviews error",
  "code":    -1,
}
```  
If add reviews success
```
{
 "message": "add reviews success",
 "code":    200,
}
```    
The result class will be same as the review struct.

## DeleteReviews
Users can delete reviews in the system.  
**Method**:POST  
**URL**:http://localhost:8000/movie/deleteReview  
**Auth required**: Need token  
If delete reviews error   
```
{
  "message": "delete reviews error",
  "code":    -1,
}
```  
If delete reviews success
```
{
 "message": "delete reviews success",
 "code":    200,
}
```   
The result class will be same as the review struct.

## UpdateReviews
Users can update thier own reviews in the system.  
**Method**:POST   
**URL**:http://localhost:8000/movie/updateReview  
**Auth required**: Need token   
If update reviews error   
```
{
  "message": "update reviews error",
  "code":    -1,
}
```  
If update reviews success
```
{
 "message": "update reviews success",
 "code":    200,
}
```   
