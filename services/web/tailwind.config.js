/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      fontFamily: {
        main: ['var(--font-main)'],
        handwritten: ['var(--font-handwritten)'],
        'stroke-order': ['var(--font-stroke-order)'],
      },
    },
  },
  plugins: [],
}
