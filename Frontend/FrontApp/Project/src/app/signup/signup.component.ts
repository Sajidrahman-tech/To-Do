import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { User } from '../user.model';
import { UserService } from '../user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {


    user:User=new User('','','','','')
  constructor(public userService:UserService) { }

  ngOnInit(): void {
  }
   
  onSubmit(form:NgForm){
      console.log("insid the onsubmit of the signup")
      console.log(form.value.Firstname)
      console.log(form.value.Lastname)
      console.log(form.value.Email)
      console.log(form.value.password)
      if(form.value.password.length<8)
      window.alert("Password too short")
      return
      console.log(form)
      this.user.Firstname=form.value.Firstname
      this.user.Lastname=form.value.Lastname
      this.user.Email=form.value.Email
      this.user.Password=form.value.password
      this.user.repassword=form.value.repassword
      this.userService.signUp(this.user);

  }
}
