import htmlmin from "html-minifier-terser"
const now = String(Date.now())
/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
export default async function(eleventyConfig) {
    eleventyConfig.setInputDirectory("src")
    eleventyConfig.setLayoutsDirectory("layouts")
    eleventyConfig.setOutputDirectory("dist")
    eleventyConfig.addPassthroughCopy({ "src/assets": "assets" });


    // allows hot reloading when tailwind reloads
    eleventyConfig.setServerOptions({
        watch: ["dist/style.css"]
    })

    // inspired from https://www.11ty.dev/docs/transforms/#minify-html-output
    eleventyConfig.addTransform("htmlmin", async function (content) {
        // Only minify if in PRODUCTION and if an HTML output
        if (process.env.ELEVENTY_PRODUCTION && (this.page.outputPath || "").endsWith(".html")) {
            let minified = htmlmin.minify(content, {
                useShortDoctype: true,
                removeComments: true,
                collapseWhitespace: true
            })
            return minified
        }
        return content
    })
}