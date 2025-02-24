import { useLocation } from "@solidjs/router"
import { createEffect, createSignal } from "solid-js"

export const title = createSignal<null | string>(null);

export default function Title() {
    const location = useLocation()
    const [t] = title;

    createEffect(() => {
        console.log(location.pathname)
        const titleEle = document.querySelector("title") as HTMLTitleElement | null;

        if (t() === null) return;
        if (titleEle === null) return;

        if (t() !== titleEle.textContent) {
            titleEle.textContent = t();
        }
    }, [location.pathname])

    return null
}
