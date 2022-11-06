/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: "#0366d6",

          secondary: "#28a745",

          accent: "#0366d6",

          neutral: "#161621",

          "base-100": "#1b1f23",

          info: "#0366d6",

          success: "#28a745",

          warning: "#f66a0a",

          error: "#d73a49",
        },
        mocha: {
          primary: "#f2cdcd",

          secondary: "#cba6f7",

          accent: "#f9e2af",

          neutral: "#1e1e2e",

          "base-100": "#11111b",

          info: "#74c7ec",

          success: "#a6e3a1",

          warning: "#fab387",

          error: "#f38ba8",
        }
      },
    ],
  },
  plugins: [
    require("daisyui")
  ],
};
