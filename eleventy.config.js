import pluginRss from "@11ty/eleventy-plugin-rss";
import htmlmin from "html-minifier-terser";
import syntaxHighlight from "@11ty/eleventy-plugin-syntaxhighlight";
import { DateTime } from "luxon";
const now = String(Date.now());
/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
export default async function (eleventyConfig) {
  eleventyConfig.setInputDirectory("src");
  eleventyConfig.setLayoutsDirectory("layouts");
  eleventyConfig.setOutputDirectory("dist");
  eleventyConfig.addPassthroughCopy({ "src/assets": "assets" });

  // allows hot reloading when tailwind reloads
  eleventyConfig.setServerOptions({
    watch: ["dist/style.css"],
  });

  // inspired from https://www.11ty.dev/docs/transforms/#minify-html-output
  eleventyConfig.addTransform("htmlmin", async function (content) {
    // Only minify if in PRODUCTION and if an HTML output
    if (
      process.env.ELEVENTY_PRODUCTION &&
      (this.page.outputPath || "").endsWith(".html")
    ) {
      let minified = htmlmin.minify(content, {
        useShortDoctype: true,
        removeComments: true,
        collapseWhitespace: true,
      });
      return minified;
    }
    return content;
  });
  eleventyConfig.addPreprocessor("drafts", "*", (data, content) => {
    if (data.draft && process.env.ELEVENTY_RUN_MODE === "build") {
      return false;
    }
  });
  eleventyConfig.addFilter("date", function (date) {
    return DateTime.fromISO(date.toISOString()).toFormat("dd-MM-yyyy");
  });
  eleventyConfig.addFilter("sortEntriesByYear", function (entries) {
    const years = {};
    entries.forEach((entry) => {
      const date = new Date(entry.date);
      const year = date.getFullYear();
      if (!years[year]) {
        years[year] = [];
      }
      years[year].push({
        ...entry,
      });
    });

    // Sort each year's entries by date descending
    Object.keys(years).forEach((year) => {
      years[year].sort((a, b) => new Date(b.date) - new Date(a.date));
    });

    // Sort the years descending
    const sortedYears = Object.entries(years).sort(
      ([yearA], [yearB]) => Number(yearB) - Number(yearA)
    );

    return sortedYears;
  });

  eleventyConfig.addFilter("sortEntriesByTopics", function (entries) {
    const topics = {};
    entries.forEach((entry) => {
      entry.data.topics.forEach((topic) => {
        if (!topics[topic]) {
          topics[topic] = 0
        }
        topics[topic] += 1
      });
    });

    console.log(topics)
    return topics;
  })

  eleventyConfig.addPairedShortcode("terminal", function (content) {
    const lines = content.split("\n").filter((l) => l.trim() !== "");
    const escaped = lines
      .map((l) => l.trim().replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;"))
      .join("\n");
    const raw = lines.map((l) => l.trim()).join("\n");
    const dataAttr = raw.replace(/"/g, "&quot;");

    const html = lines
      .map(
        (l) =>
          `<span class="select-none text-green-400">$ </span><span class="text-neutral-100">${l.trim().replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;")}</span>`
      )
      .join("\n");

    return `<div class="terminal relative rounded-lg bg-neutral-800 p-4 my-4" data-commands="${dataAttr}">
<button onclick="copyTerminal(this)" class="select-none absolute top-2 right-2 text-xs text-neutral-400 hover:text-neutral-100 transition-colors">copy</button>
<pre class="overflow-x-auto !bg-transparent !p-0 !m-0">${html}</pre>
</div>`;
  });

  eleventyConfig.addPlugin(pluginRss);

  eleventyConfig.addPlugin(syntaxHighlight)
}
