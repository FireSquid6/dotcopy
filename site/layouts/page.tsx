import { Layout } from "inkdocs";

const PageLayout: Layout = (children) => {
  return (
    <div id="layout">
      <main id="content">{children}</main>
      <footer>
        <p>© Jonathan Deiss 2024</p>
      </footer>
    </div>
  );
};

export default PageLayout;
