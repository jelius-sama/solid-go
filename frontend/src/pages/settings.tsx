import { title } from "@/components/layout/title";
import { createRenderEffect } from "solid-js";

export default function SettingsPage() {
    const [_, setTitle] = title;
    createRenderEffect(() => {
        setTitle("Settings | Solid + Go")
    })
    return (
        <div>Settings Page</div>
    )
}
