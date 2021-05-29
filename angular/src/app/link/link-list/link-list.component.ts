import { Component, OnInit } from '@angular/core';
import { Apollo, gql } from 'apollo-angular';
import LinkInterface from '../interface/link.interface';

@Component({
  selector: 'app-link-list',
  templateUrl: './link-list.component.html',
  styleUrls: ['./link-list.component.sass'],
})
export class LinkListComponent implements OnInit {
  links: LinkInterface[] = [];
  selectedLink: LinkInterface | undefined;
  loading = true;

  constructor(private apollo: Apollo) {}

  ngOnInit(): void {
    // this.apollo
    //   .watchQuery({
    //     query: gql`
    //       {
    //         links {
    //           title
    //           address
    //         }
    //       }
    //     `,
    //   })
    //   .valueChanges.subscribe((links: LinkInterface[]) => {
    //     this.links = links;
    //   });
  }

  selectLink(link: LinkInterface) {
    this.selectedLink = link;
  }
}
