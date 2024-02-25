import {defineConfig} from "turbowatch"

export default defineConfig({
  project: __dirname,
  triggers: [
    {
      expression: [
        "allof",
        ["not", ["dirname", "node_modules"]],
        ["anyof",
          ["match", "*.go", "basename"],
          ["match", "*.html", "basename"]
        ]
      ],
      interruptible: true,
      name: "go-server",
      onChange: async ({spawn}) => {
        await spawn`go run cmd/main.go`
      }
    },
    {
      expression: [
        "allof",
        // ["not", ["dirname", "node_modules"]],
        // ["not", ["dirname", "static"]],
        ["anyof", 
          ["match", "**/*.css", "wholename"],
          ["match", "**/*.html", "wholename"]
        ]
      ],
      interruptible: true,
      name: "postcss",
      onChange: async ({spawn}) => {
        await spawn`bun run cmd/css.ts`
      }
    }
  ]
})
