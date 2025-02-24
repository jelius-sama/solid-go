/* @refresh reload */
import '@/index.css';
import { render } from 'solid-js/web';
import { JSX } from 'solid-js';
import { A, Router } from '@solidjs/router';
import { routes } from '@/routes';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

const App = ({ children }: { children: JSX.Element }) => {
  return (
    <main class='m-4'>
      <nav class='flex gap-x-2'>
        <A class="text-blue-600 font-semibold text-lg" href='/'>Home</A>
        <A class="text-blue-600 font-semibold text-lg" href='/settings'>Settings</A>
        <A class="text-blue-600 font-semibold text-lg" href='/profile'>Profile</A>
      </nav>
      {children}
    </main>
  )
}

render(
  () => <Router root={(props) => <App>{props.children}</App>}>{routes}</Router>,
  root,
);