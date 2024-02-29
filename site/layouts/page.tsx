import { Layout } from "inkdocs";

const PageLayout: Layout = (children) => {
  return (
    <div id="layout">
      <header>
        <h1>My Site</h1>
      </header>
      <main id="content">{children}</main>
      <footer>
        <p>© 2021</p>
      </footer>
    </div>
  );
};

export default PageLayout;
