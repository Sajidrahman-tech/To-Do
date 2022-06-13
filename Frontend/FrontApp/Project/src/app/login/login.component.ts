import { Component, OnInit } from '@angular/core';
import { FormsModule, NgForm } from '@angular/forms';
import { Router } from '@angular/router';
FormsModule
import { from } from 'rxjs';
import { UserService } from '../user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(public userService:UserService,public router:Router) { }

  ngOnInit(): void {
  }
  
  onSubmit(form:NgForm){
    console.log("insinde the onsumit of the logint ")
    console.log(form.value.Username)
    console.log(form.value.Password)
    this.userService.validate(form.value.Username,form.value.Password).subscribe(data=>{
      console.log("inside the ts of login  ha ha")
      console.log(data)
      console.log(data.UserId)
      
      if(data.Firstname!="")
      {
        localStorage.setItem('UserId',data.UserId)
        this.router.navigate(['/home'])

      }
      else{
        console.log("user not found")
        window.alert("wrong credentials")  
      }
    })
    
  }
}
