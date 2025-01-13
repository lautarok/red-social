import { Routes } from '@angular/router';

export const routes: Routes = [
	{
		path: 'auth',
		loadChildren: () => 
				import("./features/auth/pages/auth.routes").then(m => m.routes)
	},
	{
		path: '',
		loadComponent: () =>
			import('./shared/components/frame/frame.component').then(m => m.FrameComponent),
		children: [
			{
				path: 'feed',
				loadChildren: () =>
					import('./features/posts/pages/posts.routes').then(m => m.routes)
			},
			{
				path: '**',
				redirectTo: 'feed'
			}
		]
	},
	{
		path: '**',
		pathMatch: 'full',
		redirectTo: 'auth'
	}
];
