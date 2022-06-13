import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Test } from './test.model';
@Injectable({
  providedIn: 'root'
})
export class TaskService {

  constructor(public http:HttpClient) { }

add(test:Test):Observable<any>
{
 console.log("inside the addd")
  console.log(test)

  return this.http.post('http://localhost:9090/postTask',test)
}
getTask(Valu:any):Observable<any>
{
  const take:any={
    Valu:Valu,
    
    }
console.log("isnide get task")
console.log(take)
 return this.http.post('http://localhost:9090/getTask',take)
}
delTask(Id:any) :Observable<any>
{
  const take:any={
   Id:Id
    
    }
console.log("inside the delTask servie the vlaue of taskId is"+Id)
return this.http.post('http://localhost:9090/delTask',take)
}
updTask(Id:any,Desc:any):Observable<any>
{
  const take:any={
    Id:Id,
    Desc:Desc 
     }
     console.log("inside the delTask servie the vlaue of taskId is"+Id)
     console.log("inside the delTask servie the vlaue of desc is"+Desc)
    return this.http.post('http://localhost:9090/updTask',take)

     
}


}
