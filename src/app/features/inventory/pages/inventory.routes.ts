import { Routes } from "@angular/router";

export const routes: Routes = [
    {
        path: '',
        loadComponent: () =>
            import('./sell/sell.component').then(m => m.SellComponent)
    },
]