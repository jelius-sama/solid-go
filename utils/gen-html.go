package utils

import (
	"fmt"
)

type HTMLTemplate struct {
	Body string
	Head string
}

func GenerateHTML(template HTMLTemplate) string {
	vars := GetENVVars()

	if vars.ENV == "dev" {
		return fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8" />
				<meta name="description" content="Discover, share, and support art like never before. A vibrant platform for artists and enthusiasts." />
				<meta name="theme-color" media="(prefers-color-scheme: light)" content="#ffffff">
				<meta name="theme-color" media="(prefers-color-scheme: dark)" content="#09090B">
				<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no, viewport-fit=contain">
				<meta name="format-detection" content="telephone=no" />

				<link rel="icon" href="/static/assets/icon.png" />
				<link rel="apple-touch-icon" href="/static/assets/icon.png" />
				
				%s
			</head>
			<body>
				%s
				%s
				%s
			</body>
			</html>
		`, template.Head, noScriptHTML(), template.Body, devModeScript())
	} else {
		return fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				%s
				%s
			</head>
			<body>
				%s
				%s
			</body>
			</html>
		`, template.Head, prodModeScript(), noScriptHTML(), template.Body)
	}
}

func noScriptHTML() string {
	return `
		<noscript>
		<style>
			:root {
			--background: 0 0% 100%;
			--foreground: 240 10% 3.9%;
			--card: 0 0% 100%;
			--card-foreground: 240 10% 3.9%;
			--tw-ring-offset-shadow: 0 0 #0000;
			--tw-ring-shadow: 0 0 #0000;
			--tw-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
			--muted-foreground: 240 3.8% 46.1%;
			--border: 240 5.9% 90%;
			--secondary: 240 4.8% 95.9%;
			--secondary-foreground: 240 5.9% 10%;
			--ring: 240 5.9% 10%;
			--destructive: 0 84.2% 60.2%;
			--font-feature-settings: "liga" 1, "calt" 1;
			--font-variation-settings: "wght" 700;
			}
			@media (prefers-color-scheme: dark) {
			:root {
				--background: 240 10% 3.9%;
				--foreground: 0 0% 98%;
				--secondary: 240 3.7% 15.9%;
				--secondary-foreground: 0 0% 98%;
				--card: 240 10% 3.9%;
				--card-foreground: 0 0% 98%;
				--tw-shadow-colored: 0 1px 2px 0 var(--tw-shadow-color);
				--tw-ring-offset-shadow: 0 0 #0000;
				--tw-ring-shadow: 0 0 #0000;
				--tw-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
				--destructive: 0 62.8% 30.6%;
				--ring: 240 4.9% 83.9%;
				--border: 240 3.7% 15.9%;
				--muted-foreground: 240 5% 64.9%;
			}
			}
			html,
			:host {
			line-height: 1.5;
			-webkit-text-size-adjust: 100%;
			-moz-tab-size: 4;
			tab-size: 4;
			font-size: 16px;
			font-family: theme(
				"fontFamily.sans",
				ui-sans-serif,
				system-ui,
				sans-serif,
				"Apple Color Emoji",
				"Segoe UI Emoji",
				"Segoe UI Symbol",
				"Noto Color Emoji"
			);
			-webkit-tap-highlight-color: transparent;
			font-feature-settings: var(--font-feature-settings, normal);
			font-variation-settings: var(--font-variation-settings, normal);
			}
			*,
			::before,
			::after {
			box-sizing: border-box;
			}
			body {
			margin: 0;
			line-height: inherit;
			background-color: hsl(var(--background));
			color: hsl(var(--foreground));
			}
			button {
			font-family: inherit;
			font-feature-settings: inherit;
			font-variation-settings: inherit;
			font-size: 100%;
			font-weight: inherit;
			line-height: inherit;
			letter-spacing: inherit;
			color: inherit;
			margin: 0;
			padding: 0;
			text-transform: none;
			}
			h3 {
			margin: 0;
			font-size: inherit;
			font-weight: inherit;
			}
			a {
			color: inherit;
			text-decoration: inherit;
			}
			@media (min-width: 640px) {
			.sm-max-w-xl {
				max-width: 36rem /* 576px */;
			}
			}
			.hover\:bg-secondary\/80:hover {
			background-color: hsl(var(--secondary) / 0.8);
			}
			.focus-visible\:outline-none:focus-visible {
			outline: 2px solid transparent;
			outline-offset: 2px;
			}
			.space-y-1\.5 > :not([hidden]) ~ :not([hidden]) {
			--tw-space-y-reverse: 0;
			margin-top: calc(
				0.375rem /* 6px */ * calc(1 - var(--tw-space-y-reverse))
			);
			margin-bottom: calc(0.375rem /* 6px */ * var(--tw-space-y-reverse));
			}
			.focus-visible\:ring-offset-2:focus-visible {
			--tw-ring-offset-width: 2px;
			}
			.focus-visible\:ring-ring:focus-visible {
			--tw-ring-color: hsl(var(--ring));
			}
			.ring-offset-background {
			--tw-ring-offset-color: hsl(var(--background));
			}
			.transition-colors {
			transition-property: color, background-color, border-color,
				text-decoration-color, fill, stroke;
			transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
			transition-duration: 150ms;
			}
			.focus-visible\:ring-2:focus-visible {
			--tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0
				var(--tw-ring-offset-width) var(--tw-ring-offset-color);
			--tw-ring-shadow: var(--tw-ring-inset) 0 0 0
				calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
			box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow),
				var(--tw-shadow, 0 0 #0000);
			}
		</style>
		<main
			style="
			width: calc(100vw - (1rem * 2));
			height: calc(100vh - (1rem * 2));
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			margin: 1rem;
			"
		>
			<div
			style="
				border-radius: 0.5rem;
				border-width: 1px;
				border-color: hsl(var(--border));
				border-style: solid;
				background-color: hsl(var(--card));
				color: hsl(var(--card-foreground));
				box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
				var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
				width: 100%;
			"
			class="sm-max-w-xl"
			>
			<div
				style="display: flex; flex-direction: column; padding: 1.5rem"
				class="space-y-1.5"
			>
				<h3
				style="
					font-size: 1.125rem;
					font-weight: 600;
					line-height: 1;
					letter-spacing: -0.025em;
				"
				>
				Error
				</h3>
				<p
				style="
					font-size: 0.875rem;
					line-height: 1.25rem;
					color: hsl(var(--muted-foreground));
				"
				>
				JS Disabled
				</p>
			</div>
			<div style="padding: 1.5rem; padding-top: 0">
				<span style="color: hsl(var(--destructive))">
				You need to enable JavaScript to run this app.
				</span>
			</div>
			<div
				style="
				align-items: center;
				padding: 1.5rem;
				padding-top: 0;
				display: flex;
				justify-content: end;
				"
			>
				<a
				href="https://github.com"
				style="
					display: inline-flex;
					align-items: center;
					justify-content: center;
					gap: 0.5rem;
					white-space: nowrap;
					border-radius: calc(0.5rem - 2px);
					font-size: 0.875rem;
					line-height: 1.25rem;
					font-weight: 500;
					padding-top: 0.5rem;
					padding-bottom: 0.5rem;
					padding-left: 1rem;
					padding-right: 1rem;
					height: 2.5rem;
					background-color: hsl(var(--secondary));
					color: hsl(var(--secondary-foreground));
				"
				class="focus-visible:ring-2 ring-offset-background transition-colors focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:outline-none hover:bg-secondary/80"
				>
				Return
				</a>
			</div>
			</div>
		</main>
		</noscript>
	`
}

func devModeScript() string {
	return `<script src="http://localhost:5173/src/index.tsx" type="module"></script>`
}

func prodModeScript() string {
	return `
		<script type="module" crossorigin src="/assets/main.js"></script>
    	<link rel="stylesheet" crossorigin href="/assets/index.css">
	`
}
