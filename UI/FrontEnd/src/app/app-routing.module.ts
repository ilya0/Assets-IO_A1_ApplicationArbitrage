import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SubmitComponent } from './submit/submit.component'
import { DashComponent } from './dashboard/dashboard.component';
import { NotificationsComponent } from './notifications/notifications.component';
import { ActionlistComponent } from './actionlist/actionlist.component';
// these are the routes for the different components
const routes: Routes = [
  { path: '', component: DashComponent },
  { path: 'notifications', component:NotificationsComponent },
  { path:'submit', component:SubmitComponent },
  { path:'actionlist', component:ActionlistComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
