/* @refresh reload */
import '@/index.css';
import { render, renderToString } from 'solid-js/web';
import { createEffect, createSignal, JSX, JSXElement } from 'solid-js';
import { A, Router, useLocation } from '@solidjs/router';
import { routes } from '@/routes';
import { MetaProvider, Link } from "@solidjs/meta";
import logo from "@/static/logo.jpg";

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

// type ServerEles = Array<{
//   [pathname: string]: JSXElement
// }>

// const App = ({ children }: { children: JSX.Element }) => {
//   const location = useLocation();
//   const [serverEles, setServerEles] = createSignal<ServerEles>([])

//   createEffect(() => {
//     const serverEle = document.querySelectorAll("#__server__");
//     serverEle.forEach((el) => {
//       setServerEles((prev) => ([...prev, { [location.pathname]: el }]))
//       el.remove();
//     })
//   })

//   return (
//     <main class='m-4'>
//       <nav class='flex gap-x-2'>
//         <A class="text-blue-600 font-semibold text-lg" href='/'>Home</A>
//         <A class="text-blue-600 font-semibold text-lg" href='/settings'>Settings</A>
//         <A class="text-blue-600 font-semibold text-lg" href='/profile'>Profile</A>
//       </nav>
//       {serverEles() && serverEles().map((serverEl) => {
//         const [pathname, element] = Object.entries(serverEl)[0];
//         return pathname === location.pathname ? <>{element}</> : null;
//       })}
//       {children}
//     </main>
//   )
// }

type ServerEles = { [pathname: string]: JSX.Element };

const App = ({ children }: { children: JSX.Element }) => {
  const location = useLocation();
  const [serverEles, setServerEles] = createSignal<ServerEles>({});

  // Store SSR elements on first load
  createEffect(() => {
    const serverEle = document.querySelectorAll("#__server__");
    serverEle.forEach((el) => {
      setServerEles((prev) => ({ ...prev, [location.pathname]: el }));
      el.remove();
    });
  });

  // Fetch and store SSR content dynamically based on route
  createEffect(() => {
    if (!serverEles()[location.pathname]) {
      console.log(location.pathname)
      fetch(`${location.pathname}/props`) // Dynamic URL based on current route
        .then((res) => res.json())
        .then((data) => {
          console.log("Hi", data)
          const ssrElement = parsePropsToJSX(data);
          console.log("ssrElement: ", ssrElement)
          setServerEles((prev) => ({ ...prev, [location.pathname]: ssrElement }));
        })
        .catch(console.error);
    }
  });

  createEffect(() => {
    console.log("elementes: ", serverEles()[location.pathname])
  })

  return (
    <main class="m-4">
      <nav class="flex gap-x-2">
        <A class="text-blue-600 font-semibold text-lg" href="/">Home</A>
        <A class="text-blue-600 font-semibold text-lg" href="/settings">Settings</A>
        <A class="text-blue-600 font-semibold text-lg" href="/profile">Profile</A>
      </nav>
      {/* Render stored SSR elements if available */}
      {serverEles()[location.pathname] ?? null}
      {children}
    </main>
  );
};

function parsePropsToJSX(props: any): JSX.Element {
  function renderElement(element: any, p0?: { id: any; }, p1?: any): JSX.Element {
    if (!element.tag) return element.children; // Handle text nodes

    return renderElement(
      element.tag, // The tag name (e.g., "p", "div")
      { id: element.id }, // Attributes (currently just "id")
      Array.isArray(element.children)
        ? element.children.map(renderElement) // Recursively render children
        : element.children
    );
  }

  return <>{props.body.children.map(renderElement)}</>;
}


render(
  () => (
    <Router
      root={(props) =>
        <MetaProvider>
          <Link rel="icon" href={logo} />
          <Link rel="apple-touch-icon" href={logo} />

          <App>
            {props.children}
          </App>
        </MetaProvider>
      }>
      {routes}
    </Router>
  ),
  root,
);