const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  purge: ["./src/**/*.svelte"],
  theme: {
    typography: (theme) => ({
      default: {
        css: {
          h1: {
            color: theme("colors.white"),
          },
          color: theme("colors.white"),
        },
      },
    }),
    container: {
      center: true,
      padding: "1rem",
    },
    extend: {
      fontFamily: {
        sans: ["Barlow", ...defaultTheme.fontFamily.sans],
      },
    },
  },
  variants: {},
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/ui")],
};
