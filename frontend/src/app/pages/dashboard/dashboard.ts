import { Component } from '@angular/core';

@Component({

  selector: 'app-dashboard',
  standalone: true,
  imports: [],
  styleUrl: './dashboard.scss',
  templateUrl: './dashboard.html',
})
export class Dashboard {
  stats = [
    { title: 'Total URLs', value: '128', icon: '🔗' },
    { title: 'Total Clicks', value: '2,450', icon: '👆' },
    { title: 'Created This Month', value: '18', icon: '📅' },
    { title: 'Active URLs', value: '96%', icon: '📊' }
  ];

recentUrls = [
  {
    originalUrl: 'https://www.example.com/products/electronics',
    shortUrl: 'linktrack.com/abc123',
    clicks: 245,
    status: 'Active'
  },
  {
    originalUrl: 'https://www.example.com/blog/web-development',
    shortUrl: 'linktrack.com/xyz789',
    clicks: 182,
    status: 'Active'
  },
  {
    originalUrl: 'https://www.example.com/services',
    shortUrl: 'linktrack.com/def456',
    clicks: 96,
    status: 'Expired'
  },
  {
    originalUrl: 'https://www.example.com/contact',
    shortUrl: 'linktrack.com/ghi012',
    clicks: 54,
    status: 'Active'
  }
];

}
