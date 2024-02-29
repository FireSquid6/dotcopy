import { Layout } from "inkdocs";

const PageLayout: Layout = (children) => {
  return (
    <div id="layout">
      <main id="content">{children}</main>
      <footer>
        <p>© 2021</p>
      </footer>
    </div>
  );
};

export default PageLayout;
