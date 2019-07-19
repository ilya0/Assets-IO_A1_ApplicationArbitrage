import { Component, OnInit, Input } from '@angular/core';
import { DataService } from '../data.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashComponent implements OnInit {

  loading: boolean = true;
  renewals: Object;
  h1Style: boolean = false;

  constructor(private data: DataService) { }

  ngOnInit() {
    //placing a get user subscribe on the on init lifecycle hook so it populates when app is opened
    this.data.getRenewals().subscribe(data =>{
      this.renewals = data;
      console.log(this.renewals);
      console.log(data);
      this.loading = false
    });
  }

  firstClick(){
    console.log('clicked');
    this.h1Style = true;
    this.data.firstClick();
  }

}
