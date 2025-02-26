import { Title as T } from "@solidjs/meta"

export default function Title({ children }: { children: string }) {
    return <T>{children + " - Solid + Go"}</T>
}
