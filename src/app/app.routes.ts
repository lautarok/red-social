import { Routes } from '@angular/router';
import { noAuthGuard } from './core/guards/no-auth.guard';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
	{
		path: 'auth',
		canActivate: [noAuthGuard],
		loadChildren: () => 
				import("./features/auth/pages/auth.routes").then(m => m.routes)
	},
	{
		path: '',
		loadComponent: () =>
			import('./shared/components/frame/frame.component').then(m => m.FrameComponent),
		canActivate: [authGuard],
		children: [
			{
				path: 'posts',
				loadChildren: () =>
					import('./features/posts/pages/posts.routes').then(m => m.routes)
			},
			{
				path: 'chat',
				loadChildren: () =>
					import('./features/chat/pages/chat.routes').then(m => m.routes)
			},
			{
				path: '**',
				redirectTo: 'posts'
			}
		]
	},
	{
		path: '**',
		pathMatch: 'full',
		redirectTo: ''
	}
];
