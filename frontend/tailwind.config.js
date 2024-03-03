// tailwind.config.js
const { nextui } = require("@nextui-org/react");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
    "./src/**/*.{js,ts,jsx,tsx}",
    "./src/**/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'white': '#fffefb',
        'white-accent-1': '#f5f4f1',
        'white-accent-2': '#cccbc8',
        'black': '#1d1c1c',
        'black-accent-1': '#3b3c3d',
        'black-accent-2': '#313d44',
        'blue-accent-100': '#d4eaf7',
        'blue-accent-200': '#b6ccd8',
        'blue-accent-300': '#71c4ef',
        'blue-accent-400': '#00668c',
      },
      keyframes: {
        carousel: {
          '0%, 18%, 85%, 100%': {
            transform: 'translateX(0%)',
            transitionDelay: '1s'
          },
          '20%, 38%, 70%, 83%': {
            transform: 'translateX(-33.333%)',
          },
          '40%, 68%': {
            transform: 'translateX(-66.666%)',
          },
        }
      },
      animation: {
        'carousel': 'carousel 15s linear infinite',
      }
    },
  },
  darkMode: "class",
  plugins: [nextui()]
}