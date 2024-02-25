import autoprefixer from "autoprefixer"
import cssnano from "cssnano"
import nesting from "tailwindcss/nesting"
import postcss from "postcss"
import postcssImport from "postcss-import"
import tailwindcss from "tailwindcss"
import twConfig from "../tailwind.config"

const css = await Bun.file("./css/style.css").text()
const processed = await postcss([
  postcssImport(),
  nesting(),
  tailwindcss(twConfig),
  autoprefixer(),
  cssnano()
])
  .process(css, { from: 'css/style.css', to: 'static/global.css' })

await Bun.write("./static/global.css", processed.css, {createPath: true})
if (processed.map) {
  await Bun.write("./static/global.css.map", processed.map.toString(), {createPath: true})
}

