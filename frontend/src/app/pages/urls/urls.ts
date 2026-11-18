import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({

  selector: 'app-urls',
  standalone: true,
  imports: [FormsModule],
  styleUrl: './urls.scss',
  templateUrl: './urls.html',
})
export class Urls {
  selectedStatus = 'All';
   urls = [
    {
      originalUrl: 'https://www.example.com/products/electronics',
      shortUrl: 'linktrack.com/abc123',
      clicks: 245,
      createdAt: '2026-11-10',
      status: 'Active'
    },
    {
      originalUrl: 'https://www.example.com/blog/web-development',
      shortUrl: 'linktrack.com/xyz789',
      clicks: 182,
      createdAt: '2026-11-12',
      status: 'Active'
    },
    {
      originalUrl: 'https://www.example.com/services',
      shortUrl: 'linktrack.com/def456',
      clicks: 96,
      createdAt: '2026-11-05',
      status: 'Expired'
    },
    {
      originalUrl: 'https://www.example.com/contact',
      shortUrl: 'linktrack.com/ghi012',
      clicks: 54,
      createdAt: '2026-11-14',
      status: 'Active'
    }
  ];
  searchTerm = '';

get filteredUrls() {
  return this.urls.filter(url => {
    const matchesSearch =
      url.originalUrl.toLowerCase().includes(this.searchTerm.toLowerCase()) ||
      url.shortUrl.toLowerCase().includes(this.searchTerm.toLowerCase());

    const matchesStatus =
      this.selectedStatus === 'All' ||
      url.status === this.selectedStatus;

    return matchesSearch && matchesStatus;
  });
}

copyUrl(shortUrl: string) {
  const fullUrl = `https://${shortUrl}`;

  navigator.clipboard.writeText(fullUrl)
    .then(() => {
      alert('URL copied to clipboard!');
    })
    .catch(() => {
      alert('Failed to copy URL.');
    });
}
deleteUrl(shortUrl: string) {
  const confirmed = confirm('Are you sure you want to delete this URL?');

  if (!confirmed) {
    return;
  }

  this.urls = this.urls.filter(url => url.shortUrl !== shortUrl);
}

}
