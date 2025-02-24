import { title } from "@/components/layout/title";
import { createRenderEffect } from "solid-js";

export default function NotFoundPage() {
    const [_, setTitle] = title;
    createRenderEffect(() => {
        setTitle("Not Found | Solid + Go")
    })

    return (
        <div>404 - Found Not Page</div>
    )
}
