import { Component, OnInit } from '@angular/core';
import { ReactiveFormsModule, FormControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataService } from '../data.service';

@Component({
  selector: 'app-contact',
  templateUrl: './submit.component.html',
  styleUrls: ['./submit.component.scss']
})
export class SubmitComponent implements OnInit {
  
  loading: boolean = false;
  messageForm: FormGroup;
  error:boolean = false;
  errormessage:string = "";
  testRenewalResponse;
  middleware: Object;
  errormessagecode;


  constructor(
    private formBuilder: FormBuilder,
    private data: DataService, 
    private transactionsdata: DataService, 
    private middlejson: DataService, 
    ) { }

    // this is taken from the chaincode
    // renewalId := args[0]
    // curPrice := args[1]
    // piPercent := args[2]
    // piCap := args[3]
    // piPeriod := args[4]
    // newPrice := args[5]
    // state := args[6]


  ngOnInit() {
    this.messageForm = this.formBuilder.group({
      channel: ['renewals', Validators.required],
      chaincode: ['end2end', Validators.required],

      assetid: ['1', Validators.required],
      currentprice: ['$1000', Validators.required],
      pricepercent: ['10%', Validators.required],
      pricecap: ['$10000', Validators.required],
      priceperiod: ['1 year', Validators.required],
      newprice: ['$99', Validators.required],
      state: ['Rejected', Validators.required]
    });
  }


  onSubmit(messageForm){
  //test for valid form
  //submit to api
  //turn on loading
  //if error turn off loading - turn on error message
  
/// not working

    if(this.messageForm.invalid){ // if the messae form is not valid then return
      return;
    }else{ // else set loading to true, send data to api
    
      this.loading = true // turn on loading message
      
    console.log("this is the message form value -",this.messageForm.value);
    var stringy = JSON.stringify(this.messageForm.value);
    console.log("this is the message form value stringified -",stringy)

      this.data.sendTestRenewal(stringy).subscribe( 
        res => { 
          this.testRenewalResponse = res;
          console.log(res);
          if(this.testRenewalResponse.statusCode != 200 ||this.testRenewalResponse.statusCode == null ){
            this.loading = false;
            this.error = true;
            this.errormessage = this.testRenewalResponse.error;
            this.errormessagecode = this.testRenewalResponse.statusCode
          }
        },
        err => {
          console.log("Error occured");


        }
      );

//error code
    //   {
    //     "returnCode": "Failure",
    //     "info": {
    //         "proxyError": "Proposal not pass",
    //         "peerErrors": [
    //             {
    //                 "peerId": "EquinixFounderpeer0",
    //                 "errMsg": "Sending proposal to EquinixFounderpeer0 failed because of: gRPC failure=Status{code=UNKNOWN, description=chaincode error (status: 500, message: This renewal already exists: 4), cause=null}",
    //                 "verified": false
    //             },
    //             {
    //                 "peerId": "EquinixFounderpeer1",
    //                 "errMsg": "Sending proposal to EquinixFounderpeer1 failed because of: gRPC failure=Status{code=UNKNOWN, description=chaincode error (status: 500, message: This renewal already exists: 4), cause=null}",
    //                 "verified": false
    //             }
    //         ]
    //     },
    //     "txid": "d7b172008255120053399f884bc0a52b0c598f8df031138aa841511f8e82a443"
    // }




    }

   
  

    
  }


}
