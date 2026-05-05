import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        shrine: {
          red: "#D9331D",
          green: "#3A7D44",
          teal: "#6BAFAF",
          paper: "#F9F7F4",
          ink: "#333333",
        },
      },
    },
  },
  plugins: [],
};

export default config;
