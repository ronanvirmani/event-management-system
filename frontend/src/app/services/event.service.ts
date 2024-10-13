import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class EventService {
  constructor(private http: HttpClient) {}

  getEvents() {
    return this.http.get('/api/events');
  }

  createEvent(event : any) {
    return this.http.post('/api/events', event);
  }

  uploadFile(file: File) {
    const formData = new FormData();
    formData.append('file', file, file.name);
    return this.http.post('/api/upload', formData);
  }
}
