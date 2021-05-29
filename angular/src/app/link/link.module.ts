import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LinkComponent } from './link.component';
import { LinkListComponent } from './link-list/link-list.component';

@NgModule({
  declarations: [LinkComponent, LinkListComponent],
  imports: [CommonModule],
})
export class LinkModule {}
