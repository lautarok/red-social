import { Routes } from '@angular/router';

export const routes: Routes = [
	{
		path: 'login',
		pathMatch: 'full',
		loadComponent: () =>
			import("./login/login.component").then(m => m.LoginComponent)
	},
	{
		path: '**',
		pathMatch: 'full',
		redirectTo: 'login'
	}
];
