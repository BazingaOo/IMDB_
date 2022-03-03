# UserSignUp  
User sign up  
**Method**:POST  
**URL**:http://localhost:8000/user/signUp  
**Auth required**: Need token  
If sign up error   
```
{
  "message": "sign up error",
	"code":    404,
}
```  
If sign up success
```
{
 "message": "sign up success",
 "code":    200,
}
```    
parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
