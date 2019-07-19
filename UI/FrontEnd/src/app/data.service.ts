import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'; // adding http module to use for http request



@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(private http: HttpClient) { }

  firstClick(){
    return console.log("clicked")
  }


  getBCVersion(){
    return this.http.get('https://equinixmiddleware.herokuapp.com/getbcversion')
  }


  getMiddlewareStatus(){
    return this.http.get('https://equinixmiddleware.herokuapp.com/middlewarestatus')
  }

  getOICStatus(){
    return this.http.get('https://equinixmiddleware.herokuapp.com/oicstatus')
  }

  getApp1status(){
    return this.http.get('https://equinixmiddleware.herokuapp.com/app1status')
  }

  getApp2status(){
    return this.http.get('https://equinixmiddleware.herokuapp.com/app2status')
  }

  getRenewals(){
    //return this.http.get('https://reqres.in/api/users'); //connecting to free api to pull data
    return this.http.get('https://equinixmiddleware.herokuapp.com/renewalList')
  }

  getRecentTransactions(){
    //return this.http.get('https://reqres.in/api/users'); //connecting to free api to pull data
    return this.http.get('https://equinixmiddleware.herokuapp.com/renewalHistory')
  }

  sendTestRenewal(testRenewal){
    return this.http.post('https://equinixmiddleware.herokuapp.com/sendTestRenewal', testRenewal)
    


  }

  

}
