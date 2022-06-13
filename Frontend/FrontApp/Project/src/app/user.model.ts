export class User {

  
    public UserId:number|undefined;
    public Firstname:string ;
    public Lastname:string;
    public Email:string;
    public Password:string;
    public repassword:string;



    constructor(Firstname:string,Lastname:string,Email:string,password:string,repassword:string){
        this.Firstname=Firstname
        this.Lastname=Lastname
        this.Email=Email
        this.Password=password
        this.repassword=repassword
      }
}


