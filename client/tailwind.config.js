const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  future: {
    purgeLayersByDefault: true,
    removeDeprecatedGapUtilities: true,
  },
  purge: ["./src/**/*.svelte", "./src/**/*.html"],
  theme: {
    typography: (theme) => ({
      default: {
        css: {
          h1: {
            color: theme("colors.gray.800"),
          },
          color: theme("colors.gray.800"),
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
