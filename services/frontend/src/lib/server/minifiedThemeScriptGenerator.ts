/*
 * Generates minified script that adds "dark" class for 
 * <html> tag if dark theme prefered
 *
 * @remarks
 *
 * Separated from src/components/functional/themeApplier.astro
 * to be build time compiled
 */

import * as esbuild from "esbuild";
import themeScript from "@/lib/browser/theme.js?raw";

const autoRunScript = `
${ themeScript };
document.documentElement.classList.toggle("dark", getTheme() === "dark");
`;

const autoRunMinifiedScript = esbuild.transformSync(autoRunScript, {
  minify: true,
  loader: "js",
  treeShaking: false,
  format: "esm"
});

export default autoRunMinifiedScript;
