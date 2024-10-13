import { Component } from '@angular/core';
import { EventService } from '../../services/event.service';

@Component({
  selector: 'app-event-create',
  templateUrl: './event-create.component.html'
})
export class EventCreateComponent {
  event = {
    title: '',
    description: '',
    location: '',
    date: ''
  };
  selectedFile: File | null = null;

  constructor(private eventService: EventService) {}

  onFileSelected(event : any) {
    this.selectedFile = event.target.files[0];
  }

  createEvent() {
    this.eventService.createEvent(this.event).subscribe((data: any) => {
      if (this.selectedFile) {
        this.eventService.uploadFile(this.selectedFile).subscribe();
      }
      alert('Event created successfully');
    });
  }
}
