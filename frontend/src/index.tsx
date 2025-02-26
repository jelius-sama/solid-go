/* @refresh reload */
import '@/index.css';
import { render } from 'solid-js/web';
import { createEffect, createSignal, JSX, JSXElement } from 'solid-js';
import { A, Router, useLocation } from '@solidjs/router';
import { routes } from '@/routes';
import { MetaProvider, Link } from "@solidjs/meta";
import logo from "@/static/logo.jpg";
import { Dynamic } from "solid-js/web";

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

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
      fetch(`${location.pathname}/props`)
        .then((res) => res.json())
        .then((data) => {
          const ssrElement = parsePropsToJSX(data);
          setServerEles((prev) => ({ ...prev, [location.pathname]: ssrElement }));
        })
        .catch(console.error);
    }
  });

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
  if (!props?.body?.children) {
    console.error("Invalid SSR props format:", props);
    return null;
  }

  function renderElement(element: any): JSX.Element {
    if (!element) return null;

    // Handle text nodes
    if (typeof element === "string") return element;

    // Ensure tag is valid
    if (!element.tag) {
      console.warn("Skipping element without a tag:", element);
      return null;
    }

    // Ensure children are correctly rendered
    const children = Array.isArray(element.children)
      ? element.children.map(renderElement) // Recursively render child elements
      : element.children;

    // Use `Dynamic` from solid-js/web for dynamic elements
    return (
      <Dynamic component={element.tag} id={element.id}>
        {children}
      </Dynamic>
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