import { Component } from '@angular/core';

@Component({
  selector: 'app-analytics',
  standalone: true,
  imports: [],
  styleUrl: './analytics.scss',
  templateUrl: './analytics.html',
})
export class Analytics {
  stats = [
    { title: 'Total Clicks', value: '2,450', icon: '👆' },
    { title: 'Unique Visitors', value: '1,820', icon: '👥' },
    { title: 'Average Clicks / URL', value: '61', icon: '🔗' },
    { title: 'Top Performing URL', value: 'abc123', icon: '🏆' }
  ];

  clicksData = [
    { day: 'Mon', clicks: 120 },
    { day: 'Tue', clicks: 185 },
    { day: 'Wed', clicks: 90 },
    { day: 'Thu', clicks: 240 },
    { day: 'Fri', clicks: 310 },
    { day: 'Sat', clicks: 195 },
    { day: 'Sun', clicks: 270 }
  ];

  topUrls = [
    { shortUrl: 'linktrack.com/abc123', clicks: 245 },
    { shortUrl: 'linktrack.com/xyz789', clicks: 182 },
    { shortUrl: 'linktrack.com/def456', clicks: 96 },
    { shortUrl: 'linktrack.com/ghi012', clicks: 54 }
  ];

  browsers = [
    { name: 'Chrome', percentage: 65 },
    { name: 'Firefox', percentage: 15 },
    { name: 'Edge', percentage: 12 },
    { name: 'Other', percentage: 8 }
  ];
}
