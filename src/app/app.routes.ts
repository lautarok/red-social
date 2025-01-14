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
				path: 'chat',
				loadChildren: () =>
					import('./features/chat/pages/chat.routes').then(m => m.routes)
			},
			{
				path: '**',
				redirectTo: 'chat'
			}
		]
	},
	{
		path: '**',
		pathMatch: 'full',
		redirectTo: ''
	}
];
