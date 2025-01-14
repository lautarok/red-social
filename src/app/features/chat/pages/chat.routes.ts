import { Routes, UrlSegment } from '@angular/router';

export const routes: Routes = [
	{
		matcher: (url) => {
			if (url.length === 1 && url[0].path.match(/^\d+$/g)) {
				return {consumed: url, posParams: {
					conversation: new UrlSegment(url[0].path, {})
				}}
			}
			return {consumed: url}
		},
		loadComponent: () =>
			import('./inbox/inbox.component').then(m => m.InboxComponent)
	}
]