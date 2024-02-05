/** @type {import('tailwindcss').Config} */
export const content = [
	'./templates/**/*.templ',
	'./views/**/*.templ',
];
export const theme = {};

export const daisyui = {
	themes: ["black"]
}

export const plugins = [
	require("@tailwindcss/typography"), require("daisyui")
];
export const corePlugins = { preFlight: true };
