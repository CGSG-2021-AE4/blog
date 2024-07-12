const esbuild = require("esbuild")

function buildPage(page) {
  esbuild
    .build({
      entryPoints: ["web/src/pages/" + page + ".tsx"],
      outdir: "web/bin/assets/scripts",
      bundle: true,
      minify: true,
      plugins: [],
    })
    .then(() => console.log("Build page " + page + " complete"))
    .catch(()=> process.exit(1));
}

buildPage("index")
buildPage("login")
buildPage("signup")
buildPage("account")
buildPage("article")