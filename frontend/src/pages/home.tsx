import { title } from "@/components/layout/title"
import { createRenderEffect } from "solid-js"

export default function HomePage() {
    const [_, setTitle] = title;
    createRenderEffect(() => {
        setTitle("Home | Solid + Go")
    })

    return (
        <div>Home Page</div>
    )
}
