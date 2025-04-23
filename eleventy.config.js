import pluginRss from "@11ty/eleventy-plugin-rss";
import htmlmin from "html-minifier-terser";
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
  eleventyConfig.addPlugin(pluginRss);
}
