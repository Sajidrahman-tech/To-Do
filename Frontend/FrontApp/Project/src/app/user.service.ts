import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from './user.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http:HttpClient) { 

  }


  signUp(user:User){
    console.log("inside the user service")
    console.log("Firstname is :-"+user.Firstname)
    console.log("Lastname is :-"+user.Lastname)
    console.log("Email is :-"+user.Email)
    console.log("password is :-"+user.Password)
    console.log("repassword is :-"+user.repassword)
    console.log(user)
     this.http.post('http://localhost:9090/postUser',user).subscribe(response=>{
      console.log(response)
      window.alert("Registered Successfully")
    })
  
  }
  validate(Username:string,Password:string):Observable<any>
  {

    const logIn:any={
      Firstname:Username,
      Password:Password
      }
      console.log(logIn);
    console.log("Inside the user service validate")
    console.log("the uyser nma,me is "+Username)
    console.log("the Password is "+Password)
     return this.http.post('http://localhost:9090/validate',logIn)
  }
  
  //////
  logIn(Username:string,Password:string)
  {
    const logIn:any={
      Username:Username,
      Password:Password
      }
      console.log(logIn);
    console.log("Inside the user service login")
    console.log("the uyser nma,me is "+Username)
    console.log("the Password is "+Password)
    this.http.get('http://localhost:9090/getUser',logIn).subscribe(respose=>{
      console.log("returned succesfully"   )
      console.log(respose)
    })
  }


}
