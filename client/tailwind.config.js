/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        origin: '#267365',
        origin_depth: '#12443B',
        accent: '#33F2D0',
        accent_border: '#3E2705',
      },
    },
  },
  plugins: [],
};
