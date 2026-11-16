import { Routes } from '@angular/router';

import {Layout} from './layout/layout';
import { Login } from './pages/login/login';
import { Register } from './pages/register/register';
import { Dashboard } from './pages/dashboard/dashboard';
import { Urls } from './pages/urls/urls';
import { Analytics } from './pages/analytics/analytics';

export const routes: Routes = [
  { path: 'login', component: Login },
  { path: 'register', component: Register },


  //Protested Routes
{
  path:'',
  component:Layout ,
  children:[
    { path: 'dashboard', component: Dashboard },
    { path: 'urls', component: Urls },
    { path: 'analytics', component: Analytics },

  { path: '', redirectTo: 'dashboard', pathMatch: 'full' },]

},
  { path: '**', redirectTo: 'dashboard' }
];
