import { Routes } from '@angular/router';

export const routes: Routes = [
	{
		path: 'login',
		pathMatch: 'full',
		loadComponent: () =>
			import("./login/login.component").then(m => m.LoginComponent)
	},
	{
		path: 'sign-in',
		pathMatch: 'full',
		loadComponent: () =>
			import("./sign-in/sign-in.component").then(m => m.SignInComponent)
	},
	{
		path: '**',
		pathMatch: 'full',
		redirectTo: 'login'
	}
];
