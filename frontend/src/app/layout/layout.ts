import { Component } from '@angular/core';
import {RouterOutlet , RouterLink , RouterLinkActive} from '@angular/router';

@Component({
  imports: [RouterOutlet, RouterLink, RouterLinkActive],
  selector: 'app-layout',
  styleUrl: './layout.scss',
  templateUrl: './layout.html',
})
export class Layout {}
