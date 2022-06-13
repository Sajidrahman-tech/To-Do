import { Component, OnInit } from '@angular/core';
import { TaskService } from '../task.service';
import { Test } from '../test.model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  test:Test=new Test('','')
  bool:Boolean=false
  constructor(public taskService:TaskService) { }

  ngOnInit(): void {
  }
  //to add a division */
  public items:any= [

  ];
  
  /* A two-way binding performed which
     pushes text on division */
  public newTask:any;
  public updTask:any;
  public UpdTaskId:any;

  /* When input is empty, it will
     not create a new division */
  public addToList(newTsk:string) {

    if (newTsk==undefined) {
      window.alert("Task Is Empty")
    }else{
      var valu=localStorage.getItem('UserId')
      this.test.UserId=localStorage.getItem('UserId')
      this.test.desc=newTsk
      console.log(this.test.desc)
    
    this.taskService.add(this.test).subscribe((data:any)=>{
      console.log(data)
   this.taskService.getTask(valu).subscribe((resp:any)=>{
    console.log("inside the get task servie")  
   // console.log(resp)
    this.items=resp;
    console.log(this.items)
    })
    
      
    })
    }
  }
  
  /* This function takes to input the
     task, that has to be deleted*/
  public deleteTask(index:any) {
    console.log("inside the delete function")
    console.log(index)
     this.taskService.delTask(index).subscribe((data:any)=>{
      console.log(data)
      console.log("returned back from the delTask ackend")
      var valu=localStorage.getItem('UserId')
      this.taskService.getTask(valu).subscribe((resp:any)=>{
        console.log("inside the get task servie")  
       // console.log(resp)
        this.items=resp;
        console.log(this.items)
        })
    })
     
  }

////////////

public updateTask(index:any) {
  this.bool=!this.bool
  console.log("inside the update function")

  console.log(index)
  console.log("the task to be updated is"+this.updTask)
  this.UpdTaskId=index;
  console.log("inside the update the id to be updated "+index)
  
  console.log(this.bool)
   
}
public Update(UpdTask:any){
  console.log("inside the Update")
  console.log("the task id to be updated is"+this.UpdTaskId)
  console.log("the task desc to be updated is"+this.updTask)
  this.taskService.updTask(this.UpdTaskId,UpdTask).subscribe((data)=>{
    console.log("succesfully returned the value from  backend")
    var valu=localStorage.getItem('UserId')
    this.taskService.getTask(valu).subscribe((resp:any)=>{
      console.log("inside the get task servie")  
     // console.log(resp)
      this.items=resp;
      this.bool=!this.bool
      console.log(this.items)
      })
  })

}



}

