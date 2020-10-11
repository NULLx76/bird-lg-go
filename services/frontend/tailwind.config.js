module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  purge: {
    enabled: true,
    content: ["./templates/*.qtpl"]
  },
  theme: {
    extend: {},
  },
  variants: {},
  plugins: [],
}
