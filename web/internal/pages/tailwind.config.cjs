const fs = require('fs')
const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
  	"./index.html",
  	"./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    colors,
  },
  variants: {
    extend: {},
  },
  plugins: [],
  darkMode: 'class',
}
