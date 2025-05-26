/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./views/**/*.templ}", "./**/*.templ"],
    theme: {
      extend: {
        
      },
    },
    plugins: [
      require('@tailwindcss/forms'),
    ],
  }