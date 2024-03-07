import { InkdocsOptions } from "inkdocs";
import swapRouter from "inkdocs/plugins/swap-router";
import "@kitajs/html/register";
import { devserverPlugin } from "inkdocs-server";
import PageLayout from "./layouts/page";
// this needs to be fixed
// import SyntaxHighlighter from "inkdocs-highlight-plugin";

export function getOptions(): InkdocsOptions {
  const baseHtml = `<html>
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <link rel="stylesheet" href="/styles.css" />
  <title>Dotcopy</title>

  <meta name="description" content="Keep your dotfiles DRY" />
  <meta name="author" content="Firesquid6" />
  <meta name="keywords" content="linux, dotfiles, dotcopy, html, programming, dot, copy, file, dotfiles, manager, nix, nixos, DRY, dry" />
</head>
  {slot}
</html>`;

  const options: InkdocsOptions = {
    staticFolder: "static",
    buildFolder: "build",
    contentFolder: "content",
    baseHtml,
    layouts: new Map([["page", PageLayout]]),
    craftsmen: [],
    layoutTree: {
      path: "",
      layoutName: "page",
      children: [],
    },
    plugins: [swapRouter({}), devserverPlugin()],
    server: {
      port: 3000,
    },
  };

  return options;
}

const options = getOptions();
export default options;
