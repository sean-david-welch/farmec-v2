/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ}", "./**/*.templ", "./styles/*.css"],
      theme: {
          extend: {
              screens: {
                  md: { max: '896px' },
                  lg: { min: '896px' },
              },
          },
          plugins: [],
      },
  plugins: [],
}