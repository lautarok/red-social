import { Routes } from '@angular/router';

export const routes: Routes = [
	{
		path: '',
		pathMatch: 'full',
		loadComponent: () =>
			import("./feed/feed.component").then(m => m.FeedComponent)
	},
	{
		path: '**',
		redirectTo: ''
	}
]