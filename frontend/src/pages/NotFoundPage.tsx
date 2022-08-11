import { ReactComponent as Image } from "assets/400.svg";
import { ErrorPage } from "pages/ErrorPage";

const NotFoundPage = () => (
  <ErrorPage
    svg={<Image className="h-100" />}
    attrUrl="https://storyset.com/web"
    attrContent="Web illustrations by Storyset"
  />
);

class NotFoundError extends Error {}

export default NotFoundPage;
export { NotFoundError };
