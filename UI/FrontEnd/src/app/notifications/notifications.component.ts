import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';

@Component({
  selector: 'app-about',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.scss']
})
export class NotificationsComponent implements OnInit {

//status indicators true - functioning / false - non functioning
  networkstatus: boolean;
  blockchainstatus: boolean = false; // api accesses blockchain and gets version thus verifying functionality
  middlewarestatus: boolean = false; //api accesses  middleware server status
  oicstatus: boolean = false;
  app1status: boolean = false; 
  app2status: boolean = false;


  renewals: Object; //renewal list
  transactions: Object; //transactions list
  bcversion: Object; //bc version light
  middleware: Object;
  oic: Object;
  app1: Object;
  app2: Object;


  //loading icons
  loadingoverall: boolean = true; // overall status indicator
  loadingrenewals: boolean = true; // loading status renewal list
  loadingtransactions: boolean = true; // loading status renewal list
  notifications: string = "None";

  //misc 
  h1Style: boolean = false;
  counter: number = 0;

  constructor(
    private data: DataService, 
    private transactionsdata: DataService, 
    private versionjson: DataService, 
    private middlejson: DataService, 
    private oicjson: DataService, 
    private app1json: DataService, 
    private app2json: DataService
    ) { }

  ngOnInit() {


// renewal list api
    this.data.getRenewals().subscribe(data =>{
      this.renewals = data;
      console.log(this.renewals);
      console.log(data);
      this.loadingrenewals = false;

    });


    // renewal list api
    this.transactionsdata.getRecentTransactions().subscribe(transactionsdata =>{
      this.transactions = transactionsdata;
      console.log("transactions list - ", this.transactions);
      console.log(transactionsdata);
      this.loadingtransactions = false;

    });

//blockchain version and light
    this.versionjson.getBCVersion().subscribe(versionjson =>{
      this.bcversion = versionjson;
      //this.version = null; //***** */testing ****
      console.log("version var - ", this.bcversion);
      
      if(this.bcversion != null){
        this.blockchainstatus = true;
        console.log("blockchainstatus set to functioning")
      }else{
        this.blockchainstatus = false;
        console.log("blockchainstatus set to non-functioning")
      }
      

      this.runcounter();
    });


    //middleware status
    this.middlejson.getMiddlewareStatus().subscribe(middlejson =>{
      this.middleware = middlejson;
      //this.version = null; //***** */testing ****
      console.log("middle ware status ", this.middleware);
      
      if(this.middlejson != null){
        this.middlewarestatus = true;
        console.log("middleware status set to functioning")
      }else{
        this.middlewarestatus = false;
        console.log("middleware status set to non-functioning")
      }
      this.runcounter();
    });



    //oic status
    this.oicjson.getOICStatus().subscribe(oicjson =>{
      this.oic = oicjson;
      
      //this.version = null; //***** */testing ****
      console.log("oic status ", this.oic["status"]);
      console.log('oicjson', this.oic)
      
      if(this.oic["status"] = "ACTIVATED"){
        this.oicstatus = true;
        console.log("OIC Status set to functioning")
      }else{
        this.oicstatus = false;
        console.log("OIC Status set to non-functioning")
      }
      this.runcounter();
    });


       //App1 status
       this.app1json.getApp1status().subscribe(app1json =>{
        this.app1 = app1json;
        //this.version = null; //***** */testing ****
        console.log("app1 status ", this.app1);
        
        if(this.app1json != null){
          this.app1status = true;
          console.log("app1 Status set to functioning")
        }else{
          this.app1status = false;
          console.log("app1 Status set to non-functioning")
        }
        this.runcounter();
      });

 //App2 status
 this.app2json.getApp2status().subscribe(app2json =>{
  this.app2 = app2json;
  //this.version = null; //***** */testing ****
  console.log("app2 status ", this.app2);
  
  if(this.app2json != null){
    this.app2status = true;
    console.log("app2 status set to functioning")
  }else{
    this.app2status = false;
    console.log("app2 Status set to non-functioning")
  }
  this.runcounter();
});


    

  this.loadingoverall = false;
  }

  //evaluate network status
overallStat(){
  // change light on Overall master status
  if( 
    this.blockchainstatus == false  || 
    this.oicstatus == false  || 
    this.middlewarestatus == false ){
    //if any components is false, set network to down
    this.networkstatus = false;
    console.log("downstatus", this.blockchainstatus, this.oicstatus, this.middlewarestatus);
  }else{
    this.networkstatus = true;
  }

  console.log("networkstatus - ", this.networkstatus)
  console.log("blockchainstatus - ", this.blockchainstatus)
  console.log("middlewarestatus - ", this.middlewarestatus)
  console.log("oicstatus - ", this.oicstatus)
  console.log("app1 - ", this.app1)
  console.log("app2 - ", this.app2)
}

//run counter for overall network status
runcounter(){
  this.counter++;
      if(this.counter==5){
        this.overallStat();
      }
}

//testing a click
  firstClick(){
    console.log('clicked');
    this.h1Style = true;
    this.data.firstClick();
  }


}

